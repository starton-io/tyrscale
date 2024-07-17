package service

import (
	"context"
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/dto"
	recommendationDto "github.com/starton-io/tyrscale/manager/api/recommendation/dto"

	"github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	recommendationRepository "github.com/starton-io/tyrscale/manager/api/recommendation/repository"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/route"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

var (
	routeCreatedTopic = "route_created"
	routeDeletedTopic = "route_deleted"
	routeUpdatedTopic = "route_updated"
)

type RouteService struct {
	repo               repository.IRouteRepository
	recommendationRepo recommendationRepository.IRecommendationRepository
	pub                pubsub.IPub
}

type IRouteService interface {
	Create(ctx context.Context, req *dto.Route) (*dto.CreateRouteRes, error)
	List(ctx context.Context, req *dto.ListRouteReq) ([]*pb.RouteModel, error)
	Update(ctx context.Context, req *dto.UpdateRouteReq) error
	Delete(ctx context.Context, req *dto.DeleteRouteReq) error
}

func NewRouteService(repo repository.IRouteRepository, recommendationRepo recommendationRepository.IRecommendationRepository, pub pubsub.IPub) *RouteService {
	return &RouteService{repo: repo, recommendationRepo: recommendationRepo, pub: pub}
}

func (s *RouteService) Create(ctx context.Context, req *dto.Route) (*dto.CreateRouteRes, error) {
	if err := req.LoadBalancerStrategy.Validate(); err != nil {
		return nil, err
	}
	routeUuid := uuid.New().String()

	ctx, span := tracer.Start(ctx, "UpsertRouteService", trace.WithAttributes(attribute.String("app.route.uuid", routeUuid)))
	defer tracer.SafeEndSpan(span)

	// check if there is a route with the same host and path
	routes, err := s.repo.List(ctx, &dto.ListRouteReq{Host: req.Host, Path: req.Path})
	if err != nil {
		return nil, err
	}
	if len(routes) > 0 {
		return nil, errors.New("route already exists")
	}

	route := &pb.RouteModel{
		Uuid:                 routeUuid,
		Host:                 req.Host,
		Path:                 req.Path,
		LoadBalancerStrategy: req.LoadBalancerStrategy.String(),
	}

	if req.CircuitBreaker != nil {
		route.CircuitBreaker = &pb.CircuitBreaker{
			Enabled:     req.CircuitBreaker.Enabled,
			MaxRequests: req.CircuitBreaker.MaxRequests,
			Interval:    req.CircuitBreaker.Interval,
			Timeout:     req.CircuitBreaker.Timeout,
		}
	}

	if req.HealthCheck != nil {
		route.HealthCheck = &pb.HealthCheck{
			Type:                       string(req.HealthCheck.Type),
			Interval:                   req.HealthCheck.Interval,
			CombinedWithCircuitBreaker: req.HealthCheck.CombinedWithCircuitBreaker,
			Enabled:                    req.HealthCheck.Enabled,
			Timeout:                    req.HealthCheck.Timeout,
		}
	}

	//if req.Plugins != nil {
	//	for _, plugin := range req.Plugins {
	//		route.Plugins = append(route.Plugins, &pb.Plugin{Name: plugin.Name, Priority: uint32(plugin.Priority)})
	//	}
	//}

	if err := s.repo.Upsert(ctx, route); err != nil {
		return nil, err
	}
	routeBytes, err := proto.Marshal(route)
	if err != nil {
		return nil, err
	}
	msg := message.NewMessage(uuid.New().String(), routeBytes)
	err = s.pub.Publish(ctx, routeCreatedTopic, msg)
	if err != nil {
		return nil, err
	}

	return &dto.CreateRouteRes{Uuid: route.Uuid}, nil
}

func (s *RouteService) Delete(ctx context.Context, req *dto.DeleteRouteReq) error {
	ctx, span := tracer.Start(ctx, "DeleteRouteService", trace.WithAttributes(attribute.String("app.route.uuid", req.Uuid)))
	defer tracer.SafeEndSpan(span)
	routes, err := s.repo.List(ctx, &dto.ListRouteReq{Uuid: req.Uuid})
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		return errors.New("route not found")
	}
	recommendations, err := s.recommendationRepo.List(ctx, &recommendationDto.ListRecommendationReq{RouteUuid: req.Uuid})
	if err != nil {
		return err
	}
	if len(recommendations) > 0 {
		return errors.New("cannot delete route, it is associated to a recommendation")
	}
	err = s.repo.Delete(ctx, routes[0])
	if err != nil {
		return err
	}
	routeBytes, err := proto.Marshal(routes[0])
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.New().String(), routeBytes)
	return s.pub.Publish(ctx, routeDeletedTopic, msg)
}

func (s *RouteService) List(ctx context.Context, req *dto.ListRouteReq) ([]*pb.RouteModel, error) {
	ctx, span := tracer.Start(ctx, "ListRouteService", trace.WithAttributes(attribute.String("app.route.uuid", req.Uuid)))
	defer tracer.SafeEndSpan(span)
	routes, err := s.repo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func (s *RouteService) Update(ctx context.Context, req *dto.UpdateRouteReq) error {
	ctx, span := tracer.Start(ctx, "UpdateRouteService", trace.WithAttributes(attribute.String("app.route.uuid", req.Uuid)))
	defer tracer.SafeEndSpan(span)

	if err := req.LoadBalancerStrategy.Validate(); err != nil {
		return err
	}
	routes, err := s.repo.List(ctx, &dto.ListRouteReq{Uuid: req.Uuid})
	if err != nil {
		return err
	}
	if len(routes) == 0 {
		return errors.New("route not found")
	}
	route := routes[0]
	err = copier.CopyWithOption(route, req, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return err
	}
	err = s.repo.Upsert(ctx, route)
	if err != nil {
		return err
	}
	routeBytes, err := proto.Marshal(route)
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.New().String(), routeBytes)
	return s.pub.Publish(ctx, routeUpdatedTopic, msg)
}
