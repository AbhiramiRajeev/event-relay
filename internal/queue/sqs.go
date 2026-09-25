package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQS struct {
	client   *sqs.Client
	queueURL string
}

func NewSQS(client *sqs.Client, queueURL string) *SQS {
	return &SQS{
		client:   client,
		queueURL: queueURL,
	}
}

func (q *SQS) Send(ctx context.Context, body string) error {
	_, err := q.client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    &q.queueURL,
		MessageBody: &body,
	})

	return err
}	