package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/google/uuid"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	routeDto "github.com/starton-io/tyrscale/manager/api/proxy/route/dto"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	"github.com/starton-io/tyrscale/manager/api/proxy/upstream/dto"
	"github.com/starton-io/tyrscale/manager/api/proxy/upstream/repository"
	rpcDto "github.com/starton-io/tyrscale/manager/api/rpc/dto"
	rpcRepo "github.com/starton-io/tyrscale/manager/api/rpc/repository"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

const (
	TOPIC_UPSTREAM_UPSERTED = "upstream_upserted"
	TOPIC_UPSTREAM_DELETED  = "upstream_deleted"
)

func urlParse(uri string, upstream *pb.UpstreamModel) error {
	u, err := url.Parse(uri)
	if err != nil {
		return err
	}
	port := u.Port()
	if port == "" && u.Scheme == "https" {
		port = "443"
	} else if port == "" && u.Scheme == "http" {
		port = "80"
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		return err
	}

	upstream.Host = u.Hostname()
	upstream.Port = int32(portInt)
	upstream.Path = u.Path
	upstream.Scheme = u.Scheme
	return nil
}

type UpstreamService struct {
	repo      repository.IUpstreamRepository
	routeRepo routeRepo.IRouteRepository
	rpcRepo   rpcRepo.IRPCRepository
	pub       pubsub.IPub
}

//go:generate mockery --name=IUpstreamService
type IUpstreamService interface {
	Upsert(ctx context.Context, routeUuid string, req *dto.Upstream) (*dto.UpstreamUpsertRes, error)
	List(ctx context.Context, req *dto.ListUpstreamReq) ([]*pb.UpstreamModel, error)
	Delete(ctx context.Context, req *dto.UpstreamDeleteReq) error
}

func NewUpstreamService(repo repository.IUpstreamRepository, routeRepo routeRepo.IRouteRepository, rpcRepo rpcRepo.IRPCRepository, pub pubsub.IPub) IUpstreamService {
	return &UpstreamService{repo: repo, routeRepo: routeRepo, rpcRepo: rpcRepo, pub: pub}
}

func (s *UpstreamService) Upsert(ctx context.Context, routeUuid string, upstream *dto.Upstream) (*dto.UpstreamUpsertRes, error) {
	ctx, span := tracer.Start(ctx, "UpsertUpstreamService")
	defer tracer.SafeEndSpan(span)
	if upstream.Uuid != "" {
		// check if upstream exists
		upstreams, err := s.repo.List(ctx, &dto.ListUpstreamReq{RouteUuid: routeUuid, Uuid: &upstream.Uuid})
		if err != nil {
			return nil, err
		}
		if len(upstreams) == 0 {
			return nil, errors.New("upstream not found")
		}
		fmt.Println(upstreams)
	} else {
		upstream.Uuid = uuid.New().String()
	}
	// check if route exists
	logger.Infof("Route uuid: %s", routeUuid)
	routes, err := s.routeRepo.List(ctx, &routeDto.ListRouteReq{Uuid: routeUuid})
	if err != nil {
		return nil, err
	}
	if len(routes) == 0 {
		return nil, errors.New("route not found")
	}

	upstreamModel := &pb.UpstreamModel{
		Uuid:   upstream.Uuid,
		Weight: &upstream.Weight,
		Host:   upstream.Host,
		Port:   int32(upstream.Port),
		Path:   upstream.Path,
		Scheme: upstream.Scheme,
	}

	// check if rpc exists
	if upstream.RpcUuid != "" {
		rpc, err := s.rpcRepo.List(ctx, &rpcDto.ListReq{ListFilterReq: rpcDto.ListFilterReq{UUID: upstream.RpcUuid}})
		if err != nil {
			return nil, err
		}
		if len(rpc) == 0 {
			return nil, errors.New("rpc not found")
		}
		rpcURL := rpc[0].Url
		// parse url to scheme, host, port, path
		err = urlParse(rpcURL, upstreamModel)
		if err != nil {
			return nil, err
		}
		upstreamModel.RpcUuid = &upstream.RpcUuid
	}

	// upsert upstream
	err = s.repo.Upsert(ctx, routeUuid, upstreamModel)
	if err != nil {
		return nil, err
	}
	upstreamPublishModel := &pb.UpstreamPublishUpsertModel{
		Uuid:      upstreamModel.Uuid,
		RouteHost: routes[0].Host,
		RoutePath: routes[0].Path,
		Host:      upstreamModel.Host,
		Port:      upstreamModel.Port,
		Path:      upstreamModel.Path,
		Scheme:    upstreamModel.Scheme,
		Weight:    upstream.Weight,
	}

	upstreamBytes, err := proto.Marshal(upstreamPublishModel)
	if err != nil {
		return nil, err
	}

	msg := message.NewMessage(uuid.New().String(), upstreamBytes)
	err = s.pub.Publish(ctx, TOPIC_UPSTREAM_UPSERTED, msg)
	if err != nil {
		return nil, err
	}

	return &dto.UpstreamUpsertRes{Uuid: upstream.Uuid}, nil
}

func (s *UpstreamService) UpsertScores(ctx context.Context, req *dto.Upstream) (*dto.UpstreamUpsertRes, error) {
	return nil, nil
}

func (s *UpstreamService) Delete(ctx context.Context, req *dto.UpstreamDeleteReq) error {
	ctx, span := tracer.Start(ctx, "DeleteUpstreamService", trace.WithAttributes(attribute.String("app.upstream.UUID", req.Uuid)))
	defer tracer.SafeEndSpan(span)

	err := s.repo.Delete(ctx, req)
	if err != nil {
		return err
	}
	routes, err := s.routeRepo.List(ctx, &routeDto.ListRouteReq{Uuid: req.RouteUuid})
	if err != nil {
		return err
	}
	upstreamDeleteModel := &pb.UpstreamPublishDeleteModel{
		Uuid:      req.Uuid,
		RouteHost: routes[0].Host,
		RoutePath: routes[0].Path,
	}

	upstreamBytes, err := proto.Marshal(upstreamDeleteModel)
	if err != nil {
		return err
	}
	msg := message.NewMessage(uuid.New().String(), upstreamBytes)
	return s.pub.Publish(ctx, TOPIC_UPSTREAM_DELETED, msg)
}

func (s *UpstreamService) List(ctx context.Context, req *dto.ListUpstreamReq) ([]*pb.UpstreamModel, error) {
	ctx, span := tracer.Start(ctx, "ListUpstreamService")
	defer tracer.SafeEndSpan(span)
	listUpstream, err := s.repo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	for _, upstream := range listUpstream {
		if upstream.RpcUuid != nil {
			rpcs, err := s.rpcRepo.List(ctx, &rpcDto.ListReq{ListFilterReq: rpcDto.ListFilterReq{UUID: *upstream.RpcUuid}})
			if err != nil {
				return nil, err
			}
			if len(rpcs) == 0 {
				return nil, errors.New("rpc not found")
			}
			err = urlParse(rpcs[0].Url, upstream)
			if err != nil {
				return nil, err
			}
			upstream.RpcUuid = &rpcs[0].Uuid
		}
	}
	return listUpstream, nil
}
