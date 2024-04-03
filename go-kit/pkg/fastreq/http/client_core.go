package fastreq

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/starton-io/tyrscale/go-kit/pkg/fastreq/core"
	"github.com/valyala/fasthttp"
)

const (
	defaultMaxIdleConnections = 5
	defaultResponseTimeout    = 5 * time.Second
)

func (c *fastHttpClient) do(method string, url string, headers map[string]string, body interface{}) (*core.Response, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	c.mergeWithDefaultHeaders(headers)

	for key, value := range c.builder.headers {
		req.Header.Set(key, value)
	}
	contentType := req.Header.ContentType()
	reqBody, err := c.getRequestBody(string(contentType), body)
	if err != nil {
		return nil, err
	}
	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = c.GetClient().Do(req, resp)
	if err != nil {
		fmt.Printf("err: %v", err)
		return nil, err
	}
	return &core.Response{
		StatusCode: resp.StatusCode(),
		Body:       resp.Body(),
		Header:     convertHeaders(&resp.Header),
	}, nil
}

// set default headers json
func (c *fastHttpClient) mergeWithDefaultHeaders(headers map[string]string) {
	if c.builder.headers == nil {
		c.builder.headers = make(map[string]string)
	}
	for key, value := range headers {
		c.builder.headers[key] = value
	}
}

func (c *fastHttpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {

	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case ContentTypeJson:
		return json.Marshal(body)
	case ContentTypeXml:
		return xml.Marshal(body)
	case ContentTypePlainText:
		return body.([]byte), nil
	default:
		return json.Marshal(body)
	}
}

func (c *fastHttpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

func (c *fastHttpClient) GetClient() core.HttpClient {
	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		c.client = &fasthttp.Client{
			ReadTimeout:  c.getResponseTimeout(),
			WriteTimeout: c.getResponseTimeout(),
		}
	})
	return c.client
}

// convertHeaders converts fasthttp.ResponseHeader to a map[string][]string
func convertHeaders(header *fasthttp.ResponseHeader) map[string][]string {
	headers := make(map[string][]string)
	header.VisitAll(func(key, value []byte) {
		k := string(key)
		v := string(value)
		headers[k] = append(headers[k], v)
	})
	return headers
}
