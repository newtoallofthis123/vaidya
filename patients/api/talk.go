package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/generative-ai-go/genai"
	"github.com/gorilla/websocket"
	"google.golang.org/api/option"
)

func (s *ApiServer) handleTalk(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading websocket: %v", err)
		return
	}
	log.Println("Client connected to socket service")

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GOOGLE_GEMINI_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-2.0-flash")
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text(systemPrompt())},
	}
	model.SetTemperature(1)
	model.SetTopK(64)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "application/json"

	ses := model.StartChat()
	for {
		_, prompt, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Println("Recieved Message: ", string(prompt))

		if string(prompt) == "bye" {
			delete(s.subs, conn)
			c.JSON(200, gin.H{"success": "Response generated"})
			return
		}

		fmt.Println("Sending message to Gemini")
		resp, err := ses.SendMessage(ctx, genai.Text(string(prompt)))
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}

		res := ""
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					res += fmt.Sprintln(part)
				}
			}
		}

		fmt.Println(res)

		parsed, err := parseToJSON(res)
		parsedJson, _ := json.Marshal(parsed)
		log.Println("Parsed response, writing message")
		conn.WriteMessage(websocket.TextMessage, []byte(parsedJson))
	}
}

type ParsedResponse struct {
	Think        string      `json:"think,omitempty"`
	Analysis     string      `json:"analysis,omitempty"`
	Info         PatientInfo `json:"info,omitempty"`
	NextQuestion string      `json:"next_question,omitempty"`
}

type PatientInfo struct {
	Name              string    `json:"name"`
	Age               string    `json:"age"`
	Gender            string    `json:"gender"`
	Address           string    `json:"address"`
	Identity          string    `json:"identity"`
	Phone             string    `json:"phone"`
	Problems          []Problem `json:"problems"`
	Conditions        []string  `json:"conditions"`
	Description       string    `json:"description"`
	RecommendedDoctor string    `json:"recommended_doctor"`
}

type Problem struct {
	Name        string `json:"name"`
	Duration    string `json:"duration"`
	Description string `json:"description"`
}

func parseToJSON(raw string) (ParsedResponse, error) {
	raw = strings.TrimSpace(raw)

	pattern := regexp.MustCompile(`"((?:[^"\n]|\\.|\\\n)*)(\n)"`)
	raw = pattern.ReplaceAllString(raw, `"${1}\\n"`)

	var parsed ParsedResponse

	err := json.Unmarshal([]byte(raw), &parsed)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return ParsedResponse{}, err
	}

	return parsed, nil
}
