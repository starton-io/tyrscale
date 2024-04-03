package interceptor

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

type DefaultRequestInterceptor struct {
	Host        string
	Path        string
	Port        int32
	Scheme      string
	ExtraHeader map[string]string
}

func (r *DefaultRequestInterceptor) Intercept(req *fasthttp.Request) error {
	req.URI().SetPath(r.Path)
	req.URI().SetHost(fmt.Sprintf("%s:%d", r.Host, r.Port))
	req.URI().SetScheme(r.Scheme)
	for k, v := range r.ExtraHeader {
		req.Header.Set(k, v)
	}
	return nil
}
