package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	networkDto "github.com/starton-io/tyrscale/manager/api/network/dto"
	networkRepo "github.com/starton-io/tyrscale/manager/api/network/repository"
	routeDto "github.com/starton-io/tyrscale/manager/api/proxy/route/dto"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	"github.com/starton-io/tyrscale/manager/api/recommendation/dto"
	"github.com/starton-io/tyrscale/manager/api/recommendation/repository"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/recommendation"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

//go:generate mockery --name IRecommendationService --output ./mocks
type IRecommendationService interface {
	Create(ctx context.Context, recommendation *dto.CreateRecommendationReq) (*dto.CreateRecommendationRes, error)
	Update(ctx context.Context, recommendation *dto.UpdateRecommendationReq) error
	List(ctx context.Context, listRecommendationsParams *dto.ListRecommendationReq) ([]*pb.RecommendationModel, error)
	Delete(ctx context.Context, deleteRecommendationParams *dto.DeleteRecommendationReq) error
}

type RecommendationService struct {
	repo      repository.IRecommendationRepository
	netRepo   networkRepo.INetworkRepository
	routeRepo routeRepo.IRouteRepository
	pub       pubsub.IPub
}

func NewRecommendationService(repo repository.IRecommendationRepository, netRepo networkRepo.INetworkRepository, routeRepo routeRepo.IRouteRepository, pub pubsub.IPub) *RecommendationService {
	return &RecommendationService{repo: repo, netRepo: netRepo, routeRepo: routeRepo, pub: pub}
}

func (s *RecommendationService) Create(ctx context.Context, req *dto.CreateRecommendationReq) (*dto.CreateRecommendationRes, error) {
	recommendationUuid := uuid.New().String()
	ctx, span := tracer.Start(ctx, "CreateRecommendationService", trace.WithAttributes(attribute.String("app.recommendation.uuid", recommendationUuid)))
	defer tracer.SafeEndSpan(span)

	err := req.Strategy.Validate()
	if err != nil {
		return nil, err
	}
	recommendation := &pb.RecommendationModel{
		RouteUuid:   recommendationUuid,
		Schedule:    req.Schedule,
		NetworkName: req.NetworkName,
		Strategy:    req.Strategy.String(),
	}

	// check if the route and network exist
	err = s.check(ctx, req.RouteUuid, req.NetworkName)
	if err != nil {
		return nil, err
	}

	err = s.repo.Upsert(ctx, recommendation)
	if err != nil {
		return nil, err
	}

	payload, err := proto.Marshal(recommendation)
	if err != nil {
		return nil, err
	}

	msg := message.NewMessage(uuid.NewString(), payload)
	err = s.pub.Publish(ctx, "recommendation_created", msg)
	if err != nil {
		return nil, err
	}

	return &dto.CreateRecommendationRes{RouteUuid: recommendation.RouteUuid}, nil
}

func (s *RecommendationService) check(ctx context.Context, routeUuid string, networkName string) error {
	// Check if the network exist
	netRes, err := s.netRepo.List(ctx, &networkDto.ListNetworkReq{Name: networkName})
	if err != nil {
		return err
	}
	if len(netRes) == 0 {
		return errors.New("network not found")
	}

	// Check if routes exist
	routeRes, err := s.routeRepo.List(ctx, &routeDto.ListRouteReq{Uuid: routeUuid})
	if err != nil {
		return err
	}
	if len(routeRes) == 0 {
		return errors.New("route not found")
	}

	// check if the network is associate to a existing recommendation
	recommendationRes, err := s.repo.List(ctx, &dto.ListRecommendationReq{
		NetworkName: networkName,
	})
	if err != nil {
		return err
	}
	if len(recommendationRes) > 0 {
		return errors.New("network already associated to a recommendation")
	}

	return nil
}

func (s *RecommendationService) Update(ctx context.Context, req *dto.UpdateRecommendationReq) error {
	ctx, span := tracer.Start(ctx, "UpdateRecommendationService", trace.WithAttributes(attribute.String("app.recommendation.route_uuid", req.RouteUuid)))
	defer tracer.SafeEndSpan(span)

	err := req.Strategy.Validate()
	if err != nil {
		return err
	}

	res, err := s.repo.List(ctx, &dto.ListRecommendationReq{
		RouteUuid: req.RouteUuid,
	})
	if err != nil {
		return err
	}
	if len(res) == 0 {
		return errors.New("recommendation not found")
	}
	recommendationUpdate := res[0]
	err = copier.CopyWithOption(recommendationUpdate, req, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		return err
	}

	err = s.repo.Upsert(ctx, recommendationUpdate)
	if err != nil {
		return err
	}
	payload, err := json.Marshal(recommendationUpdate)
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.NewString(), payload)
	return s.pub.Publish(ctx, "recommendation_updated", msg)
}

func (s *RecommendationService) List(ctx context.Context, req *dto.ListRecommendationReq) ([]*pb.RecommendationModel, error) {
	ctx, span := tracer.Start(ctx, "ListRecommendationService", trace.WithAttributes(attribute.String("app.recommendation.route_uuid", req.RouteUuid)))
	defer tracer.SafeEndSpan(span)

	return s.repo.List(ctx, req)
}

func (s *RecommendationService) Delete(ctx context.Context, deleteRecommendationParams *dto.DeleteRecommendationReq) error {
	ctx, span := tracer.Start(ctx, "DeleteRecommendationService", trace.WithAttributes(attribute.String("app.recommendation.route_uuid", deleteRecommendationParams.RouteUuid)))
	defer tracer.SafeEndSpan(span)

	res, err := s.repo.List(ctx, &dto.ListRecommendationReq{
		RouteUuid: deleteRecommendationParams.RouteUuid,
	})
	if err != nil {
		return err
	}
	if len(res) == 0 {
		return errors.New("recommendation not found")
	}
	recommendation := res[0]
	err = s.repo.Delete(ctx, recommendation.RouteUuid)
	if err != nil {
		return err
	}
	payload, err := json.Marshal(recommendation)
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.NewString(), payload)
	return s.pub.Publish(ctx, "recommendation_deleted", msg)
}
