package repository

import (
	"context"
	"testing"

	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	mocks "github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv/redismocks"
	"github.com/starton-io/tyrscale/manager/api/recommendation/dto"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/recommendation"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
)

type RecommendationRepositoryTestSuite struct {
	suite.Suite
	mockDB *mocks.IRedisStore
	repo   IRecommendationRepository
}

func (suite *RecommendationRepositoryTestSuite) SetupTest() {
	suite.mockDB = mocks.NewIRedisStore(suite.T())
	suite.repo = NewRecommendationRepository(suite.mockDB)
}

func TestRecommendationRepositoryRunSuite(t *testing.T) {
	suite.Run(t, new(RecommendationRepositoryTestSuite))
}

func (suite *RecommendationRepositoryTestSuite) TestUpsert() {
	recommendation := &pb.RecommendationModel{
		RouteUuid:   "test",
		Schedule:    "* * * * *",
		NetworkName: "test",
		Strategy:    dto.StrategyHighestBlock.String(),
	}
	ctx := context.Background()
	serializedValueBytes, err := proto.Marshal(recommendation)
	suite.Nil(err)
	suite.mockDB.On("StoreHash", ctx, "recommendations", recommendation.RouteUuid, serializedValueBytes).Return(nil)
	err = suite.repo.Upsert(ctx, recommendation)
	suite.NoError(err)
}

func (suite *RecommendationRepositoryTestSuite) TestDelete() {
	recommendation := &pb.RecommendationModel{
		RouteUuid: "test",
	}
	ctx := context.Background()
	suite.mockDB.On("DeleteHash", ctx, "recommendations", recommendation.RouteUuid).Return(nil)
	err := suite.repo.Delete(ctx, recommendation.RouteUuid)
	suite.Nil(err)
}

func (suite *RecommendationRepositoryTestSuite) TestList() {
	recommendationRes := &pb.RecommendationModel{
		RouteUuid:   "test",
		Schedule:    "* * * * *",
		NetworkName: "test",
		Strategy:    dto.StrategyHighestBlock.String(),
	}
	serializedValueBytes, err := proto.Marshal(recommendationRes)
	suite.Nil(err)
	responses := [][]byte{
		serializedValueBytes,
	}
	listRecommendationsParams := &dto.ListRecommendationReq{
		RouteUuid: "test",
	}
	filterParams := kv.NewParamsFilterPB[*pb.RecommendationModel](map[string]string{"route_uuid": "test"}, true, "", 0)

	ctx := context.Background()
	suite.mockDB.On("ScanHash", ctx, "recommendations", filterParams).Return(responses, nil)
	recommendations, err := suite.repo.List(ctx, listRecommendationsParams)
	suite.Nil(err)
	suite.NotNil(recommendations)
	suite.Equal(1, len(recommendations))
	suite.Equal("test", recommendations[0].RouteUuid)
	suite.Equal("* * * * *", recommendations[0].Schedule)
	suite.Equal("test", recommendations[0].NetworkName)
	suite.Equal(dto.StrategyHighestBlock.String(), recommendations[0].Strategy)
}
