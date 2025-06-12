package service

import (
	"context"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

var logger = slog.Default()

func ChatMessage(w http.ResponseWriter, r *http.Request) {
	logger.Info("Start chat response")

	err := godotenv.Load()
	if err != nil {
		logger.Error("Error load env", slog.String("error", err.Error()))
		log.Fatal(err)
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if err != nil {
		log.Fatal(err)
	}

	question, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error("Error reading request body", slog.String("error", err.Error()))
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(string(question)),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte(result.Text()))
	if err != nil {
		logger.Error("Error writing response", slog.String("error", err.Error()))
		log.Fatal(err)
	}
}
