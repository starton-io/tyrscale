package plugin

import (
	"context"
	"errors"
	"fmt"
	"log"
	"plugin"
	"strings"

	"github.com/starton-io/tyrscale/gateway/pkg/config"
	"github.com/starton-io/tyrscale/gateway/pkg/interceptor/types"
	typesMiddleware "github.com/starton-io/tyrscale/gateway/pkg/middleware/types"
	"github.com/starton-io/tyrscale/gateway/pkg/utils"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
)

type PluginType string

const (
	PluginTypeResponseInterceptor PluginType = "ResponseInterceptor"
	PluginTypeRequestInterceptor  PluginType = "RequestInterceptor"
	PluginTypeMiddleware          PluginType = "Middleware"
)

func (p PluginType) String() string {
	return string(p)
}

type PluginResponseInterceptor struct {
	Name        string
	Interceptor types.IResponseInterceptor
}

type PluginManager struct {
	responseInterceptors   map[string]types.IResponseInterceptor
	requestInterceptors    map[string]types.IRequestInterceptor
	middlewareInterceptors map[string]typesMiddleware.MiddlewareFunc
	storage                IPluginStorage
}

type IPluginManager interface {
	GetPluginRespInterceptor(name string) (types.IResponseInterceptor, error)
	GetPluginReqInterceptor(name string) (types.IRequestInterceptor, error)
	GetPluginMiddleware(name string) (typesMiddleware.MiddlewareFunc, error)
	LoadPlugins(config *config.PluginConfig) (int, error)
}

func NewPluginManager(pluginStorage IPluginStorage) IPluginManager {
	return &PluginManager{
		responseInterceptors:   make(map[string]types.IResponseInterceptor),
		requestInterceptors:    make(map[string]types.IRequestInterceptor),
		middlewareInterceptors: make(map[string]typesMiddleware.MiddlewareFunc),
		storage:                pluginStorage,
	}
}

func (pm *PluginManager) LoadPlugins(config *config.PluginConfig) (int, error) {
	errors := []error{}
	loadedPlugins := 0
	for k, plugin := range config.Plugins {
		if err := pm.loadPlugin(plugin); err != nil {
			errors = append(errors, fmt.Errorf("plugin #%d (%s): %s", k, plugin.Name, err.Error()))
			continue
		}
		loadedPlugins++
	}
	if len(errors) > 0 {
		return loadedPlugins, loaderError{errors: errors}
	}
	return loadedPlugins, nil
}

func (pm *PluginManager) loadPlugin(p *config.Plugin) error {
	logger.Info("Loading plugin", "name", p.Name, "path", p.Path, "sha256sum", p.Sha256sum)
	ok, err := utils.GetSHA256Checksum(p.Path, p.Sha256sum)
	if err != nil {
		return fmt.Errorf("error getting checksum: %v", err)
	}
	if !ok {
		return fmt.Errorf("checksum %s does not match to %s", p.Sha256sum, p.Path)
	}

	plugin, err := plugin.Open(p.Path)
	if err != nil {
		return fmt.Errorf("error opening plugin: %v", err)
	}

	// Factorized plugin lookup
	pluginTypes := []PluginType{
		PluginTypeResponseInterceptor,
		PluginTypeRequestInterceptor,
		PluginTypeMiddleware,
	}
	for _, name := range pluginTypes {
		if pluginInstance, err := plugin.Lookup(string(name)); err == nil {
			return pm.registerPlugin(pluginInstance)
		}
	}

	return fmt.Errorf("plugin does not implement any of the required interfaces")
}

func (pm *PluginManager) registerPlugin(p interface{}) error {
	switch plugin := p.(type) {
	case RegistererResponse:
		log.Println("registering response interceptor")
		plugin.RegisterResponseInterceptors(func(name string, interceptor types.IResponseInterceptor) {
			pm.responseInterceptors[name] = interceptor
			_ = pm.storage.AddPlugin(context.Background(), name, string(PluginTypeResponseInterceptor))
		})
	case RegistererRequest:
		plugin.RegisterRequestInterceptors(func(name string, interceptor types.IRequestInterceptor) {
			pm.requestInterceptors[name] = interceptor
			_ = pm.storage.AddPlugin(context.Background(), name, string(PluginTypeRequestInterceptor))
		})
	case RegistererMiddleware:
		plugin.RegisterMiddleware(func(name string, middleware typesMiddleware.MiddlewareFunc) {
			logger.Info("registering middleware", "name", name)
			pm.middlewareInterceptors[name] = middleware
			_ = pm.storage.AddPlugin(context.Background(), name, string(PluginTypeMiddleware))
		})
		log.Println("registering middleware")
	default:
		return fmt.Errorf("unknown plugin type: %T", plugin)
	}
	return nil
}

func (pm *PluginManager) GetPluginRespInterceptor(name string) (types.IResponseInterceptor, error) {
	interceptor, ok := pm.responseInterceptors[name]
	if !ok {
		return nil, errors.New("plugin interceptor response not found")
	}
	return interceptor, nil
}

func (pm *PluginManager) GetPluginMiddleware(name string) (typesMiddleware.MiddlewareFunc, error) {
	interceptor, ok := pm.middlewareInterceptors[name]
	if !ok {
		return nil, errors.New("plugin interceptor middleware not found")
	}
	return interceptor, nil
}

func (pm *PluginManager) GetPluginReqInterceptor(name string) (types.IRequestInterceptor, error) {
	interceptor, ok := pm.requestInterceptors[name]
	if !ok {
		return nil, errors.New("plugin interceptor request not found")
	}
	return interceptor, nil
}

type RegistererResponse interface {
	RegisterResponseInterceptors(func(
		name string,
		interceptor types.IResponseInterceptor,
	))
}

type RegistererMiddleware interface {
	RegisterMiddleware(func(
		name string,
		middleware typesMiddleware.MiddlewareFunc,
	))
}

type RegistererRequest interface {
	RegisterRequestInterceptors(func(
		name string,
		interceptor types.IRequestInterceptor,
	))
}

type loaderError struct {
	errors []error
}

func (l loaderError) Error() string {
	msgs := make([]string, len(l.errors))
	for i, err := range l.errors {
		msgs[i] = err.Error()
	}
	return fmt.Sprintf("plugin loader found %d error(s): \n%s", len(msgs), strings.Join(msgs, "\n"))
}

func (l loaderError) Len() int {
	return len(l.errors)
}

func (l loaderError) Errs() []error {
	return l.errors
}
