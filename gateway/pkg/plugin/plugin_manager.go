package plugin

import (
	"context"
	"errors"
	"fmt"
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
	responsePlugins        map[string]RegistererResponse
	requestPlugins         map[string]RegistererRequest
	middlewarePlugins      map[string]RegistererMiddleware
	storage                IPluginStorage
}

type IPluginManager interface {
	RegisterResponseInterceptor(pluginName string, configPayload []byte) (types.IResponseInterceptor, error)
	RegisterRequestInterceptor(pluginName string, configPayload []byte) (types.IRequestInterceptor, error)
	RegisterMiddleware(pluginName string, configPayload []byte) (typesMiddleware.MiddlewareFunc, error)
	ValidatePlugin(pluginName string, pluginType string, configPayload []byte) error
	LoadPlugins(config *config.PluginConfig) (int, error)
	Store() IPluginStorage
}

func NewPluginManager(pluginStorage IPluginStorage) IPluginManager {
	return &PluginManager{
		responsePlugins:   make(map[string]RegistererResponse),
		requestPlugins:    make(map[string]RegistererRequest),
		middlewarePlugins: make(map[string]RegistererMiddleware),
		storage:           pluginStorage,
	}
}

func (pm *PluginManager) ValidatePlugin(pluginName string, pluginType string, configPayload []byte) error {

	switch pluginType {
	case string(PluginTypeResponseInterceptor):
		plugin, ok := pm.responsePlugins[pluginName]
		if !ok {
			return errors.New("plugin not found")
		}
		return plugin.Validate(configPayload)
	case string(PluginTypeRequestInterceptor):
		plugin, ok := pm.requestPlugins[pluginName]
		if !ok {
			return errors.New("plugin not found")
		}
		return plugin.Validate(configPayload)
	case string(PluginTypeMiddleware):
		plugin, ok := pm.middlewarePlugins[pluginName]
		if !ok {
			return errors.New("plugin not found")
		}
		return plugin.Validate(configPayload)
	}
	return nil
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

func (pm *PluginManager) Store() IPluginStorage {
	return pm.storage
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
			return pm.registerPlugin(p.Name, pluginInstance)
		}
	}

	return fmt.Errorf("plugin does not implement any of the required interfaces")
}

func (pm *PluginManager) registerPlugin(pluginName string, p interface{}) error {
	switch plugin := p.(type) {
	case RegistererResponse:
		pm.responsePlugins[pluginName] = plugin
		_ = pm.storage.AddPlugin(context.Background(), pluginName, string(PluginTypeResponseInterceptor))
	case RegistererRequest:
		pm.requestPlugins[pluginName] = plugin
		_ = pm.storage.AddPlugin(context.Background(), pluginName, string(PluginTypeRequestInterceptor))
	case RegistererMiddleware:
		pm.middlewarePlugins[pluginName] = plugin
		_ = pm.storage.AddPlugin(context.Background(), pluginName, string(PluginTypeMiddleware))
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

func (pm *PluginManager) RegisterResponseInterceptor(pluginName string, configPayload []byte) (types.IResponseInterceptor, error) {
	responsePlugin, exists := pm.responsePlugins[pluginName]
	if !exists {
		return nil, errors.New("response interceptor plugin not found")
	}
	var interceptorResponse types.IResponseInterceptor
	err := responsePlugin.RegisterResponseInterceptor(func(interceptorName string, interceptor types.IResponseInterceptor) {
		interceptorResponse = interceptor
	}, configPayload)
	if err != nil {
		return nil, err
	}
	return interceptorResponse, nil
}

func (pm *PluginManager) RegisterRequestInterceptor(pluginName string, configPayload []byte) (types.IRequestInterceptor, error) {
	requestPlugin, exists := pm.requestPlugins[pluginName]
	if !exists {
		return nil, errors.New("request interceptor plugin not found")
	}
	var interceptorRequest types.IRequestInterceptor
	err := requestPlugin.RegisterRequestInterceptor(func(interceptorName string, interceptor types.IRequestInterceptor) {
		interceptorRequest = interceptor
	}, configPayload)
	if err != nil {
		return nil, err
	}
	return interceptorRequest, nil
}

func (pm *PluginManager) RegisterMiddleware(pluginName string, configPayload []byte) (typesMiddleware.MiddlewareFunc, error) {
	middlewarePlugin, exists := pm.middlewarePlugins[pluginName]
	if !exists {
		return nil, errors.New("middleware plugin not found")
	}
	var interceptorMiddleware typesMiddleware.MiddlewareFunc
	err := middlewarePlugin.RegisterMiddleware(func(name string, middleware typesMiddleware.MiddlewareFunc) {
		interceptorMiddleware = middleware
	}, configPayload)
	if err != nil {
		return nil, err
	}
	return interceptorMiddleware, nil
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
	RegisterResponseInterceptor(func(
		name string,
		interceptor types.IResponseInterceptor,
	), []byte) error
	Validate([]byte) error
}

type RegistererMiddleware interface {
	RegisterMiddleware(func(
		name string,
		middleware typesMiddleware.MiddlewareFunc,
	), []byte) error
	Validate([]byte) error
}

type RegistererRequest interface {
	RegisterRequestInterceptor(func(
		name string,
		interceptor types.IRequestInterceptor,
	), []byte) error
	Validate([]byte) error
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
