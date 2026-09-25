package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/AbhiramiRajeev/event-relay/internal/handler"
	"github.com/AbhiramiRajeev/event-relay/internal/queue"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	queueURL := os.Getenv("SQS_QUEUE_URL")

	if queueURL == "" {
		log.Fatal("SQS_QUEUE_URL is not set")
	}

	messageQueue := queue.NewSQS(sqsClient, queueURL)
	messageHandler := handler.NewMessageHandler(messageQueue)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /messages", messageHandler.Publish)

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("event-relay listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
