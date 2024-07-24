package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/manager/api/proxy/plugin/service"
)

func Routes(r fiber.Router, listenerURL string) {
	pluginSvc := service.NewPluginService(listenerURL)
	pluginHandler := NewPluginHandler(pluginSvc)
	grp := r.Group("/plugins")
	{
		grp.Get("/", pluginHandler.ListPlugins)
	}
}
