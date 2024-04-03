package tracer

import (
	"context"
	"net/url"
	"runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer

func InitTracer(name string) trace.Tracer {
	if Tracer != nil {
		return Tracer
	}
	Tracer = otel.Tracer(name)
	return Tracer
}

func GetTracer() trace.Tracer {
	return Tracer
}

func Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	// check if tracer if empty
	if Tracer == nil {
		return ctx, nil
	}
	pc, file, line, _ := runtime.Caller(1) // Caller(1) to get the function that called StartTrace
	funcName := runtime.FuncForPC(pc).Name()
	opts = append(opts, trace.WithAttributes(attribute.String("code.function", funcName)))
	opts = append(opts, trace.WithAttributes(attribute.String("code.filepath", file)))
	opts = append(opts, trace.WithAttributes(attribute.Int("code.lineno", line)))

	return Tracer.Start(ctx, spanName, opts...)
}

func Trace(ctx *context.Context, spanName string, opts ...trace.SpanStartOption) func() {
	if Tracer == nil {
		return func() {}
	}
	c, span := Tracer.Start(*ctx, spanName, opts...)
	*ctx = c
	return func() {
		span.End()
	}
}

func SafeEndSpan(span trace.Span) {
	if span != nil {
		span.End()
	}
}

type OptTracer struct {
	Endpoint        string
	TracingProvider *sdktrace.TracerProvider
	ServiceInstance string
}

func InitTracing(serviceName string, opts *OptTracer) *sdktrace.TracerProvider {
	endpoint := opts.Endpoint
	var u *url.URL
	var err error
	tp := opts.TracingProvider
	// get the scheme from the url
	if endpoint != "" {
		u, err = url.Parse(endpoint)
		if err != nil {
			panic(err)
		}

	}
	var serviceInstance string
	if opts.ServiceInstance != "" {
		serviceInstance = opts.ServiceInstance
	}
	var exporter sdktrace.SpanExporter
	switch {
	case u == nil:
		exporter, err = otlptracehttp.New(context.Background(), otlptracehttp.WithInsecure())
		if err != nil {
			panic(err)
		}
	case u.Scheme == "http":
		exporter, err = otlptracehttp.New(context.Background(), otlptracehttp.WithEndpointURL(endpoint), otlptracehttp.WithInsecure())
		if err != nil {
			panic(err)
		}

	case u.Scheme == "grpc":
		exporter, err = otlptracegrpc.New(context.Background(), otlptracegrpc.WithEndpoint(endpoint), otlptracegrpc.WithInsecure())
		if err != nil {
			panic(err)
		}
	}

	if tp == nil {
		tp = DefaultTracerProvider(serviceName, serviceInstance, exporter)
	}
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	Tracer = otel.Tracer(serviceName)
	return tp
}

func DefaultTracerProvider(serviceName string, serviceInstance string, exporter sdktrace.SpanExporter) *sdktrace.TracerProvider {
	var attrs []attribute.KeyValue
	attrs = append(attrs, semconv.ServiceNameKey.String(serviceName))
	if serviceInstance != "" {
		attrs = append(attrs, semconv.ServiceInstanceIDKey.String(serviceInstance))
	}

	return sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				attrs...,
			)),
	)
}
