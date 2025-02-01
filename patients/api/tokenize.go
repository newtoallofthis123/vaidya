package api

import (
	"context"
	"fmt"

	"github.com/newtoallofthis123/patients/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type JsonSymptom struct {
	Type string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

func (s *ApiServer) tokenizeText(text string) ([]JsonSymptom, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		s.logger.Error("Unable to connect to tokenization server: " + err.Error())
		return nil, err
	}
	defer conn.Close()

	fmt.Println("Trying to connect to microservice")
	client := types.NewTokensServerClient(conn)
	symptoms, err := client.FindSymptoms(context.Background(), &types.SymptomsRequest{Text: text})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected")

	var symptomsList []JsonSymptom

	for _, sym := range symptoms.Symptoms {
		symptomsList = append(symptomsList, JsonSymptom{Type: sym.Type, Name: sym.Name})
	}

	return symptomsList, nil
}
