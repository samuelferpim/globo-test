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
	conn ports.AMQPConnection
}

func NewRabbitMQ(conn ports.AMQPConnection) ports.QueueService {
	return &RabbitMQ{conn: conn}
}

func NewRabbitMQConnection(url string) (ports.AMQPConnection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return &amqpConnectionWrapper{conn}, nil
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

	done := make(chan error)
	go func() {
		err := ch.Publish(
			"",         // exchange
			VotesQueue, // routing key
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        message,
				Headers: amqp.Table{
					"is_error": isError,
				},
			},
		)
		done <- err
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-done:
		return err
	}
}

func (r *RabbitMQ) Close() error {
	return r.conn.Close()
}

type amqpConnectionWrapper struct {
	*amqp.Connection
}

func (w *amqpConnectionWrapper) Channel() (ports.AMQPChannel, error) {
	ch, err := w.Connection.Channel()
	if err != nil {
		return nil, err
	}
	return &amqpChannelWrapper{ch}, nil
}

type amqpChannelWrapper struct {
	*amqp.Channel
}

func (w *amqpChannelWrapper) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	return w.Channel.Publish(exchange, key, mandatory, immediate, msg)
}
