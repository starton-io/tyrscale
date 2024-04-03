package fastreq

import (
	"fmt"
	"sync"

	"github.com/starton-io/tyrscale/go-kit/pkg/fastreq/core"
	"github.com/valyala/fasthttp"
)

type fastHttpClient struct {
	builder    *clientBuilder
	client     *fasthttp.Client
	clientOnce sync.Once
}

type IClient interface {
	Do(url, method string, headers map[string]string, body interface{}) (*core.Response, error)
	Get(url string, headers map[string]string) (*core.Response, error)
	Post(url string, body interface{}, headers map[string]string) (*core.Response, error)
	Put(url string, body interface{}, headers map[string]string) (*core.Response, error)
	Patch(url string, body interface{}, headers map[string]string) (*core.Response, error)
	Delete(url string, headers map[string]string) (*core.Response, error)
	Options(url string, headers map[string]string) (*core.Response, error)
}

func (c *fastHttpClient) Do(url, method string, headers map[string]string, body interface{}) (*core.Response, error) {
	switch method {
	case "GET":
		return c.Get(url, headers)
	case "POST":
		return c.Post(url, body, headers)
	case "PUT":
		return c.Put(url, body, headers)
	case "PATCH":
		return c.Patch(url, body, headers)
	case "DELETE":
		return c.Delete(url, headers)
	case "OPTIONS":
		return c.Options(url, headers)
	default:
		return nil, fmt.Errorf("invalid method")
	}
}

func (c *fastHttpClient) Get(url string, headers map[string]string) (*core.Response, error) {
	return c.do("GET", url, headers, nil)
}

func (c *fastHttpClient) Post(url string, body interface{}, headers map[string]string) (*core.Response, error) {
	return c.do("POST", url, headers, body)
}

func (c *fastHttpClient) Put(url string, body interface{}, headers map[string]string) (*core.Response, error) {
	return c.do("PUT", url, headers, body)
}

func (c *fastHttpClient) Patch(url string, body interface{}, headers map[string]string) (*core.Response, error) {
	return c.do("PATCH", url, headers, body)
}

func (c *fastHttpClient) Delete(url string, headers map[string]string) (*core.Response, error) {
	return c.do("DELETE", url, headers, nil)
}

func (c *fastHttpClient) Options(url string, headers map[string]string) (*core.Response, error) {
	return c.do("OPTIONS", url, headers, nil)
}
