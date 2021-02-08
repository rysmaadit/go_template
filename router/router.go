package router

import (
	"example.com/m/handler"
	"example.com/m/service"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func NewRouter(_ service.Dependencies) http.Handler {
	r := mux.NewRouter()

	setPingHandler(r)
	setPostSheetHandler(r)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setPingHandler(r *mux.Router) {
	r.Methods(http.MethodGet).Path("/ping").Handler(handler.Ping())
}

func setPostSheetHandler(r *mux.Router) {
	r.Methods(http.MethodPost).Path("/webhook").Handler(handler.PostWebhookHandler())
}
