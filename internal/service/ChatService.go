package service

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/ClaytonMatos84/go-geminiapi/pkg"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

var logger = slog.Default()

func ChatMessage(w http.ResponseWriter, r *http.Request) {
	logger.Info("Start chat response")

	err := godotenv.Load()
	if pkg.CheckError(err, "Error load env") {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: os.Getenv("GEMINI_API_KEY"),
	})
	if pkg.CheckError(err, "Error create client") {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	question, err := io.ReadAll(r.Body)
	if pkg.CheckError(err, "Error reading request body") {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	} else if len(question) == 0 {
		logger.Error("Empty question received")
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	maxOutputTokens := int32(80)
	config := &genai.GenerateContentConfig{
		MaxOutputTokens:  maxOutputTokens,
		ResponseMIMEType: "text/plain",
		SystemInstruction: genai.NewContentFromText(
			"Responda a questão de forma resumida e somente se tiver a certeza da resposta. Retorne de onde foi retirada a resposta e a data.",
			genai.RoleUser,
		),
	}
	logger.Info("Received question", slog.String("question", string(question)))
	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash",
		genai.Text(string(question)),
		config,
	)
	if pkg.CheckError(err, "Error interact with model") {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(result.Text()))
	if pkg.CheckError(err, "Error writing response") {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
}
