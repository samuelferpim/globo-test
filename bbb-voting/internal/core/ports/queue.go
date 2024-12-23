package ports

import (
	"context"

	"github.com/streadway/amqp"
)

type QueueService interface {
	Publish(ctx context.Context, topic string, message []byte) error
	PublishError(ctx context.Context, originalMessage []byte, errorMsg string) error
	Close() error
}

type AMQPConnection interface {
	Channel() (AMQPChannel, error)
	Close() error
}

type AMQPChannel interface {
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	Close() error
}
