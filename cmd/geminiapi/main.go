package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/ClaytonMatos84/go-geminiapi/internal/routers"
	"github.com/gorilla/handlers"
)

func main() {
	mux := routers.HandleRouter()

	slog.Info("Starting Gemini API server...")
	if err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(mux)); err != nil {
		slog.Error("Failed to start server", "error", err)
		log.Fatalf("Failed to start server: %v", err)
	}
}
