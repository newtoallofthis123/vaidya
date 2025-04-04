package api

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/newtoallofthis123/patients/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *ApiServer) handleTranscribe(c *gin.Context) {
	fileHeader, err := c.FormFile("content")
	if err != nil {
		c.JSON(500, gin.H{"err": "Unable to read file: " + err.Error()})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(500, gin.H{"err": "Unable to read file: " + err.Error()})
		return
	}
	defer file.Close()

	msg, err := s.transcribe(file)
	if err != nil {
		c.JSON(500, gin.H{"err": "Unable to read file: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"msg": msg})
}

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

	fmt.Println(len(content))

	for i := 0; i < len(content); i += 1024 {
		chunk := content[i:min(i+1024, len(content))]

		err = stream.Send(&types.AudioFile{
			Filename:  "",
			Format:    "ogg",
			AudioData: chunk,
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
