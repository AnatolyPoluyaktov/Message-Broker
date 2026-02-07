package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter(msgHandler *MessageHandler) *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/v1").Subrouter()
	api.HandleFunc("queues/{queue_name}/subscriptions", msgHandler.Subscribe).Methods(http.MethodPost)
	api.HandleFunc("queues/{queue_name}/messages", msgHandler.ProduceMessage).Methods(http.MethodPost)
	api.HandleFunc("queues/{queue_name}/messages", msgHandler.FetchMessages).Methods(http.MethodGet)
	return r
}
