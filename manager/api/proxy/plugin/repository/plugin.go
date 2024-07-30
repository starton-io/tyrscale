package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/manager/api/proxy/plugin/dto"
	routeRepo "github.com/starton-io/tyrscale/manager/api/proxy/route/repository"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"

	//"github.com/starton-io/tyrscale/manager/api/proxy/upstream/dto"
	pb "github.com/starton-io/tyrscale/gateway/proto/gen/go/plugin"
)

const (
	BASE_KEY = "plugins"
)

type PluginRepository struct {
	baseKey string
	kvDB    kv.IRedisStore
}

//go:generate mockery --name=IPluginRepository
type IPluginRepository interface {
	List(ctx context.Context, req *dto.ListPluginReq) ([]*pb.Plugin, error)
	Upsert(ctx context.Context, routeUuid string, plugin *pb.Plugin) error
	Delete(ctx context.Context, routeUuid string, req *dto.DetachPluginReq) error
}

func NewPluginRepository(kvDB kv.IRedisStore) IPluginRepository {
	return &PluginRepository{
		baseKey: BASE_KEY,
		kvDB:    kvDB,
	}
}

func (r *PluginRepository) List(ctx context.Context, req *dto.ListPluginReq) ([]*pb.Plugin, error) {
	ctx, span := tracer.Start(ctx, "ListPluginRepository")
	defer tracer.SafeEndSpan(span)
	matchCriteria, err := utils.StructToMapStr(req, "query")
	if err != nil {
		return nil, err
	}
	filter := kv.NewParamsFilterPB[*pb.Plugin](matchCriteria, true, "", 0)
	if matchCriteria["name"] != "" {
		filter.MatchCriteria["name"] = matchCriteria["name"]
		filter.EnablePrefilter = true
	}

	keyPuginHash := fmt.Sprintf("%s:%s:%s", routeRepo.BASE_KEY, req.RouteUuid, BASE_KEY)
	listPlugin, err := r.kvDB.ScanHash(ctx, keyPuginHash, filter)
	if err != nil {
		return nil, err
	}

	return utils.UnmarshalSliceBytesToProto[*pb.Plugin](listPlugin)

	//var res = &dto.Plugins{
	//	InterceptorRequest:  make([]*dto.Plugin, 0),
	//	Middleware:          make([]*dto.Plugin, 0),
	//	InterceptorResponse: make([]*dto.Plugin, 0),
	//}

	//for _, pluginResp := range listPbPlugin {
	//	var payloadConfig map[string]interface{}
	//	if err := json.Unmarshal(pluginResp.Config, &payloadConfig); err != nil {
	//		return nil, fmt.Errorf("failed to unmarshal plugin config: %w", err)
	//	}
	//	switch pluginResp.Name {
	//	case string(plugin.PluginTypeMiddleware):
	//		res.Middleware = append(res.Middleware, &dto.Plugin{
	//			Name:        pluginResp.Name,
	//			Description: pluginResp.Description,
	//			Config:      pluginResp.Config,
	//			Priority:    int(pluginResp.Priority),
	//		})
	//	case string(plugin.PluginTypeRequestInterceptor):
	//		res.InterceptorRequest = append(res.InterceptorRequest, &dto.Plugin{
	//			Name:        pluginResp.Name,
	//			Description: pluginResp.Description,
	//			Config:      pluginResp.Config,
	//			Priority:    int(pluginResp.Priority),
	//		})
	//	case string(plugin.PluginTypeResponseInterceptor):
	//		res.InterceptorResponse = append(res.InterceptorResponse, &dto.Plugin{
	//			Name:        pluginResp.Name,
	//			Description: pluginResp.Description,
	//			Config:      pluginResp.Config,
	//			Priority:    int(pluginResp.Priority),
	//		})
	//	}
	//}
	//return res, nil
}

func (r *PluginRepository) Upsert(ctx context.Context, routeUuid string, plugin *pb.Plugin) error {
	ctx, span := tracer.Start(ctx, "UpsertPluginRepository", trace.WithAttributes(attribute.String("app.route.uuid", routeUuid)))
	defer tracer.SafeEndSpan(span)

	keyPuginHash := fmt.Sprintf("%s:%s:%s", routeRepo.BASE_KEY, routeUuid, r.baseKey)
	serializedValue, err := proto.Marshal(plugin)
	if err != nil {
		return err
	}
	return r.kvDB.StoreHash(ctx, keyPuginHash, plugin.Name, serializedValue)
}

func (r *PluginRepository) Delete(ctx context.Context, routeUuid string, req *dto.DetachPluginReq) error {
	ctx, span := tracer.Start(ctx, "DeletePluginRepository", trace.WithAttributes(attribute.String("app.route.UUID", routeUuid)))
	defer tracer.SafeEndSpan(span)
	log.Printf("DeletePluginRepository: %s", req.Name)

	hashKey := fmt.Sprintf("%s:%s:%s", routeRepo.BASE_KEY, routeUuid, r.baseKey)
	return r.kvDB.DeleteHash(ctx, hashKey, req.Name)
}
