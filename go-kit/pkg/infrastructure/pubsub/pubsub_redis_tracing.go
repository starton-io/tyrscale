package pubsub

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	subscriberTracerName = "redis.subscriber"
	publisherTracerName  = "redis.publisher"
)

type TracingPublisherDecorator struct {
	tracerPub IPub
}

func NewTracingPublisherDecorator(tracerPub IPub) IPub {
	return &TracingPublisherDecorator{tracerPub: tracerPub}
}

func (t *TracingPublisherDecorator) Publish(ctx context.Context, topic string, messages ...*message.Message) error {
	_, span := tracer.Start(ctx, publisherTracerName, trace.WithSpanKind(trace.SpanKindProducer))
	defer tracer.SafeEndSpan(span)
	if len(messages) == 0 {
		return nil
	}

	messages[0].SetContext(ctx)
	spanAttributes := []attribute.KeyValue{
		semconv.MessagingDestinationKindTopic,
		semconv.MessagingDestinationKey.String(topic),
		semconv.MessagingOperationProcess,
	}
	spanAttributes = append(spanAttributes, spanAttributes...)
	span.SetAttributes(spanAttributes...)

	spanContext := trace.SpanContextFromContext(ctx)
	for _, msg := range messages {
		if spanContext.HasTraceID() {
			msg.Metadata.Set("trace_id", spanContext.TraceID().String())
		}
		if spanContext.HasSpanID() {
			msg.Metadata.Set("span_id", spanContext.SpanID().String())
		}
	}

	return t.tracerPub.Publish(ctx, topic, messages...)
}

func (t *TracingPublisherDecorator) Close() error {
	return t.tracerPub.Close()
}

// ExtractRemoteParentSpanContext defines a middleware that will extract trace/span id
// from the message metadata and creates a child span for the message.
func ExtractRemoteParentSpanContext() message.HandlerMiddleware {
	return func(h message.HandlerFunc) message.HandlerFunc {
		return ExtractRemoteParentSpanContextHandler(h)
	}
}

// ExtractRemoteParentSpanContextHandler decorates a watermill HandlerFunc to extract
// trace/span id from the metadata when a message is received and set a child span context.
func ExtractRemoteParentSpanContextHandler(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) ([]*message.Message, error) {

		if msg.Metadata.Get("trace_id") != "" {

			var traceID trace.TraceID
			var spanID trace.SpanID
			var err error

			traceID, err = trace.TraceIDFromHex(msg.Metadata.Get("trace_id"))
			if err != nil {
				return nil, err
			}

			if msg.Metadata.Get("span_id") != "" {
				spanID, err = trace.SpanIDFromHex(msg.Metadata.Get("span_id"))
				if err != nil {
					return nil, err
				}
			}

			spanContext := trace.NewSpanContext(
				trace.SpanContextConfig{
					TraceID:    traceID,
					SpanID:     spanID,
					TraceFlags: 01,
					Remote:     true,
				},
			)

			if spanContext.IsValid() {
				msg.SetContext(trace.ContextWithSpanContext(msg.Context(), spanContext))
			}
		}

		return h(msg)
	}
}

// config represents the configuration options available for subscriber
// middlewares and publisher decorators.
type configTracing struct {
	spanAttributes []attribute.KeyValue
	tracer         trace.Tracer
}

// Option provides a convenience wrapper for simple options that can be
type OptionTracing func(*configTracing)

// WithSpanAttributes includes the given attributes to the generated Spans.
func WithSpanAttributes(attributes ...attribute.KeyValue) OptionTracing {
	return func(c *configTracing) {
		c.spanAttributes = attributes
	}
}

func WithTracer(tracer trace.Tracer) OptionTracing {
	return func(c *configTracing) {
		c.tracer = tracer
	}
}

// Trace defines a middleware that will add tracing.
func Trace(options ...OptionTracing) message.HandlerMiddleware {
	return func(h message.HandlerFunc) message.HandlerFunc {
		return TraceHandler(h, options...)
	}
}

// TraceHandler decorates a watermill HandlerFunc to add tracing when a message is received.
func TraceHandler(h message.HandlerFunc, options ...OptionTracing) message.HandlerFunc {
	config := &configTracing{}

	for _, opt := range options {
		opt(config)
	}

	if config.tracer == nil {
		config.tracer = otel.Tracer(subscriberTracerName)
	}

	spanOptions := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindConsumer),
		trace.WithAttributes(config.spanAttributes...),
	}

	return func(msg *message.Message) ([]*message.Message, error) {
		spanName := message.HandlerNameFromCtx(msg.Context())
		ctx, span := config.tracer.Start(msg.Context(), spanName, spanOptions...)
		span.SetAttributes(
			semconv.MessagingDestinationKindTopic,
			semconv.MessagingDestinationKey.String(message.SubscribeTopicFromCtx(ctx)),
			semconv.MessagingOperationReceive,
		)
		msg.SetContext(ctx)

		events, err := h(msg)

		if err != nil {
			span.RecordError(err)
		}
		span.End()

		return events, err
	}
}

// TraceNoPublishHandler decorates a watermill NoPublishHandlerFunc to add tracing when a message is received.
func TraceNoPublishHandler(h message.NoPublishHandlerFunc, options ...OptionTracing) message.NoPublishHandlerFunc {
	decoratedHandler := TraceHandler(func(msg *message.Message) ([]*message.Message, error) {
		return nil, h(msg)
	}, options...)

	return func(msg *message.Message) error {
		_, err := decoratedHandler(msg)

		return err
	}
}
