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
	dependencies := service.InstantiateDependencies()

	setPingHandler(r)
	setPostSheetHandler(r, dependencies)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}

func setPingHandler(r *mux.Router) {
	r.Methods(http.MethodGet).Path("/ping").Handler(handler.Ping())
}

func setPostSheetHandler(r *mux.Router, dependencies service.Dependencies) {
	r.Methods(http.MethodPost).Path("/webhook").Handler(handler.PostWebhookHandler(dependencies))
}
