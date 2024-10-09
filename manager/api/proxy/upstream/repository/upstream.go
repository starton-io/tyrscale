package repository

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"

	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	"github.com/starton-io/tyrscale/manager/api/proxy/upstream/dto"
	rpcRepo "github.com/starton-io/tyrscale/manager/api/rpc/repository"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
)

const (
	BASE_KEY      = "upstreams"
	HASH_BASE_KEY = "hash_upstreams"
	LuaScriptList = `
    local cursor = "0"
    local result = {}
    repeat
        local scan_result = redis.call("SCAN", cursor, "MATCH", "ARGV[1]")
        cursor = scan_result[1]
        local keys = scan_result[2]
        for i, key in ipairs(keys) do
            local upstreams = redis.call("ZRANGE", key, 0, -1, "WITHSCORES")
            local route_upstreams = {key, upstreams}
            table.insert(result, route_upstreams)
        end
    until cursor == "0"
    return result
    `
)

type UpstreamRepository struct {
	baseKey string
	kvDB    kv.IRedisStore
}

//go:generate mockery --name=IUpstreamRepository
type IUpstreamRepository interface {
	List(ctx context.Context, req *dto.ListUpstreamReq) ([]*pb.UpstreamModel, error)
	Upsert(ctx context.Context, routeUuid string, upstream *pb.UpstreamModel) error
	Delete(ctx context.Context, req *dto.UpstreamDeleteReq) error
}

func NewUpstreamRepository(kvDB kv.IRedisStore) IUpstreamRepository {
	return &UpstreamRepository{
		baseKey: BASE_KEY,
		kvDB:    kvDB,
	}
}

func (r *UpstreamRepository) List(ctx context.Context, req *dto.ListUpstreamReq) ([]*pb.UpstreamModel, error) {
	ctx, span := tracer.Start(ctx, "ListUpstreamRepository")
	defer tracer.SafeEndSpan(span)
	filter := kv.NewParamsFilterPB[*pb.UpstreamModel](make(map[string]string), true, "", 0)

	keyUpstreamHash := fmt.Sprintf("%s:%s:%s", routeRepo.BASE_KEY, req.RouteUuid, HASH_BASE_KEY)
	keyUpstreamScore := fmt.Sprintf("%s:%s:%s", routeRepo.BASE_KEY, req.RouteUuid, BASE_KEY)
	listUpstream, err := r.kvDB.ScanHash(ctx, keyUpstreamHash, filter)
	if err != nil {
		return nil, err
	}
	listPbUpstream, err := utils.UnmarshalSliceBytesToProto[*pb.UpstreamModel](listUpstream)
	if err != nil {
		return nil, err
	}

	// filter by uuid get only the right uuid
	if req.Uuid != nil {

		// filter by uuid to get only the right uuid
		filteredUpstreams := []*pb.UpstreamModel{}
		for _, upstream := range listPbUpstream {
			if upstream.Uuid == *req.Uuid {
				filteredUpstreams = append(filteredUpstreams, upstream)
			}
		}
		listPbUpstream = filteredUpstreams
	}

	// list the score
	listScore, err := r.kvDB.Zget(ctx, keyUpstreamScore, 0, -1)
	if err != nil {
		return nil, err
	}
	// loop the listscore and add the score to the upstream
	for _, upstream := range listPbUpstream {
		for _, score := range listScore {
			if score.Member == upstream.Uuid {
				upstream.Weight = &score.Score
				break
			}
		}
	}

	return listPbUpstream, nil
}

func (r *UpstreamRepository) Upsert(ctx context.Context, routeUuid string, upstream *pb.UpstreamModel) error {
	ctx, span := tracer.Start(ctx, "UpsertUpstreamRepository")
	defer tracer.SafeEndSpan(span)

	globalPrefix, rClient := r.kvDB.GetClient()
	scoreKey := fmt.Sprintf("%s%s:%s:%s", globalPrefix, routeRepo.BASE_KEY, routeUuid, BASE_KEY)
	keyUpstream := fmt.Sprintf("%s%s:%s:%s", globalPrefix, routeRepo.BASE_KEY, routeUuid, HASH_BASE_KEY)
	tx := rClient.TxPipeline()
	err := tx.ZAdd(ctx, scoreKey, redis.Z{Score: *upstream.Weight, Member: upstream.Uuid}).Err()
	if err != nil {
		return err
	}

	// remove weight from upstream
	upstream.Weight = nil
	serializedUpstream, err := proto.Marshal(upstream)
	if err != nil {
		return err
	}

	err = tx.HSet(ctx, keyUpstream, upstream.Uuid, serializedUpstream).Err()
	if err != nil {
		return err
	}

	// if the upstream has a rpc uuid, add the association to the rpc
	if upstream.RpcUuid != nil {
		keyRpc := fmt.Sprintf("%s%s:%s:%s", globalPrefix, rpcRepo.BASE_KEY, *upstream.RpcUuid, BASE_KEY)
		upstreamAssociation := &pb.UpstreamRPCRouteAssociation{
			Uuid:      upstream.Uuid,
			RpcUuid:   *upstream.RpcUuid,
			RouteUuid: routeUuid,
		}
		serializedUpstreamAssociation, err := proto.Marshal(upstreamAssociation)
		if err != nil {
			return err
		}
		err = tx.HSet(ctx, keyRpc, upstream.Uuid, serializedUpstreamAssociation).Err()
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec(ctx)
	return err
}

func (r *UpstreamRepository) UpdateScore(ctx context.Context, upstream *dto.UpstreamUpdateReq) error {
	ctx, span := tracer.Start(ctx, "UpdateScoreUpstreamRepository")
	defer tracer.SafeEndSpan(span)
	key := fmt.Sprintf("%s:%s:%s:%s", r.baseKey, routeRepo.BASE_KEY, upstream.RouteUuid, BASE_KEY)
	return r.kvDB.Zset(ctx, key, upstream.Uuid, upstream.Weight)
}

func (r *UpstreamRepository) Delete(ctx context.Context, req *dto.UpstreamDeleteReq) error {
	ctx, span := tracer.Start(ctx, "DeleteUpstreamRepository", trace.WithAttributes(attribute.String("app.upstream.uuid", req.Uuid)))
	defer tracer.SafeEndSpan(span)

	globalPrefix, rClient := r.kvDB.GetClient()
	scoreKey := fmt.Sprintf("%s%s:%s:%s", globalPrefix, routeRepo.BASE_KEY, req.RouteUuid, BASE_KEY)
	hashKey := fmt.Sprintf("%s%s:%s:%s", globalPrefix, routeRepo.BASE_KEY, req.RouteUuid, HASH_BASE_KEY)
	tx := rClient.TxPipeline()
	err := tx.HDel(ctx, hashKey, req.Uuid).Err()
	if err != nil {
		return err
	}
	err = tx.ZRem(ctx, scoreKey, req.Uuid).Err()
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx)
	return err
}
