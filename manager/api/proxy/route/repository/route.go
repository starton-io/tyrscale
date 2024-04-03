package repository

import (
	"context"
	"fmt"

	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/dto"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/route"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

var (

	// TODO: use Lua script to retrieve all upstreams for each route
	LuaScriptGetAll = `
    local cursor = "0"
    local result = {}
    repeat
        local scan_result = redis.call("SCAN", cursor, "MATCH", "route:*")
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
	LuaScriptGettest = `
    local cursor = "0"
    local result = {}
    repeat
        local scan_result = redis.call("SCAN", cursor, "MATCH", "routes:*")
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

	LuaScriptDeleteUpstreamsOfRoute = `
    local cursor = "0"
    local result = {}
    repeat
        local scan_result = redis.call("SCAN", cursor, "MATCH", "route:" .. ARGV[1])
        cursor = scan_result[1]
        local keys = scan_result[2]
        for i, key in ipairs(keys) do
            local upstreams = redis.call("ZRANGE", key, 0, -1, "WITHSCORES")
			// deletes each key upstreams in hset upstreams
			for _, upstream in ipairs(upstreams) do
				redis.call("HDEL", "upstreams", upstream)
			end
        end
    until cursor == "0"
    return result
	`
)

const (
	BASE_KEY      = "routes"
	UPSTREAMS_KEY = "upstreams"
)

type RouteRepository struct {
	baseKey string
	kvDB    kv.IRedisStore
}

//go:generate mockery --name=IRouteRepository
type IRouteRepository interface {
	Upsert(ctx context.Context, route *pb.RouteModel) error
	List(ctx context.Context, filterParams *dto.ListRouteReq) ([]*pb.RouteModel, error)
	Delete(ctx context.Context, route *pb.RouteModel) error
}

func NewRouteRepository(kvDB kv.IRedisStore) *RouteRepository {
	return &RouteRepository{
		baseKey: BASE_KEY,
		kvDB:    kvDB,
	}
}

func (r *RouteRepository) Upsert(ctx context.Context, route *pb.RouteModel) error {
	ctx, span := tracer.Start(ctx, "UpsertRouteRepository", trace.WithAttributes(attribute.String("app.route.uuid", route.Uuid)))
	defer tracer.SafeEndSpan(span)

	key := r.baseKey
	serializedValue, err := proto.Marshal(route)
	if err != nil {
		return err
	}
	return r.kvDB.StoreHash(ctx, key, route.Uuid, serializedValue)
}

func (r *RouteRepository) List(ctx context.Context, filterParams *dto.ListRouteReq) ([]*pb.RouteModel, error) {
	ctx, span := tracer.Start(ctx, "ListRouteRepository")
	defer tracer.SafeEndSpan(span)
	matchCriteria, err := utils.StructToMapStr(filterParams, "query")
	if err != nil {
		return nil, err
	}
	fmt.Println(matchCriteria)
	filter := kv.NewParamsFilterPB[*pb.RouteModel](matchCriteria, false, "", 0)
	if matchCriteria["uuid"] != "" {
		filter.EnablePrefilter = true
		filter.PrefilterPattern = matchCriteria["uuid"]
	}
	routeMap, err := r.kvDB.ScanHash(ctx, r.baseKey, filter)
	if err != nil {
		return nil, err
	}
	return utils.UnmarshalSliceBytesToProto[*pb.RouteModel](routeMap)
}

func (r *RouteRepository) Delete(ctx context.Context, route *pb.RouteModel) error {
	ctx, span := tracer.Start(ctx, "DeleteRouteRepository", trace.WithAttributes(attribute.String("app.route.UUID", route.Uuid)))
	defer tracer.SafeEndSpan(span)

	return r.kvDB.DeleteHash(ctx, r.baseKey, route.Uuid)
}
