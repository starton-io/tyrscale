package middleware

import "github.com/valyala/fasthttp"

// Middleware chainable fasthttp middleware func
type Middleware func(next fasthttp.RequestHandler) fasthttp.RequestHandler

// Apply process multiple middleware
func Apply(h fasthttp.RequestHandler, m ...Middleware) fasthttp.RequestHandler {
	for i := 0; i < len(m); i++ {
		h = m[i](h)
	}
	return h
}

// ComposeMiddleware compose multiple middleware
func ComposeMiddleware(m ...Middleware) Middleware {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return Apply(next, m...)
	}
}

// MiddlewareFunc defines a function type for middleware operations.
type MiddlewareFunc func(next fasthttp.RequestHandler) fasthttp.RequestHandler

// ChainMiddleware applies a sequence of pre-middlewares and post-middlewares to a handler.
func ChainMiddleware(h fasthttp.RequestHandler, listMiddleware []MiddlewareFunc) fasthttp.RequestHandler {
	// Apply pre-middlewares in the order provided.
	for _, preMiddleware := range listMiddleware {
		h = preMiddleware(h)
	}
	return h
}

// MiddlewareComposer returns a MiddlewareFunc that chains pre and post middlewares around a handler.
func MiddlewareComposer(pre []MiddlewareFunc) MiddlewareFunc {
	return func(next fasthttp.RequestHandler) fasthttp.RequestHandler {
		return ChainMiddleware(next, pre)
	}
}
