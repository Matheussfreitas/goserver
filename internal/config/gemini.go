package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func Gemini(prompt string) string {
	ctx := context.Background()
	
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY não configurada no ambiente")
	}

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Text())
	return result.Text()
}
