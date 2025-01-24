from concurrent.futures import ThreadPoolExecutor
import io
import logging
import wave
import grpc
import transcribe_pb2
import transcribe_pb2_grpc
from whisper import WhisperTranscribe

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s - %(levelname)s - %(message)s",
)


class AudioServiceServicer(transcribe_pb2_grpc.AudioServiceServicer):
    def __init__(self) -> None:
        self.logger = logging.getLogger(__name__)
        self.transcriber = WhisperTranscribe()

    def cal_chunk_size(self, audio_data, seconds=30):
        try:
            with wave.open(io.BytesIO(audio_data), "rb") as wav_file:
                sample_rate = wav_file.getframerate()
                bit_depth = wav_file.getsampwidth() * 8
                num_channels = wav_file.getnchannels()
                bytes_per_second = (sample_rate * bit_depth * num_channels) // 8
                return bytes_per_second * seconds
        except wave.Error as e:
            self.logger.warning(f"Not enough data to parse WAV header: {e}")
            return 0

    def TranscribeAudio(self, request_iterator, context):
        audio_data = bytearray()
        chunk_size = None

        for req in request_iterator:
            audio_data.extend(req.audio_data)

            if chunk_size is None:
                chunk_size = self.cal_chunk_size(audio_data)
                if chunk_size is None:
                    continue

            while len(audio_data) >= chunk_size:
                chunk = audio_data[:chunk_size]
                audio_data = audio_data[chunk_size:]

                self.logger.info("Transcribing 30-second chunk...")
                try:
                    transcription = self.transcriber.transcribe_audio(chunk)
                    yield transcribe_pb2.TranscribeResponse(
                        status="Success", message=transcription
                    )
                except Exception as e:
                    self.logger.error(f"Error during transcription: {e}")
                    context.set_code(grpc.StatusCode.INTERNAL)
                    context.set_details(f"Error: {e}")
                    yield transcribe_pb2.TranscribeResponse(
                        status="Error", message=str(e)
                    )

        if audio_data:
            self.logger.info("Transcribing remaining audio...")
            try:
                transcription = self.transcriber.transcribe_audio(audio_data)
                yield transcribe_pb2.TranscribeResponse(
                    status="Success", message=transcription
                )
            except Exception as e:
                self.logger.error(f"Error during transcription: {e}")
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(f"Error: {e}")
                yield transcribe_pb2.TranscribeResponse(status="Error", message=str(e))


def serve():
    server = grpc.server(
        thread_pool=ThreadPoolExecutor(max_workers=10),
    )
    audio_server = AudioServiceServicer()
    transcribe_pb2_grpc.add_AudioServiceServicer_to_server(audio_server, server)

    server.add_insecure_port("[::]:50052")

    server.start()
    audio_server.logger.info("Audio server is running on port 50052...")

    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    serve()
