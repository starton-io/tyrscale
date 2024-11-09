package interceptor

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"

	"github.com/starton-io/tyrscale/gateway/pkg/normalizer"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/valyala/fasthttp"
)

type DefaultResponseInterceptor struct {
}

func (r *DefaultResponseInterceptor) Intercept(res *fasthttp.Response) error {

	// Check if the response is gzipped
	var body []byte
	if bytes.Equal(res.Header.Peek("Content-Encoding"), []byte("gzip")) {
		// Read the entire response body
		gunzipBody, err := gzip.NewReader(bytes.NewReader(res.Body()))
		if err != nil {
			return err
		}
		defer gunzipBody.Close()
		body, err = io.ReadAll(gunzipBody)
		if err != nil {
			return err
		}
	} else {
		body = res.Body()
	}

	// Parse the response body
	normalizedResponse := &normalizer.NormalizedResponse{}
	if err := normalizedResponse.ParseResponse(body); err != nil {
		logger.Errorf("Failed to parse response: %v", err)
		return err
	}

	// Check for specific error codes in the response
	if normalizedResponse.RpcResponse.Error != nil {
		if normalizedResponse.RpcResponse.Error.Code == -32601 {
			logger.Debugf("Status 405, ignoring method %s", normalizedResponse.RpcResponse.Error.Message)
			res.SetStatusCode(fasthttp.StatusMethodNotAllowed)
			res.SetBody(body)
			return nil
		}
	}

	// Set the response body
	res.SetBody(body)

	// Check for specific HTTP status codes
	switch res.StatusCode() {
	case fasthttp.StatusTooManyRequests:
		return errors.New("too many requests")
	case fasthttp.StatusInternalServerError, fasthttp.StatusBadGateway, fasthttp.StatusServiceUnavailable, fasthttp.StatusGatewayTimeout:
		return errors.New("invalid response")
	}

	return nil
}

// TODO: implement default last response interceptor GZIP
type DefaultLastResponseInterceptor struct {
}

func gZipData(data []byte) (compressedData []byte, err error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	_, err = gz.Write(data)
	if err != nil {
		return
	}
	if err = gz.Flush(); err != nil {
		return
	}
	if err = gz.Close(); err != nil {
		return
	}
	compressedData = b.Bytes()
	return
}

func (r *DefaultLastResponseInterceptor) Intercept(res *fasthttp.Response) error {
	//gzip the body
	body, err := gZipData(res.Body())
	if err != nil {
		return err
	}
	res.SetBody(body)
	return nil
}
