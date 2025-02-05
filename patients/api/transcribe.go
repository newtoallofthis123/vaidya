package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"github.com/newtoallofthis123/patients/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *ApiServer) transcribe(file multipart.File) (string, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50052", opts...)
	if err != nil {
		s.logger.Error("Unable to read content file with err: " + err.Error())
		return "", err
	}
	defer conn.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	fmt.Println("Trying to connect to microservice")
	client := types.NewAudioServiceClient(conn)
	stream, err := client.TranscribeAudio(context.Background())
	if err != nil {
		s.logger.Error("Unable to read content file with err: " + err.Error())
		return "", err
	}
	fmt.Println("Connected to microservice")

	// send 5kb buffer at a time
	buff := make([]byte, 1024*5)
	fmt.Println(len(content))

	for i := 0; i < len(content); i += 1024 * 5 {
		buff = content[i : i+1024*5]

		err = stream.Send(&types.AudioFile{
			Filename:  "",
			Format:    "ogg",
			AudioData: buff,
		})
		if err != nil {
			s.logger.Error("Unable to read content file with err: " + err.Error())
		}
	}
	fmt.Println("Sent to microservice")
	stream.CloseSend()

	var msg types.TranscribeResponse
	err = stream.RecvMsg(&msg)
	if err != nil {
		log.Fatalf("Failed to receive message: %v", err)
	}

	return msg.Message, nil
}
