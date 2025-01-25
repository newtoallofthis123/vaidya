package api

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/newtoallofthis123/patients/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (s *ApiServer) handleTranscribe(c *gin.Context) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50052", opts...)
	if err != nil {
		s.logger.Error("Unable to read content file with err: " + err.Error())
		c.JSON(500, gin.H{"err": "Unable to read file: " + err.Error()})
		return
	}
	defer conn.Close()

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
	content, err := io.ReadAll(file)
	if err != nil {
		c.JSON(500, gin.H{"err": "Unable to read file: " + err.Error()})
		return
	}

	fmt.Println("Trying to connect to microservice")
	client := types.NewAudioServiceClient(conn)
	stream, err := client.TranscribeAudio(context.Background())
	if err != nil {
		s.logger.Error("Unable to read content file with err: " + err.Error())
		c.JSON(500, gin.H{"err": "Unable to connect to microservice: " + err.Error()})
		return
	}
	fmt.Println("Connected to microservice")

	// send 5kb buffer at a time
	buff := make([]byte, 1024*5)
	fmt.Println(len(content))

	for i := 0; i < len(content); i += 1024 * 5 {
		buff = content[i : i+1024*5]

		err = stream.Send(&types.AudioFile{
			Filename:  fileHeader.Filename,
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

	c.JSON(200, gin.H{"message": msg.Message})
}
