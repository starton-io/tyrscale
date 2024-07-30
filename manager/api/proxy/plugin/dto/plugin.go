package dto

import "github.com/starton-io/tyrscale/gateway/pkg/plugin"

type Plugins struct {
	Middleware          []*Plugin `json:"Middleware,omitempty" validate:"omitempty"`
	InterceptorRequest  []*Plugin `json:"RequestInterceptor,omitempty" validate:"omitempty"`
	InterceptorResponse []*Plugin `json:"ResponseInterceptor,omitempty" validate:"omitempty"`
}

type Plugin struct {
	Name        string      `json:"name" validate:"required"`
	Description string      `json:"description" validate:"omitempty"`
	Priority    int         `json:"priority" validate:"required,gte=1,lte=1000"`
	Config      interface{} `json:"config,omitempty" validate:"omitempty"`
}

type AttachPluginReq struct {
	Name        string            `json:"name" validate:"required"`
	Type        plugin.PluginType `json:"type" validate:"required"`
	Config      interface{}       `json:"config" validate:"required"`
	Description string            `json:"description" validate:"omitempty"`
	Priority    int               `json:"priority" validate:"required,gte=1,lte=1000"`
}

type DetachPluginReq struct {
	Name string            `json:"name" validate:"required"`
	Type plugin.PluginType `json:"type" validate:"required"`
}

type ListPluginReq struct {
	RouteUuid string            `query:"route_uuid" validate:"required"`
	Type      plugin.PluginType `query:"type" validate:"required"`
	Name      string            `query:"name" validate:"omitempty"`
}
