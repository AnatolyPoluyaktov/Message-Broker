package http

import "net/http"

type MessageHandler struct {
}

func (mh *MessageHandler) ProduceMessage(w http.ResponseWriter, r *http.Request) {

}

func (mh *MessageHandler) Subscribe(w http.ResponseWriter, r *http.Request) {

}

func (mh *MessageHandler) FetchMessages(w http.ResponseWriter, r *http.Request) {

}
