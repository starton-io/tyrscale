package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	pluginRepo "github.com/starton-io/tyrscale/manager/api/proxy/plugin/repository"
	"github.com/starton-io/tyrscale/manager/api/proxy/plugin/service"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
)

func Routes(r fiber.Router, validator validation.Validation, listenerURL string, kvDB kv.IRedisStore, publisher pubsub.IPub) {
	routeRepo := routeRepo.NewRouteRepository(kvDB)
	pluginRepo := pluginRepo.NewPluginRepository(kvDB)
	pluginSvc := service.NewPluginService(listenerURL, routeRepo, pluginRepo, publisher)
	pluginHandler := NewPluginHandler(pluginSvc, validator)

	grp := r.Group("/plugins")
	{
		grp.Get("/", pluginHandler.ListPlugins)
	}

	grpRoute := r.Group("/routes")
	{
		grpRoute.Post("/:uuid/attach-plugin", pluginHandler.AttachPlugin)
		grpRoute.Post("/:uuid/detach-plugin", pluginHandler.DetachPlugin)
		grpRoute.Get("/:uuid/plugins", pluginHandler.ListPluginsFromRoute)
	}
}
