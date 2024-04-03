package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	networkRepo "github.com/starton-io/tyrscale/manager/api/network/repository"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	recommendationRepo "github.com/starton-io/tyrscale/manager/api/recommendation/repository"
	recommendationService "github.com/starton-io/tyrscale/manager/api/recommendation/service"
)

func Routes(r fiber.Router, kvDB kv.IRedisStore, validator validation.Validation, pubsub pubsub.IPub) {
	netRepo := networkRepo.NewNetworkRepository(kvDB)
	recommendationRepo := recommendationRepo.NewRecommendationRepository(kvDB)
	routeRepo := routeRepo.NewRouteRepository(kvDB)
	recommendationSvc := recommendationService.NewRecommendationService(recommendationRepo, netRepo, routeRepo, pubsub)
	handler := NewHandler(recommendationSvc, validator)
	grp := r.Group("/recommendations")
	{
		grp.Post("/", handler.CreateRecommendation)
		grp.Put("/", handler.UpdateRecommendation)
		grp.Get("/", handler.ListRecommendation)
		grp.Delete("/:route_uuid", handler.DeleteRecommendation)
	}
}
