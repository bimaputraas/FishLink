package publisher

import (
	"context"
)

type Publisher interface {
	PublishMessage(ctx context.Context,queueName string, msg []byte) error
}