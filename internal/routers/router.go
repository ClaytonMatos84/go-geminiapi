package routers

import (
	"github.com/ClaytonMatos84/go-geminiapi/internal/service"
	"github.com/gorilla/mux"
)

func HandleRouter() *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/chat", service.ChatMessage).Methods("POST")

	return mux
}
