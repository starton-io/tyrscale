package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	pluginService "github.com/starton-io/tyrscale/manager/api/proxy/plugin/service"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/service"
	upstreamRepository "github.com/starton-io/tyrscale/manager/api/proxy/upstream/repository"
	upstreamService "github.com/starton-io/tyrscale/manager/api/proxy/upstream/service"
	recommendationRepository "github.com/starton-io/tyrscale/manager/api/recommendation/repository"
	rpcRepo "github.com/starton-io/tyrscale/manager/api/rpc/repository"
)

func Routes(r fiber.Router, kvDB kv.IRedisStore, validator validation.Validation, pub pubsub.IPub, gatewayUrl string) {
	routeRepo := repository.NewRouteRepository(kvDB)
	upstreamRepo := upstreamRepository.NewUpstreamRepository(kvDB)
	recommendationRepo := recommendationRepository.NewRecommendationRepository(kvDB)
	rpcRepo := rpcRepo.NewRPCRepository(kvDB)
	pluginSvc := pluginService.NewPluginService(gatewayUrl)
	routeSvc := service.NewRouteService(routeRepo, recommendationRepo, pub, pluginSvc)

	upstreamSvc := upstreamService.NewUpstreamService(upstreamRepo, routeRepo, rpcRepo, pub)
	routeHandler := NewRouteHandler(routeSvc, upstreamSvc, validator)
	grp := r.Group("/routes")
	{
		grp.Post("/", routeHandler.CreateRoute)
		grp.Put("/:uuid", routeHandler.UpdateRoute)
		grp.Get("/", routeHandler.ListRoutes)
		grp.Delete("/:uuid", routeHandler.DeleteRoute)
		grp.Post("/:uuid/attach-plugin", routeHandler.AttachPlugin)
		grp.Post("/:uuid/detach-plugin", routeHandler.DetachPlugin)
	}
}
