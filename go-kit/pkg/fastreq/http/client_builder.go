package fastreq

import (
	"time"

	"github.com/valyala/fasthttp"
)

type clientBuilder struct {
	headers           map[string]string
	connectionTimeout time.Duration
	responseTimeout   time.Duration
	maxConnections    int
	disableTimeouts   bool
	client            *fasthttp.Client
	url               string
}

type IClientBuilder interface {
	SetHeaders(headers map[string]string) IClientBuilder
	SetUrl(url string) IClientBuilder
	SetConnectionTimeout(timeout time.Duration) IClientBuilder
	SetResponseTimeout(timeout time.Duration) IClientBuilder
	SetMaxConnections(maxConn int) IClientBuilder
	DisableTimeouts(disable bool) IClientBuilder
	SetUserAgent(userAgent string) IClientBuilder
	Build() IClient
}

func NewBuilder() IClientBuilder {
	return &clientBuilder{}
}

func (c *clientBuilder) Build() IClient {
	client := &fastHttpClient{
		builder: c,
	}
	return client
}

func (c *clientBuilder) SetHeaders(headers map[string]string) IClientBuilder {
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	for k, v := range headers {
		c.headers[k] = v
	}
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) IClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) IClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxConnections(i int) IClientBuilder {
	c.maxConnections = i
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) IClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHttpClient(client *fasthttp.Client) IClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(userAgent string) IClientBuilder {
	if c.headers == nil {
		c.headers = make(map[string]string)
	}
	c.headers["User-Agent"] = userAgent
	return c
}

func (c *clientBuilder) SetUrl(url string) IClientBuilder {
	c.url = url
	return c
}
