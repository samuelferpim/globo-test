package ports

import "context"

type QueueService interface {
	Publish(ctx context.Context, topic string, message []byte) error
	Close() error
}
