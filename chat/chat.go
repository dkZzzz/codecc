package chat

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
)

const (
	init_question string = `Now I use SonarQube Scan my project, 
	after that, I gain some issue to fix. 
	I want to know how to fix them. 
	Next I will give the code and problem description, 
	and you tell me how to fix it`
)

func Optimize(secret_key string) []string {
	client := openai.NewClient(secret_key)
	req := openai.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: "",
			},
		},
	}
	questions := make([]string, 0)
	answers := make([]string, 0)
	questions = append(questions, init_question)
	questions = append(questions, "	password := \"123456\"，\"password\" detected here, make sure this is not a hard-coded credential.")
	// TODO: Add Question

	for _, question := range questions {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    "user",
			Content: question,
		})
		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			log.Printf("ChatCompletion error: %v\n", err)
			continue
		}
		answers = append(answers, resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
	}
	return answers
}
