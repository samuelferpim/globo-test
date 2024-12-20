package queue

import (
	"bbb-voting/internal/core/ports"
	"context"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

const (
	VotesQueue = "votes"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ(url string) (ports.QueueService, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{conn: conn}, nil
}

func (r *RabbitMQ) Publish(ctx context.Context, topic string, message []byte) error {
	if topic != VotesQueue {
		return fmt.Errorf("invalid topic: %s, only 'votes' is supported", topic)
	}

	return r.publishToQueue(ctx, message, false)
}

func (r *RabbitMQ) PublishError(ctx context.Context, originalMessage []byte, errorMsg string) error {
	errorPayload := struct {
		OriginalMessage json.RawMessage `json:"original_message"`
		Error           string          `json:"error"`
		IsError         bool            `json:"is_error"`
	}{
		OriginalMessage: originalMessage,
		Error:           errorMsg,
		IsError:         true,
	}

	errorJSON, err := json.Marshal(errorPayload)
	if err != nil {
		return fmt.Errorf("failed to marshal error payload: %w", err)
	}

	return r.publishToQueue(ctx, errorJSON, true)
}

func (r *RabbitMQ) publishToQueue(ctx context.Context, message []byte, isError bool) error {
	ch, err := r.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open channel: %w", err)
	}
	defer ch.Close()

	// Use a channel to signal when the publish is done
	done := make(chan error, 1)

	go func() {
		err := ch.Publish(
			"",         // exchange (use default exchange)
			VotesQueue, // routing key (queue name for default exchange)
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        message,
				Headers: amqp.Table{
					"is_error": isError,
				},
			})
		done <- err
	}()

	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("failed to publish message: %w", err)
		}
		return nil
	case <-ctx.Done():
		return fmt.Errorf("publish operation cancelled: %w", ctx.Err())
	}
}

func (r *RabbitMQ) Close() error {
	return r.conn.Close()
}
