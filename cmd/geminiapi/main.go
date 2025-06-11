package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text("Boa noite tudo bom?"),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Text())
}
