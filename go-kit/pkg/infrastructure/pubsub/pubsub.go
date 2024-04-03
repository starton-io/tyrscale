package pubsub

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

//go:generate mockery --name=IPub --output=./mocks
type IPub interface {
	Publish(ctx context.Context, topic string, message ...*message.Message) error
	Close() error
}

type ISub interface {
	Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error)
	Close() error
}
