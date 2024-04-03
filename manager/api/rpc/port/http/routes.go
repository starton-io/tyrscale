package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	networkRepo "github.com/starton-io/tyrscale/manager/api/network/repository"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	upstreamService "github.com/starton-io/tyrscale/manager/api/proxy/upstream/service"

	upstreamRepo "github.com/starton-io/tyrscale/manager/api/proxy/upstream/repository"
	"github.com/starton-io/tyrscale/manager/api/rpc/repository"
	"github.com/starton-io/tyrscale/manager/api/rpc/service"
)

func Routes(r fiber.Router, kvDB kv.IRedisStore, validator validation.Validation, pub pubsub.IPub) {
	netRepo := networkRepo.NewNetworkRepository(kvDB)
	rpcRepo := repository.NewRPCRepository(kvDB)
	routeRepo := routeRepo.NewRouteRepository(kvDB)
	upstreamRepo := upstreamRepo.NewUpstreamRepository(kvDB)
	upstreamSvc := upstreamService.NewUpstreamService(upstreamRepo, routeRepo, rpcRepo, pub)
	rpcSvc := service.NewRPCService(rpcRepo, netRepo, upstreamSvc, pub)
	handler := NewRPCHandler(rpcSvc, validator)
	grp := r.Group("/rpcs")
	{
		grp.Post("/", handler.CreateRPC)
		grp.Get("/", handler.ListRPCs)
		grp.Delete("/:uuid", handler.DeleteRPC)
		grp.Put("/", handler.UpdateRPC)
	}

}
