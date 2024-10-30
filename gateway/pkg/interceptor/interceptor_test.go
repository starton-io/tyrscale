package interceptor

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/starton-io/tyrscale/gateway/pkg/normalizer"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
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
	_ = logger.InitLogger()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a valid JSON-RPC response
		response := &normalizer.RPCResponse{
			JsonrpcVersion: "2.0",
			Result:         "OK", // or nil if you want to simulate an error
			Error:          nil,  // or provide an error struct if simulating an error
			ID:             1,
		}

		// Marshal the response to JSON
		responseBody, err := normalizer.SonicCfg.Marshal(response)
		if err != nil {
			http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
			return
		}

		// Check if the request should return a gzip response
		if r.Header.Get("Accept-Encoding") == "gzip" {
			// encode response body with gzip
			gzipResponseBody, err := gZipData(responseBody)
			if err != nil {
				http.Error(w, "Failed to gzip response", http.StatusInternalServerError)
				return
			}
			// Set response headers for gzip
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Encoding", "gzip")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(gzipResponseBody)
		} else {
			// Set response headers for plain JSON
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(responseBody)
		}
	}))
	defer server.Close()

	chainResp := NewInterceptorResponseChain()
	chainReq := NewInterceptorRequestChain()

	// Test AddFirst
	firstInterceptor := &mockInterceptorResponse{}
	firstInterceptorReq := &mockInterceptorRequest{}
	chainResp.AddFirst(firstInterceptor, "interceptorFirst")
	chainReq.AddFirst(firstInterceptorReq, "interceptorFirst")
	assert.Equal(t, 1, len(chainResp.interceptors))
	assert.Equal(t, 1, len(chainReq.interceptors))
	assert.Equal(t, firstInterceptor, chainResp.interceptors[0].interceptor)
	assert.Equal(t, firstInterceptorReq, chainReq.interceptors[0].interceptor)

	// Test AddLast
	lastInterceptor := &mockInterceptorResponse{}
	lastInterceptorReq := &mockInterceptorRequest{}
	chainResp.AddLast(lastInterceptor, "interceptorLast")
	chainReq.AddLast(lastInterceptorReq, "interceptorLast")
	assert.Equal(t, 2, len(chainResp.interceptors))
	assert.Equal(t, lastInterceptor, chainResp.interceptors[1].interceptor)
	assert.Equal(t, 2, len(chainReq.interceptors))
	assert.Equal(t, lastInterceptorReq, chainReq.interceptors[1].interceptor)

	// Test AddOrdered
	orderedInterceptor := &DefaultResponseInterceptor{}
	parsedURL, _ := url.Parse(server.URL)
	port, _ := strconv.Atoi(parsedURL.Port())
	orderedInterceptorReq := &DefaultRequestInterceptor{
		Host:   parsedURL.Hostname(),
		Path:   "/",
		Port:   int32(port),
		Scheme: parsedURL.Scheme,
	}
	chainResp.AddOrdered(orderedInterceptor, "interceptorOrdered", 4)
	chainReq.AddOrdered(orderedInterceptorReq, "interceptorOrdered", 4)
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
	err = fasthttp.Do(req, resp)
	assert.NoError(t, err)

	err = chainResp.Intercept(resp)
	assert.NoError(t, err)
	assert.True(t, firstInterceptor.called)
	assert.True(t, lastInterceptor.called)

	// make request with gzip
	req.Header.Set("Accept-Encoding", "gzip")
	err = fasthttp.Do(req, resp)
	assert.NoError(t, err)

	err = chainResp.Intercept(resp)
	assert.NoError(t, err)
	assert.True(t, firstInterceptor.called)
	assert.True(t, lastInterceptor.called)

	// Test Remove
	chainResp.Remove("interceptorOrdered")

	assert.Equal(t, 2, len(chainResp.interceptors))
	assert.Equal(t, firstInterceptor, chainResp.interceptors[0].interceptor)
	assert.Equal(t, lastInterceptor, chainResp.interceptors[1].interceptor)
}
