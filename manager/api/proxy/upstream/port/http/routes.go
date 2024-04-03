package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	"github.com/starton-io/tyrscale/manager/api/proxy/upstream/repository"
	"github.com/starton-io/tyrscale/manager/api/proxy/upstream/service"
	rpcRepo "github.com/starton-io/tyrscale/manager/api/rpc/repository"
)

func Routes(r fiber.Router, kvDB kv.IRedisStore, validator validation.Validation, pub pubsub.IPub) {
	upstreamRepo := repository.NewUpstreamRepository(kvDB)
	routeRepo := routeRepo.NewRouteRepository(kvDB)
	rpcRepo := rpcRepo.NewRPCRepository(kvDB)
	upstreamSvc := service.NewUpstreamService(upstreamRepo, routeRepo, rpcRepo, pub)
	upstreamHandler := NewUpstreamHandler(upstreamSvc, validator)
	grp := r.Group("/routes/:route_uuid/upstreams")
	{
		grp.Put("/", upstreamHandler.UpsertUpstream)
		grp.Get("/", upstreamHandler.ListUpstreams)
		grp.Delete("/:uuid", upstreamHandler.DeleteUpstream)
	}
	grp = r.Group("/upstreams")
	{
		grp.Get("/", upstreamHandler.ListUpstreams)
	}
}
