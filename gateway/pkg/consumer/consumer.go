package consumer

import (
	"context"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
)

const (
	RouteCreated     = "route_created"
	RouteDeleted     = "route_deleted"
	UpstreamUpserted = "upstream_upserted"
	UpstreamDeleted  = "upstream_deleted"
)

type IConsumerControl interface {
	Pause()
	Resume()
}

type IConsumerReset interface {
	Reset() error
}

type IConsumer interface {
	IConsumerControl
	Start(ctx context.Context)
	Shutdown(ctx context.Context) error
}

type Consumer struct {
	RouteHandler *RouteHandler
	Subscriber   pubsub.ISub
	Router       *message.Router
	Suspend      bool
}

func NewConsumer(routeHandler *RouteHandler, subscriber pubsub.ISub, router *message.Router) IConsumer {
	return &Consumer{
		RouteHandler: routeHandler,
		Subscriber:   subscriber,
		Router:       router,
		Suspend:      true,
	}
}

func (c *Consumer) Pause() {
	c.Suspend = true
}

func (c *Consumer) Resume() {
	c.Suspend = false
}

func (c *Consumer) Start(ctx context.Context) {
	logger.Info("Starting consumer")

	c.Router.AddMiddleware(SetupMiddleware()...)
	c.Router.AddNoPublisherHandler(RouteCreated, RouteCreated,
		c.Subscriber, func(msg *message.Message) error {

			logger.Infof("Adding route: %s", string(msg.Payload))
			if c.Suspend {
				return fmt.Errorf("consumer is suspended")
			}
			return c.RouteHandler.HandleRouteCreated(msg)
		})
	c.Router.AddNoPublisherHandler(RouteDeleted, RouteDeleted,
		c.Subscriber,
		func(msg *message.Message) error {
			if c.Suspend {
				return fmt.Errorf("consumer is suspended")
			}
			return c.RouteHandler.HandleRouteDeleted(msg)
		},
	)
	c.Router.AddNoPublisherHandler(UpstreamUpserted, UpstreamUpserted,
		c.Subscriber,
		func(msg *message.Message) error {
			if c.Suspend {
				return fmt.Errorf("consumer is suspended")
			}
			return c.RouteHandler.HandleUpstreamUpserted(msg)
		},
	)
	c.Router.AddNoPublisherHandler(UpstreamDeleted, UpstreamDeleted,
		c.Subscriber,
		func(msg *message.Message) error {
			if c.Suspend {
				return fmt.Errorf("consumer is suspended")
			}
			return c.RouteHandler.HandleUpstreamDeleted(msg)
		},
	)

	go func() {
		if err := c.Router.Run(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
}

func (c *Consumer) Shutdown(ctx context.Context) error {
	return c.Router.Close()
}
