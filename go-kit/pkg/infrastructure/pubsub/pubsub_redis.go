package pubsub

import (
	"context"

	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
)

//go:generate mockery --name=IPubRedis --output=./mocks
type IPubRedis interface {
	Publish(topic string, msg ...*message.Message) error
	Close() error
}

type RedisPub struct {
	// contains filtered or unexported fields
	globalPrefix string
	red          IPubRedis
}

func NewRedisPub(config redisstream.PublisherConfig, opts ...Option) IPub {
	redisStream, _ := redisstream.NewPublisher(config, nil)
	r := &RedisPub{red: redisStream}
	for _, opt := range opts {
		r.withOption(opt)
	}
	return r
}

func (r *RedisPub) withOption(opt Option) {
	switch opt.Type() {
	case OptionTypeGlobalPrefix:
		r.globalPrefix = opt.(globalPrefix).prefix
	}
}

func (r *RedisPub) Publish(ctx context.Context, topic string, msg ...*message.Message) error {
	topic = r.globalPrefix + topic
	return r.red.Publish(topic, msg...)
}

func (r *RedisPub) Close() error {
	return r.red.Close()
}

type RedisSub struct {
	globalPrefix string
	red          *redisstream.Subscriber
}

func NewRedisSub(config redisstream.SubscriberConfig, opts ...Option) *RedisSub {
	redisStream, _ := redisstream.NewSubscriber(config, nil)
	r := &RedisSub{red: redisStream}
	for _, opt := range opts {
		r.withOption(opt)
	}
	return r
}

func (r *RedisSub) withOption(opt Option) {
	switch opt.Type() {
	case OptionTypeGlobalPrefix:
		r.globalPrefix = opt.(globalPrefix).prefix
	}
}

func (r *RedisSub) Subscribe(ctx context.Context, topic string) (<-chan *message.Message, error) {
	topic = r.globalPrefix + topic
	return r.red.Subscribe(ctx, topic)
}

func (r *RedisSub) Close() error {
	return r.red.Close()
}
