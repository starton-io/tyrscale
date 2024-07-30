package main

import (
	"encoding/json"

	"github.com/starton-io/tyrscale/gateway/pkg/interceptor/types"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/valyala/fasthttp"
)

var validate validation.Validation

func init() {
	validate = validation.New()
}

type Config struct {
	Headers map[string]string `json:"headers" validate:"required"`
}

type MyResponseInterceptor struct {
	Config Config
}

func (m *MyResponseInterceptor) Intercept(response *fasthttp.Response) error {
	response.Header.Set("X-Intercepted", "true")
	for k, v := range m.Config.Headers {
		response.Header.Set(k, v)
	}
	return nil
}

type MyResponseInterceptorRegister struct{}

func (p *MyResponseInterceptorRegister) RegisterResponseInterceptor(registerFunc func(name string, interceptor types.IResponseInterceptor), payload []byte) error {
	var config Config
	if err := json.Unmarshal(payload, &config); err != nil {
		return err
	}
	if err := validate.ValidateStruct(config); err != nil {
		return err
	}
	registerFunc("MyResponseInterceptor", &MyResponseInterceptor{Config: config})
	return nil
}

func (p *MyResponseInterceptorRegister) Validate(configPayload []byte) error {
	var config Config
	if err := json.Unmarshal(configPayload, &config); err != nil {
		return err
	}
	return validate.ValidateStruct(config)
}

var ResponseInterceptor MyResponseInterceptorRegister

func main() {}
