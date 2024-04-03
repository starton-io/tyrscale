package repository

import (
	"context"

	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/manager/api/network/dto"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/network"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

const (
	BASE_KEY = "networks"
)

// INetworkRepository is an interface for network repository
//
//go:generate mockery --name=INetworkRepository
type INetworkRepository interface {
	Create(ctx context.Context, network *pb.NetworkModel) error
	List(ctx context.Context, filterParams *dto.ListNetworkReq) ([]*pb.NetworkModel, error)
	Delete(ctx context.Context, name string) error
}

type NetworkRepository struct {
	baseKey string
	kvDB    kv.IRedisStore
}

func NewNetworkRepository(kvDB kv.IRedisStore) INetworkRepository {
	return &NetworkRepository{
		baseKey: BASE_KEY,
		kvDB:    kvDB,
	}
}

func (r *NetworkRepository) Create(ctx context.Context, network *pb.NetworkModel) error {
	ctx, span := tracer.Start(ctx, "CreateNetworkRepository", trace.WithAttributes(attribute.String("app.network.name", network.Name)))
	defer tracer.SafeEndSpan(span)

	key := r.baseKey
	serializedValue, err := proto.Marshal(network)
	if err != nil {
		return err
	}
	return r.kvDB.StoreHash(ctx, key, network.Name, serializedValue)
}

func (r *NetworkRepository) List(ctx context.Context, filterParams *dto.ListNetworkReq) ([]*pb.NetworkModel, error) {
	ctx, span := tracer.Start(ctx, "ListNetworkRepository")
	defer tracer.SafeEndSpan(span)

	matchCriteria, err := utils.StructToMapStr(filterParams, "query")
	if err != nil {
		return nil, err
	}
	filter := kv.NewParamsFilterPB[*pb.NetworkModel](matchCriteria, false, "", 0)
	if matchCriteria["name"] != "" {
		filter.EnablePrefilter = true
		filter.PrefilterPattern = matchCriteria["name"]
	}
	networksMap, err := r.kvDB.ScanHash(ctx, r.baseKey, filter)
	if err != nil {
		return nil, err
	}
	return utils.UnmarshalSliceBytesToProto[*pb.NetworkModel](networksMap)
}

func (r *NetworkRepository) Delete(ctx context.Context, name string) error {
	ctx, span := tracer.Start(ctx, "DeleteNetworkRepository", trace.WithAttributes(attribute.String("app.network.name", name)))
	defer tracer.SafeEndSpan(span)

	return r.kvDB.DeleteHash(ctx, r.baseKey, name)
}
