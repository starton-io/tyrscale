package repository

import (
	"context"

	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/manager/api/recommendation/dto"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/recommendation"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

const (
	BASE_KEY = "recommendations"
)

//go:generate mockery --name=IRecommendationRepository
type IRecommendationRepository interface {
	Upsert(ctx context.Context, recommendation *pb.RecommendationModel) error
	List(ctx context.Context, filterParams *dto.ListRecommendationReq) ([]*pb.RecommendationModel, error)
	Delete(ctx context.Context, uuid string) error
}

type RecommendationRepository struct {
	baseKey string
	kvDB    kv.IRedisStore
}

func NewRecommendationRepository(kvDB kv.IRedisStore) IRecommendationRepository {
	return &RecommendationRepository{
		baseKey: BASE_KEY,
		kvDB:    kvDB,
	}
}

func (r *RecommendationRepository) Upsert(ctx context.Context, recommendation *pb.RecommendationModel) error {
	ctx, span := tracer.Start(ctx, "CreateRecommendationRepository", trace.WithAttributes(attribute.String("app.recommendation.route_uuid", recommendation.RouteUuid)))
	defer tracer.SafeEndSpan(span)
	key := r.baseKey
	serializedValue, err := proto.Marshal(recommendation)
	if err != nil {
		return err
	}
	return r.kvDB.StoreHash(ctx, key, recommendation.RouteUuid, serializedValue)
}

func (r *RecommendationRepository) List(ctx context.Context, filterParams *dto.ListRecommendationReq) ([]*pb.RecommendationModel, error) {
	ctx, span := tracer.Start(ctx, "ListRecommendationRepository", trace.WithAttributes(attribute.String("app.recommendation.route_uuid", filterParams.RouteUuid)))
	defer tracer.SafeEndSpan(span)
	matchCriteria, err := utils.StructToMapStr(filterParams, "query")
	if err != nil {
		return nil, err
	}
	filter := kv.NewParamsFilterPB[*pb.RecommendationModel](matchCriteria, false, "", 0)
	if matchCriteria["route_uuid"] != "" {
		filter.MatchCriteria["route_uuid"] = matchCriteria["route_uuid"]
		filter.EnablePrefilter = true
	}

	recommendations, err := r.kvDB.ScanHash(ctx, r.baseKey, filter)
	if err != nil {
		return nil, err
	}
	return utils.UnmarshalSliceBytesToProto[*pb.RecommendationModel](recommendations)
}

func (r *RecommendationRepository) Delete(ctx context.Context, uuid string) error {
	key := r.baseKey
	return r.kvDB.DeleteHash(ctx, key, uuid)
}
