package router

import (
	"example.com/m/handler"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/ping").Handler(handler.Ping())

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	return loggedRouter
}
