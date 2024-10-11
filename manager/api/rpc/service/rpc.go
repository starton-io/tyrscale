package service

import (
	"context"
	"fmt"
	"net/url"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	networkDto "github.com/starton-io/tyrscale/manager/api/network/dto"
	networkRepo "github.com/starton-io/tyrscale/manager/api/network/repository"
	upstreamDto "github.com/starton-io/tyrscale/manager/api/proxy/upstream/dto"
	upstreamService "github.com/starton-io/tyrscale/manager/api/proxy/upstream/service"
	"github.com/starton-io/tyrscale/manager/api/rpc/dto"
	"github.com/starton-io/tyrscale/manager/api/rpc/repository"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/rpc"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

//go:generate mockery --name=IRPCService --output=./mocks
type IRPCService interface {
	Create(ctx context.Context, rpc *dto.CreateRpcReq) (*dto.CreateRpcRes, *dto.CreateRpcCtx, error)
	Update(ctx context.Context, rpc *dto.UpdateRpcReq) error
	List(ctx context.Context, filterParams *dto.ListReq) ([]*pb.RpcModel, error)
	Delete(ctx context.Context, req *dto.DeleteRpcReq) error
}

type RPCService struct {
	repo            repository.IRPCRepository
	netRepo         networkRepo.INetworkRepository
	upstreamService upstreamService.IUpstreamService
	pub             pubsub.IPub
}

func NewRPCService(repo repository.IRPCRepository, netRepo networkRepo.INetworkRepository, upstreamService upstreamService.IUpstreamService, pub pubsub.IPub) *RPCService {
	return &RPCService{
		repo:            repo,
		netRepo:         netRepo,
		upstreamService: upstreamService,
		pub:             pub,
	}
}

func (s *RPCService) Create(ctx context.Context, req *dto.CreateRpcReq) (*dto.CreateRpcRes, *dto.CreateRpcCtx, error) {
	//validate request
	err := req.Type.Validate()
	if err != nil {
		return nil, nil, err
	}

	// check if uuid is empty and generate new one
	if req.UUID == "" {
		req.UUID = uuid.New().String()
	}

	ctx, span := tracer.Start(ctx, "CreateRPCService", trace.WithAttributes(attribute.String("app.rpc.UUID", req.UUID)))
	defer tracer.SafeEndSpan(span)

	var rpc pb.RpcModel

	//check if network exists by Name
	netRes, err := s.netRepo.List(ctx, &networkDto.ListNetworkReq{Name: req.NetworkName})
	if err != nil {
		return nil, nil, err
	}
	if len(netRes) == 0 {
		return nil, nil, fmt.Errorf("network with name %s not found", req.NetworkName)
	}

	// check if rpc with same uuid already exists
	existingRpc, err := s.repo.List(ctx, &dto.ListReq{ListFilterReq: dto.ListFilterReq{URL: req.URL}})
	if err != nil {
		return nil, nil, err
	}
	if len(existingRpc) > 0 {
		// transform url to dns
		parsedURL, err := url.Parse(req.URL)
		if err != nil {
			logger.Errorf("error parsing url %s", err)
		}
		return nil, &dto.CreateRpcCtx{UUID: existingRpc[0].Uuid}, fmt.Errorf("rpc with url %s already exists", parsedURL.Host)
	}

	utils.Copy(&rpc, req)
	rpc.ChainId = netRes[0].ChainId
	err = s.repo.Create(ctx, &rpc)
	if err != nil {
		return nil, nil, err
	}
	rpcBytes, err := proto.Marshal(&rpc)
	if err != nil {
		return nil, nil, err
	}
	msg := message.NewMessage(uuid.New().String(), rpcBytes)
	err = s.pub.Publish(ctx, "rpc_created", msg)
	if err != nil {
		return nil, nil, err
	}
	return &dto.CreateRpcRes{UUID: req.UUID}, nil, nil
}

func (s *RPCService) List(ctx context.Context, listReq *dto.ListReq) ([]*pb.RpcModel, error) {
	ctx, span := tracer.Start(ctx, "ListRPCService")
	defer tracer.SafeEndSpan(span)

	rpcs, err := s.repo.List(ctx, listReq)
	if err != nil {
		return nil, err
	}
	return rpcs, nil
}

func (s *RPCService) Delete(ctx context.Context, rpc *dto.DeleteRpcReq) error {
	ctx, span := tracer.Start(ctx, "DeleteRPCService", trace.WithAttributes(attribute.String("app.rpc.UUID", rpc.UUID)))
	defer tracer.SafeEndSpan(span)

	res, err := s.repo.List(ctx, &dto.ListReq{ListFilterReq: dto.ListFilterReq{UUID: rpc.UUID}})
	if err != nil {
		return err
	}
	if len(res) == 0 {
		return fmt.Errorf("rpc with uuid %s does not exist", rpc.UUID)
	}
	rpcRes := res[0]
	associatedUpstream, err := s.repo.ListAssociatedUpstreamByRPCUuid(ctx, rpcRes.Uuid)
	if err != nil {
		return err
	}
	logger.Debugf("associatedUpstream: %v", associatedUpstream)
	if len(associatedUpstream) > 0 {
		err = s.repo.DeleteAssociatedUpstream(ctx, rpcRes.Uuid)
		if err != nil {
			return err
		}
		for _, upstream := range associatedUpstream {
			logger.Debugf("upstream: %v", upstream)
			err = s.upstreamService.Delete(ctx, &upstreamDto.UpstreamDeleteReq{Uuid: upstream.Uuid, RouteUuid: upstream.RouteUuid})
			if err != nil {
				return err
			}
		}
	}
	err = s.repo.Delete(ctx, rpcRes)
	if err != nil {
		return err
	}
	rpcBytes, err := proto.Marshal(rpcRes)
	if err != nil {
		return err
	}
	logger.Debug("rpc deleted", zap.String("uuid", rpc.UUID))
	msg := message.NewMessage(uuid.New().String(), rpcBytes)
	return s.pub.Publish(ctx, "rpc_deleted", msg)
}

func (s *RPCService) Update(ctx context.Context, rpc *dto.UpdateRpcReq) error {
	ctx, span := tracer.Start(ctx, "UpdateRPCService", trace.WithAttributes(attribute.String("app.rpc.UUID", rpc.UUID)))
	defer tracer.SafeEndSpan(span)

	if rpc.Type != nil {
		err := rpc.Type.Validate()
		if err != nil {
			return err
		}
	}

	res, err := s.repo.List(ctx, &dto.ListReq{ListFilterReq: dto.ListFilterReq{UUID: rpc.UUID}})
	if err != nil {
		return err
	}
	if len(res) == 0 {
		return fmt.Errorf("rpc with uuid %s does not exist", rpc.UUID)
	}
	updateRpc := res[0]
	err = copier.CopyWithOption(updateRpc, rpc, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	if err != nil {
		return err
	}
	err = s.repo.Update(ctx, updateRpc)
	if err != nil {
		return err
	}
	rpcBytes, err := proto.Marshal(updateRpc)
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.New().String(), rpcBytes)
	return s.pub.Publish(ctx, "rpc_updated", msg)
}
