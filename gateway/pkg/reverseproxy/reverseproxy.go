package reverseproxy

import (
	"net"

	"github.com/starton-io/tyrscale/gateway/pkg/handler"
	"github.com/valyala/fasthttp"
)

var hopHeaders = [][]byte{
	[]byte("Connection"),
	[]byte("Proxy-Connection"), // non-standard but still sent by libcurl and rejected by e.g. google
	[]byte("Keep-Alive"),
	[]byte("Proxy-Authenticate"),
	[]byte("Proxy-Authorization"),
	[]byte("Te"),      // canonicalized version of "TE"
	[]byte("Trailer"), // not Trailers per URL above; https://www.rfc-editor.org/errata_search.php?eid=4522
	[]byte("Transfer-Encoding"),
	[]byte("Upgrade"),
}

type ReverseProxyHandler struct {
	ReverseProxy handler.ProxyHandler
}

type ProxyHandler interface {
	ReverseProxyHandler(ctx *fasthttp.RequestCtx)
}

func NewReverseProxyHandler(handler handler.ProxyHandler) ProxyHandler {
	return &ReverseProxyHandler{
		ReverseProxy: handler,
	}
}

func (p *ReverseProxyHandler) ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	res := &ctx.Response

	if ip, _, err := net.SplitHostPort(ctx.RemoteAddr().String()); err == nil {
		req.Header.Add("X-Forwarded-For", ip)
	}

	// remove hop-by-hop headers
	for _, hopHeader := range hopHeaders {
		req.Header.DelBytes(hopHeader)
	}

	p.ReverseProxy.Handle(ctx)

	// remove hop-by-hop headers
	for _, hopHeader := range hopHeaders {
		res.Header.DelBytes(hopHeader)
	}
}
