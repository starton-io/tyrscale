package interceptor

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type mockInterceptorResponse struct {
	called bool
}

func (m *mockInterceptorResponse) Intercept(res *fasthttp.Response) error {
	m.called = true
	return nil
}

type mockInterceptorRequest struct {
	called bool
}

func (m *mockInterceptorRequest) Intercept(req *fasthttp.Request) error {
	m.called = true
	return nil
}

func TestInterceptorResponseChain(t *testing.T) {
	// init fake server
	logger.InitLogger()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	chainResp := NewInterceptorResponseChain()
	chainReq := NewInterceptorRequestChain()

	// Test AddFirst
	firstInterceptor := &mockInterceptorResponse{}
	firstInterceptorReq := &mockInterceptorRequest{}
	chainResp.AddFirst(firstInterceptor)
	chainReq.AddFirst(firstInterceptorReq)
	assert.Equal(t, 1, len(chainResp.interceptors))
	assert.Equal(t, 1, len(chainReq.interceptors))
	assert.Equal(t, firstInterceptor, chainResp.interceptors[0].interceptor)
	assert.Equal(t, firstInterceptorReq, chainReq.interceptors[0].interceptor)

	// Test AddLast
	lastInterceptor := &mockInterceptorResponse{}
	lastInterceptorReq := &mockInterceptorRequest{}
	chainResp.AddLast(lastInterceptor)
	chainReq.AddLast(lastInterceptorReq)
	assert.Equal(t, 2, len(chainResp.interceptors))
	assert.Equal(t, lastInterceptor, chainResp.interceptors[1].interceptor)
	assert.Equal(t, 2, len(chainReq.interceptors))
	assert.Equal(t, lastInterceptorReq, chainReq.interceptors[1].interceptor)

	// Test AddOrdered
	orderedInterceptor := &DefaultResponseInterceptor{}
	parsedURL, _ := url.Parse(server.URL)
	port, _ := strconv.Atoi(parsedURL.Port())
	orderedInterceptorReq := &DefaultRequestInterceptor{
		Host:   parsedURL.Host,
		Path:   "/",
		Port:   int32(port),
		Scheme: parsedURL.Scheme,
	}
	chainResp.AddOrdered(orderedInterceptor, 4)
	chainReq.AddOrdered(orderedInterceptorReq, 4)
	assert.Equal(t, 3, len(chainResp.interceptors))
	assert.Equal(t, orderedInterceptor, chainResp.interceptors[2].interceptor)

	// make request
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(server.URL)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := chainReq.Intercept(req)
	assert.NoError(t, err)

	// make request
	fasthttp.Do(req, resp)

	err = chainResp.Intercept(resp)
	assert.NoError(t, err)
	assert.True(t, firstInterceptor.called)
	assert.True(t, lastInterceptor.called)

	// Test Remove
	chainResp.Remove(orderedInterceptor)
	assert.Equal(t, 2, len(chainResp.interceptors))
	assert.Equal(t, firstInterceptor, chainResp.interceptors[0].interceptor)
	assert.Equal(t, lastInterceptor, chainResp.interceptors[1].interceptor)
}
