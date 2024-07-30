package middleware

import (
	"sort"

	"github.com/starton-io/tyrscale/gateway/pkg/middleware/types"
	"github.com/valyala/fasthttp"
)

// Middleware chainable fasthttp middleware func
type Middleware func(next fasthttp.RequestHandler) fasthttp.RequestHandler

// Apply process multiple middleware
func Apply(h fasthttp.RequestHandler, m ...types.MiddlewareFunc) fasthttp.RequestHandler {
	for i := 0; i < len(m); i++ {
		h = m[i](h)
	}
	return h
}

// ComposeMiddleware compose multiple middleware
func ComposeMiddleware(m ...types.MiddlewareFunc) types.MiddlewareFunc {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return Apply(next, m...)
	}
}

// middlewarecomposer with priority
func MiddlewareComposerWithPriority(middlewares []*MiddlewareWithPriority) types.MiddlewareFunc {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return ChainMiddlewareWithPriority(next, middlewares)
	}
}

// MiddlewareWithPriority holds a middleware function and its priority.
type MiddlewareWithPriority struct {
	Middleware types.MiddlewareFunc
	Name       string
	Priority   int
}

func ChainMiddlewareWithPriority(h fasthttp.RequestHandler, listMiddleware []*MiddlewareWithPriority) fasthttp.RequestHandler {
	// Sort middleware by priority in ascending order
	sort.SliceStable(listMiddleware, func(i, j int) bool {
		return listMiddleware[i].Priority < listMiddleware[j].Priority
	})
	// Apply middleware in reverse order to maintain priority
	for i := len(listMiddleware) - 1; i >= 0; i-- {
		h = listMiddleware[i].Middleware(h)
	}
	return h
}

func MiddlewareWithPriorityComposer(m ...*MiddlewareWithPriority) types.MiddlewareFunc {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return ChainMiddlewareWithPriority(next, m)
	}
}
