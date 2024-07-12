package consumer

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/cenkalti/backoff/v4"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
)

//func logRetryAttempt(retryNum int, delay time.Duration) {
//	logger.Error("retry attempt", "retry", retryNum, "delay", delay)
//}

func SetupMiddleware() []message.HandlerMiddleware {
	return []message.HandlerMiddleware{
		AckOnMaxRetry{
			middleware.Retry{
				MaxRetries:      5,
				InitialInterval: 1 * time.Second,
				MaxInterval:     20 * time.Second,
				Multiplier:      3.0,
			},
		}.MiddlewareAckOnMaxRetry,
	}
}

// AckOnMaxRetry provides a middleware that retries the handler if errors are returned and acknowledges the message when max retries are reached.
type AckOnMaxRetry struct {
	middleware.Retry
}

// Middleware returns the AckOnMaxRetry middleware.
func (a AckOnMaxRetry) MiddlewareAckOnMaxRetry(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) ([]*message.Message, error) {
		producedMessages, err := h(msg)
		if err == nil {
			return producedMessages, nil
		}

		expBackoff := backoff.NewExponentialBackOff()
		expBackoff.InitialInterval = a.InitialInterval
		expBackoff.MaxInterval = a.MaxInterval
		expBackoff.Multiplier = a.Multiplier
		expBackoff.MaxElapsedTime = a.MaxElapsedTime
		expBackoff.RandomizationFactor = a.RandomizationFactor

		ctx := msg.Context()
		if a.MaxElapsedTime > 0 {
			var cancel func()
			ctx, cancel = context.WithTimeout(ctx, a.MaxElapsedTime)
			defer cancel()
		}

		retryNum := 1
		expBackoff.Reset()
	retryLoop:
		for {
			waitTime := expBackoff.NextBackOff()
			select {
			case <-ctx.Done():
				return producedMessages, err
			case <-time.After(waitTime):
				// go on
			}

			producedMessages, err = h(msg)
			if err == nil {
				return producedMessages, nil
			}

			if a.Logger != nil {
				a.Logger.Error("Error occurred, retrying", err, watermill.LogFields{
					"retry_no":     retryNum,
					"max_retries":  a.MaxRetries,
					"wait_time":    waitTime,
					"elapsed_time": expBackoff.GetElapsedTime(),
				})
			}
			if a.OnRetryHook != nil {
				a.OnRetryHook(retryNum, waitTime)
			}

			retryNum++
			if retryNum > a.MaxRetries {
				logger.Info("message automatically acked")
				msg.Ack()
				break retryLoop
			}
		}

		return nil, err
	}
}
