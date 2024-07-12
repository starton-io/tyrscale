package interceptor

import (
	"log"
	"path/filepath"
	"plugin"
)

// InitRespInterceptor is the chain of response interceptors
// InitReqInterceptor is the chain of request interceptors
var (
	InitRespInterceptor InterceptorResponseChain = InterceptorResponseChain{}
	InitReqInterceptor  InterceptorRequestChain  = InterceptorRequestChain{}
)

func NewInterceptorResponseChainFromFile(pluginPath string) *InterceptorResponseChain {
	plugin, err := plugin.Open(pluginPath)
	if err != nil {
		log.Fatalf("Error opening plugin: %v", err)
	}
	interceptorPlugin, err := plugin.Lookup("IResponseInterceptor")
	if err != nil {
		log.Fatal(err)
	}
	interceptor := interceptorPlugin.(IResponseInterceptor)
	return &InterceptorResponseChain{
		interceptors: []IResponseInterceptor{interceptor},
	}
}

func NewInterceptorRequestChainFromFile(pluginPath string) *InterceptorRequestChain {
	plugin, err := plugin.Open(pluginPath)
	if err != nil {
		log.Fatalf("Error opening plugin: %v", err)
	}
	interceptorPlugin, err := plugin.Lookup("IRequestInterceptor")
	if err != nil {
		log.Fatal(err)
	}
	interceptor := interceptorPlugin.(IRequestInterceptor)
	return &InterceptorRequestChain{
		interceptors: []IRequestInterceptor{interceptor},
	}
}

func init() {
	// get all files in ./plugins/interceptor/response
	interceptorRespFiles, err := filepath.Glob("./plugins/interceptor/response/*")
	if err != nil {
		log.Fatalf("Error opening plugin: %v", err)
	}
	for _, file := range interceptorRespFiles {
		InitRespInterceptor.Add(NewInterceptorResponseChainFromFile(file))
	}
	// get all files in ./plugins/interceptor/request
	interceptorReqFiles, err := filepath.Glob("./plugins/interceptor/request/*")
	if err != nil {
		log.Fatalf("Error opening plugin: %v", err)
	}
	for _, file := range interceptorReqFiles {
		InitReqInterceptor.Add(NewInterceptorRequestChainFromFile(file))
	}
}
