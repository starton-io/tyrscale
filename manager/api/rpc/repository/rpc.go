package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/starton-io/tyrscale/manager/api/network"
	"github.com/starton-io/tyrscale/manager/api/rpc/dto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"

	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/sorter"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	constrpc "github.com/starton-io/tyrscale/manager/api/rpc"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/rpc"
	upstreamPb "github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
)

const (
	BASE_KEY          = "rpcs"
	BASE_KEY_UPSTREAM = "upstreams"
)

//go:generate mockery --name=IRPCRepository
type IRPCRepository interface {
	Create(ctx context.Context, rpc *pb.RpcModel) error
	List(ctx context.Context, filter *dto.ListReq) ([]*pb.RpcModel, error)
	Update(ctx context.Context, rpc *pb.RpcModel) error
	Delete(ctx context.Context, rpc *pb.RpcModel) error
	DeleteAssociatedUpstream(ctx context.Context, uuid string) error
	ListAssociatedUpstream(ctx context.Context, uuid string) ([]*upstreamPb.UpstreamRPCRouteAssociation, error)
}

type RPCRepository struct {
	baseKey string
	kvDB    kv.IRedisStore
}

func NewRPCRepository(kvDB kv.IRedisStore) IRPCRepository {
	return &RPCRepository{
		baseKey: BASE_KEY,
		kvDB:    kvDB,
	}
}

func (r *RPCRepository) Create(ctx context.Context, rpc *pb.RpcModel) error {
	ctx, span := tracer.Start(ctx, "CreateRPCRepository", trace.WithAttributes(attribute.String("app.rpc.UUID", rpc.Uuid)))
	defer tracer.SafeEndSpan(span)

	globalPrefix, rClient := r.kvDB.GetClient()
	exist, err := rClient.HExists(ctx, globalPrefix+BASE_KEY, rpc.Uuid).Result()
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("rpc with uuid %s already exist", rpc.Uuid)
	}

	rpcBytes, err := proto.Marshal(rpc)
	if err != nil {
		return err
	}
	tx := rClient.TxPipeline()
	err = tx.HSet(ctx, globalPrefix+constrpc.BASE_KEY, rpc.Uuid, rpcBytes).Err()
	if err != nil {
		return err
	}
	err = tx.ZAdd(ctx, fmt.Sprintf("%s:%s:%s", globalPrefix+network.BASE_KEY, rpc.NetworkName, constrpc.BASE_KEY), redis.Z{Score: 0, Member: rpc.Uuid}).Err()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx)
	return err
}

func (r *RPCRepository) List(ctx context.Context, listReq *dto.ListReq) ([]*pb.RpcModel, error) {
	ctx, span := tracer.Start(ctx, "ListRPCRepository")
	defer tracer.SafeEndSpan(span)

	matchCriteria, err := utils.StructToMapStr(listReq.ListFilterReq, "query")
	if err != nil {
		return nil, err
	}
	filter := kv.NewParamsFilterPB[*pb.RpcModel](matchCriteria, false, "", 0)
	if matchCriteria["uuid"] != "" {
		filter.EnablePrefilter = true
		filter.PrefilterPattern = matchCriteria["uuid"]
	}

	listRPC, err := r.kvDB.ScanHash(ctx, r.baseKey, filter)
	if err != nil {
		return nil, err
	}
	protoRPC, err := utils.UnmarshalSliceBytesToProto[*pb.RpcModel](listRPC)
	if err != nil {
		return nil, err
	}

	// sort rpcs
	if listReq.ListSortReq != nil && listReq.ListSortReq.SortBy != "" {
		sortStrategy := sorter.NewSortByFieldPB[*pb.RpcModel](listReq.ListSortReq.SortBy, listReq.ListSortReq.SortDescending)
		protoRPC, err = sortStrategy.Sort(protoRPC)
		if err != nil {
			return nil, err
		}
	}
	return protoRPC, nil
}

func (r *RPCRepository) Delete(ctx context.Context, rpc *pb.RpcModel) error {
	ctx, span := tracer.Start(ctx, "DeleteRPCRepository", trace.WithAttributes(attribute.String("app.rpc.UUID", rpc.Uuid)))
	defer tracer.SafeEndSpan(span)

	globalPrefix, rClient := r.kvDB.GetClient()
	tx := rClient.TxPipeline()
	err := tx.HDel(ctx, globalPrefix+constrpc.BASE_KEY, rpc.Uuid).Err()
	if err != nil {
		return err
	}
	err = tx.ZRem(ctx, fmt.Sprintf("%s:%s:%s", globalPrefix+network.BASE_KEY, rpc.NetworkName, constrpc.BASE_KEY), rpc.Uuid).Err()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx)
	return err
}

func (r *RPCRepository) DeleteAssociatedUpstream(ctx context.Context, uuid string) error {
	ctx, span := tracer.Start(ctx, "DeleteAssociatedUpstreamRPCRepository", trace.WithAttributes(attribute.String("app.rpc.UUID", uuid)))
	defer tracer.SafeEndSpan(span)
	key := fmt.Sprintf("%s:%s", uuid, BASE_KEY_UPSTREAM)
	return r.kvDB.Delete(ctx, key)
}

func (r *RPCRepository) ListAssociatedUpstream(ctx context.Context, uuid string) ([]*upstreamPb.UpstreamRPCRouteAssociation, error) {
	ctx, span := tracer.Start(ctx, "ListAssociatedUpstreamRPCRepository", trace.WithAttributes(attribute.String("app.rpc.UUID", uuid)))
	defer tracer.SafeEndSpan(span)
	key := fmt.Sprintf("%s:%s", uuid, BASE_KEY_UPSTREAM)
	filter := kv.NewParamsFilterPB[*upstreamPb.UpstreamRPCRouteAssociation](make(map[string]string), true, "*", 0)
	list, err := r.kvDB.ScanHash(ctx, key, filter)
	if err != nil {
		return nil, err
	}
	protoRPC, err := utils.UnmarshalSliceBytesToProto[*upstreamPb.UpstreamRPCRouteAssociation](list)
	if err != nil {
		return nil, err
	}
	return protoRPC, nil
}

func (r *RPCRepository) Update(ctx context.Context, rpc *pb.RpcModel) error {
	ctx, span := tracer.Start(ctx, "UpdateRPCRepository", trace.WithAttributes(attribute.String("app.rpc.UUID", rpc.Uuid)))
	defer tracer.SafeEndSpan(span)

	key := r.baseKey
	serialized, err := json.Marshal(rpc)
	if err != nil {
		return err
	}
	return r.kvDB.StoreHash(ctx, key, rpc.Uuid, string(serialized))
}
