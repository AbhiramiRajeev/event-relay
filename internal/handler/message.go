package handler

import (
	"encoding/json"
	"net/http"

	model "github.com/AbhiramiRajeev/event-relay/internal/models"
	"github.com/AbhiramiRajeev/event-relay/internal/queue"
)

type MessageHandler struct{
	queue *queue.SQS
}

func NewMessageHandler(queue *queue.SQS) *MessageHandler {
	return &MessageHandler{
		queue:queue ,
	}
}

func (h *MessageHandler) Publish(w http.ResponseWriter, r *http.Request) {
	var message model.Message

	if err := json.NewDecoder(r.Body).Decode(&message); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	body, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "failed to encode message", http.StatusInternalServerError)
		return
	}

	if err := h.queue.Send(r.Context(), string(body)); err != nil {
		http.Error(w, "failed to queue message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
