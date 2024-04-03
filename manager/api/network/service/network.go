package service

import (
	"context"
	"fmt"

	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/manager/api/network/dto"
	"github.com/starton-io/tyrscale/manager/api/network/repository"
	recommendationDto "github.com/starton-io/tyrscale/manager/api/recommendation/dto"
	recommendationSvc "github.com/starton-io/tyrscale/manager/api/recommendation/service"
	dtoRpc "github.com/starton-io/tyrscale/manager/api/rpc/dto"
	rpcSvc "github.com/starton-io/tyrscale/manager/api/rpc/service"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/network"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

//go:generate mockery --name=INetworkService
type INetworkService interface {
	Create(ctx context.Context, network *dto.Network) error
	List(ctx context.Context, filterParams *dto.ListNetworkReq) ([]*pb.NetworkModel, error)
	Delete(ctx context.Context, name string) error
}

type NetworkService struct {
	repo              repository.INetworkRepository
	rpcSvc            rpcSvc.IRPCService
	recommendationSvc recommendationSvc.IRecommendationService
}

func NewNetworkService(repo repository.INetworkRepository, rpcSvc rpcSvc.IRPCService, recommendationSvc recommendationSvc.IRecommendationService) *NetworkService {
	return &NetworkService{
		repo:              repo,
		rpcSvc:            rpcSvc,
		recommendationSvc: recommendationSvc,
	}
}

func (s *NetworkService) Create(ctx context.Context, req *dto.Network) error {
	// list network if the network exist with the same name or chain id
	ctx, span := tracer.Start(ctx, "CreateNetworkService", trace.WithAttributes(attribute.String("app.network.name", req.Name)))
	defer tracer.SafeEndSpan(span)

	var network pb.NetworkModel
	networks, err := s.repo.List(ctx, &dto.ListNetworkReq{
		Name: req.Name,
	})
	if err != nil {
		return err
	}
	if len(networks) > 0 {
		return fmt.Errorf("network with the same name already exist")
	}

	utils.Copy(&network, req)
	return s.repo.Create(ctx, &network)
}

func (s *NetworkService) List(ctx context.Context, req *dto.ListNetworkReq) ([]*pb.NetworkModel, error) {
	ctx, span := tracer.Start(ctx, "ListNetworkService")
	defer tracer.SafeEndSpan(span)

	return s.repo.List(ctx, req)
}

func (s *NetworkService) Delete(ctx context.Context, name string) error {
	ctx, span := tracer.Start(ctx, "DeleteNetworkService", trace.WithAttributes(attribute.String("app.network.name", name)))
	defer tracer.SafeEndSpan(span)

	//get all the rpc associate to the network
	rpcs, err := s.rpcSvc.List(ctx, &dtoRpc.ListReq{
		ListFilterReq: dtoRpc.ListFilterReq{
			NetworkName: name,
		},
	})
	if err != nil {
		return err
	}
	for _, rpc := range rpcs {
		if err := s.rpcSvc.Delete(ctx, &dtoRpc.DeleteRpcReq{
			UUID: rpc.Uuid,
		}); err != nil {
			return err
		}
	}

	// check if the network is associate to a recommendation
	recommendationRes, err := s.recommendationSvc.List(ctx, &recommendationDto.ListRecommendationReq{
		NetworkName: name,
	})
	if err != nil {
		return err
	}
	if len(recommendationRes) > 0 {
		return fmt.Errorf("network is associated to a recommendation")
	}
	return s.repo.Delete(ctx, name)
}
