package core

import (
	"github.com/valyala/fasthttp"
)

type HttpClient interface {
	Do(req *fasthttp.Request, resp *fasthttp.Response) error
}
