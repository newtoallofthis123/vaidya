package api

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ollama/ollama/api"
)

func (s *ApiServer) handleGetSummary(c *gin.Context) {
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatal(err)
	}

	prompt := c.PostForm("prompt")

	req := &api.GenerateRequest{
		Model:  "llama3.2:1b",
		Prompt: prompt,
		System: `You are a highly advanced language model trained to analyze and summarize complex medical data. Your task is to process the provided patient data and generate a title and description, each enclosed within their respective tags.
1. **Title**: Provide a brief, descriptive heading that captures the essence of the patient's case, enclosed within "< Title >" tags.
2. **Description**: Offer a detailed description of the patient's condition, including symptoms, test results, diagnosis, and treatment plans, enclosed within "< Description >" tags.
Ensure that the output is in English, avoids any conversational language, and strictly follows the specified format without any additional text. The LLM should only use the data provided and not include any external information. It can interpret the data and suggest possible next steps for the doctor, but it should not schedule appointments or invent any data.
<Title>
[Descriptive title of the patient's case]
</Title>
<Description>
[Detailed description of the patient's condition, history, and relevant details]
</Description>`,

		Stream: new(bool),
	}

	var res api.GenerateResponse
	ctx := context.Background()
	respFunc := func(resp api.GenerateResponse) error {
		res = resp
		return nil
	}

	err = client.Generate(ctx, req, respFunc)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, res)
}
