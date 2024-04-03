package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/starton-io/tyrscale/manager/api/network/repository"
	"github.com/starton-io/tyrscale/manager/api/network/service"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	upstreamRepo "github.com/starton-io/tyrscale/manager/api/proxy/upstream/repository"
	upstreamSvc "github.com/starton-io/tyrscale/manager/api/proxy/upstream/service"
	recommendationRepo "github.com/starton-io/tyrscale/manager/api/recommendation/repository"
	recommendationSvc "github.com/starton-io/tyrscale/manager/api/recommendation/service"
	rpcRepo "github.com/starton-io/tyrscale/manager/api/rpc/repository"
	rpcSvc "github.com/starton-io/tyrscale/manager/api/rpc/service"
)

func Routes(r fiber.Router, kvDB kv.IRedisStore, validator validation.Validation, pub pubsub.IPub) {
	networkRepo := repository.NewNetworkRepository(kvDB)
	recommendationRepo := recommendationRepo.NewRecommendationRepository(kvDB)
	routeRepo := routeRepo.NewRouteRepository(kvDB)
	recommendationSvc := recommendationSvc.NewRecommendationService(recommendationRepo, networkRepo, routeRepo, pub)
	rpcRepo := rpcRepo.NewRPCRepository(kvDB)
	upstreamRepo := upstreamRepo.NewUpstreamRepository(kvDB)
	upstreamSvc := upstreamSvc.NewUpstreamService(upstreamRepo, routeRepo, rpcRepo, pub)
	rpcSvc := rpcSvc.NewRPCService(rpcRepo, networkRepo, upstreamSvc, pub)
	networkSvc := service.NewNetworkService(networkRepo, rpcSvc, recommendationSvc)
	networkHandler := NewNetworkHandler(networkSvc, validator)
	grp := r.Group("/networks")
	{
		grp.Post("/", networkHandler.CreateNetwork)
		grp.Get("/", networkHandler.ListNetworks)
		grp.Delete("/:name", networkHandler.DeleteNetwork)
	}
}
