package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
	"go.uber.org/zap/zaptest/observer"
)

func TestApply(t *testing.T) {
	var order []string

	m1 := func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			order = append(order, "m1")
			next(ctx)
		}
	}

	m2 := func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return func(ctx *fasthttp.RequestCtx) {
			order = append(order, "m2")
			next(ctx)
		}
	}

	finalHandler := func(ctx *fasthttp.RequestCtx) {
		order = append(order, "final")
	}

	handler := Apply(finalHandler, m2, m1)
	ctx := &fasthttp.RequestCtx{}
	handler(ctx)

	assert.Equal(t, []string{"m1", "m2", "final"}, order)
}

func TestChainMiddlewareWithPriority(t *testing.T) {
	var order []string

	m2 := &MiddlewareWithPriority{
		Middleware: func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
			return func(ctx *fasthttp.RequestCtx) {
				order = append(order, "m2")
				t.Log("append m2")
				next(ctx)
			}
		},
		Name:     "m2",
		Priority: 100,
	}

	m1 := &MiddlewareWithPriority{
		Middleware: func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
			return func(ctx *fasthttp.RequestCtx) {
				order = append(order, "m1")
				t.Log("append m1")
				next(ctx)
			}
		},
		Name:     "m1",
		Priority: 200,
	}

	m3 := &MiddlewareWithPriority{
		Middleware: func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
			return func(ctx *fasthttp.RequestCtx) {
				order = append(order, "m3")
				t.Log("append m3")
				next(ctx)
			}
		},
		Name:     "m3",
		Priority: 150,
	}

	finalHandler := func(ctx *fasthttp.RequestCtx) {
		order = append(order, "final")
		t.Log("append final")
	}

	//handler := ChainMiddlewareWithPriority(finalHandler, []*MiddlewareWithPriority{m1, m2, m3})
	composerMiddleware := MiddlewareWithPriorityComposer(m1, m2, m3)
	handler := composerMiddleware(finalHandler)
	ctx := &fasthttp.RequestCtx{}
	handler(ctx)

	assert.Equal(t, []string{"m1", "m3", "m2", "final"}, order)
}

func TestNewLogger(t *testing.T) {
	// Create a zap test logger
	core, recorded := observer.New(zapcore.InfoLevel)
	logger := zap.New(core)

	// Create the middleware
	middleware := NewLogger(logger)

	// Create a dummy handler
	next := func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
	}

	// Create a fasthttp request context
	ctx := &fasthttp.RequestCtx{}

	// Call the middleware
	handler := middleware(next)
	handler(ctx)

	// Check the logs
	logs := recorded.All()
	if len(logs) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(logs))
	}

	log := logs[0]
	if log.Level != zapcore.InfoLevel {
		t.Errorf("expected log level %v, got %v", zapcore.InfoLevel, log.Level)
	}

	if log.Message != "access" {
		t.Errorf("expected log message 'access', got %v", log.Message)
	}

	fields := log.ContextMap()
	if fields["code"] != int64(fasthttp.StatusOK) {
		t.Errorf("expected status code %v, got %v", fasthttp.StatusOK, fields["code"])
	}

	if _, ok := fields["time"]; !ok {
		t.Errorf("expected 'time' field in log")
	}

	if _, ok := fields["host"]; !ok {
		t.Errorf("expected 'host' field in log")
	}

	if _, ok := fields["method"]; !ok {
		t.Errorf("expected 'method' field in log")
	}
}

func TestNewCors(t *testing.T) {
	// Create a dummy handler to pass to the middleware
	dummyHandler := func(ctx *fasthttp.RequestCtx) {
		ctx.SetStatusCode(fasthttp.StatusOK)
	}

	// Create the CORS middleware
	corsMiddleware := NewCors()

	// Wrap the dummy handler with the CORS middleware
	handler := corsMiddleware(dummyHandler)

	// Create a new fasthttp request context
	ctx := &fasthttp.RequestCtx{}

	// Call the handler
	handler(ctx)

	// Check the response headers
	assert.Equal(t, "*", string(ctx.Response.Header.Peek("Access-Control-Allow-Origin")))
	assert.Equal(t, "POST, OPTIONS", string(ctx.Response.Header.Peek("Access-Control-Allow-Methods")))
	assert.Equal(t, "Content-Type", string(ctx.Response.Header.Peek("Access-Control-Allow-Headers")))
}

func TestNewRecover(t *testing.T) {
	logger := zaptest.NewLogger(t)
	middleware := NewRecover(logger)

	handler := middleware(func(ctx *fasthttp.RequestCtx) {
		panic("test panic")
	})

	ctx := &fasthttp.RequestCtx{}
	handler(ctx)

	// Check logs for the expected panic message
	// This part is more complex and might require a custom zapcore to capture logs
}
