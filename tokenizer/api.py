from concurrent.futures import ThreadPoolExecutor
import logging
import grpc
import ml_pb2
import ml_pb2_grpc
from parsers.symptoms import SymptomTokenizer


class TokensServerServicer(ml_pb2_grpc.TokensServerServicer):
    def __init__(self):
        self.symptom_tokenizer = SymptomTokenizer()
        self.logger = logging.getLogger(__name__)

    def FindSymptoms(self, request: ml_pb2.SymptomsRequest, context: grpc.RpcContext):
        desc = request.text

        grouped_tokens = self.symptom_tokenizer.get_symptoms(desc)
        self.logger.info(f'Sending over {len(grouped_tokens)}s for request containing text: {desc[:10]}...')

        for symptom in grouped_tokens["symptoms"]:
            res = ml_pb2.Symptom(type="symptoms", name=symptom["name"], loc=symptom["loc"], confidence=float(symptom["score"]))
            yield res

        for symptom in grouped_tokens["treatments"]:
            res = ml_pb2.Symptom(type="treatments", name=symptom["name"], loc=symptom["loc"], confidence=float(symptom["score"]))
            yield res

        for symptom in grouped_tokens["tests"]:
            res = ml_pb2.Symptom(type="tests", name=symptom["name"], loc=symptom["loc"], confidence=float(symptom["score"]))
            yield res
    
    def SayHello(self, request, context):
        hi = ml_pb2.Hello(res="Cool Guy")

        return hi

def serve():
    server = grpc.server(thread_pool=ThreadPoolExecutor(max_workers=10))
    token_server = TokensServerServicer()
    ml_pb2_grpc.add_TokensServerServicer_to_server(
        token_server, server)
    
    server.add_insecure_port('[::]:50051')
    
    server.start()
    token_server.logger.info("ML server is running on port 50051...")
    
    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()
