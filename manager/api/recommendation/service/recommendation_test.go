package service

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/starton-io/tyrscale/gateway/pkg/balancer"
	pubsubMocks "github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub/mocks"
	networkDto "github.com/starton-io/tyrscale/manager/api/network/dto"
	netMocks "github.com/starton-io/tyrscale/manager/api/network/repository/mocks"
	routeMocks "github.com/starton-io/tyrscale/manager/api/proxy/route/repository/mocks"
	recommendationDto "github.com/starton-io/tyrscale/manager/api/recommendation/dto"
	repoMocks "github.com/starton-io/tyrscale/manager/api/recommendation/repository/mocks"
	networkPb "github.com/starton-io/tyrscale/manager/pkg/pb/network"
	recommendationPb "github.com/starton-io/tyrscale/manager/pkg/pb/recommendation"
	routePb "github.com/starton-io/tyrscale/manager/pkg/pb/route"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RecommendationServiceTest struct {
	suite.Suite
	service     IRecommendationService
	netMocks    *netMocks.INetworkRepository
	pubsubMocks *pubsubMocks.IPub
	repoMocks   *repoMocks.IRecommendationRepository
	routeMocks  *routeMocks.IRouteRepository
}

func (suite *RecommendationServiceTest) SetupTest() {
	suite.netMocks = netMocks.NewINetworkRepository(suite.T())
	suite.pubsubMocks = pubsubMocks.NewIPub(suite.T())
	suite.repoMocks = repoMocks.NewIRecommendationRepository(suite.T())
	suite.routeMocks = routeMocks.NewIRouteRepository(suite.T())
	suite.service = NewRecommendationService(suite.repoMocks, suite.netMocks, suite.routeMocks, suite.pubsubMocks)
}

func TestRecommendationTestSuite(t *testing.T) {
	suite.Run(t, new(RecommendationServiceTest))
}

func (suite *RecommendationServiceTest) TestCreateRecommendation_Success() {
	ctx := context.Background()
	networkDto := &networkDto.ListNetworkReq{
		Name: "eth-mainnet",
	}
	//routeDto := &routeDto.ListRouteReq{
	//	Uuid: "test",
	//}
	recommendationReqDto := &recommendationDto.CreateRecommendationReq{
		Schedule:    "* * * * *",
		NetworkName: "eth-mainnet",
		Strategy:    "STRATEGY_HIGHEST_BLOCK",
	}
	netRes := []*networkPb.NetworkModel{
		{
			ChainId:    1,
			Name:       "eth-mainnet",
			Blockchain: "ethereum",
		},
	}
	suite.netMocks.On("List", ctx, networkDto).Return(netRes, nil)
	suite.routeMocks.On("List", ctx, mock.Anything).Return([]*routePb.RouteModel{
		{
			Uuid:                 "test",
			Host:                 "starton-local.com",
			Path:                 "/api/v1/test",
			LoadBalancerStrategy: string(balancer.BalancerPriority),
		},
	}, nil)
	suite.repoMocks.On("List", ctx, &recommendationDto.ListRecommendationReq{
		NetworkName: "eth-mainnet",
	}).Return([]*recommendationPb.RecommendationModel{}, nil)
	suite.repoMocks.On("Upsert", ctx, mock.Anything).Return(nil)
	suite.pubsubMocks.On("Publish", ctx, "recommendation_created", mock.Anything).Return(nil)
	recommendationRes, err := suite.service.Create(ctx, recommendationReqDto)
	suite.NoError(err)

	err = uuid.Validate(recommendationRes.RouteUuid)
	suite.NoError(err)
}

func (suite *RecommendationServiceTest) TestCreateRecommendation_NetworkNotFound() {
	ctx := context.Background()
	networkDto := &networkDto.ListNetworkReq{
		Name: "eth-mainnet",
	}
	recommendationReqDto := &recommendationDto.CreateRecommendationReq{
		Schedule:    "* * * * *",
		NetworkName: "eth-mainnet",
		Strategy:    "STRATEGY_HIGHEST_BLOCK",
	}
	netRes := make([]*networkPb.NetworkModel, 0)
	suite.netMocks.On("List", ctx, networkDto).Return(netRes, nil)
	_, err := suite.service.Create(ctx, recommendationReqDto)
	suite.Error(err)
}

func (suite *RecommendationServiceTest) TestDeleteRecommendation_Success() {
	ctx := context.Background()
	recommendationReqDto := &recommendationDto.DeleteRecommendationReq{
		RouteUuid: "test",
	}
	listRecommendationReqDto := &recommendationDto.ListRecommendationReq{
		RouteUuid: "test",
	}
	suite.repoMocks.On("List", ctx, listRecommendationReqDto).Return([]*recommendationPb.RecommendationModel{
		{
			RouteUuid: "test",
		},
	}, nil)
	suite.repoMocks.On("Delete", ctx, "test").Return(nil)
	suite.pubsubMocks.On("Publish", ctx, "recommendation_deleted", mock.Anything).Return(nil)
	err := suite.service.Delete(ctx, recommendationReqDto)
	suite.NoError(err)
}

func (suite *RecommendationServiceTest) TestListRecommendation_Success() {
	ctx := context.Background()
	listRecommendationReqDto := &recommendationDto.ListRecommendationReq{
		RouteUuid: "test",
	}
	suite.repoMocks.On("List", ctx, listRecommendationReqDto).Return([]*recommendationPb.RecommendationModel{
		{
			RouteUuid: "test",
		},
	}, nil)
	recommendationRes, err := suite.service.List(ctx, listRecommendationReqDto)
	suite.NoError(err)
	suite.Equal(1, len(recommendationRes))
}

func (suite *RecommendationServiceTest) TestUpdateRecommendation_Success() {
	ctx := context.Background()
	listRecommendationReqDto := &recommendationDto.ListRecommendationReq{
		RouteUuid: "test",
	}
	listRecommendationRes := []*recommendationPb.RecommendationModel{
		{
			RouteUuid:   "test",
			Schedule:    "* * * * *",
			NetworkName: "eth-mainnet",
			Strategy:    recommendationDto.StrategyHighestBlock.String(),
		},
	}
	updateRecommendationReqDto := &recommendationDto.UpdateRecommendationReq{
		RouteUuid: "test",
		Strategy:  "STRATEGY_AIR_UNDER_THE_CURVE",
	}
	updateRecommendation := listRecommendationRes[0]
	_ = copier.CopyWithOption(updateRecommendation, updateRecommendationReqDto, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
	suite.repoMocks.On("List", ctx, listRecommendationReqDto).Return(listRecommendationRes, nil)
	suite.repoMocks.On("Upsert", ctx, updateRecommendation).Return(nil)
	suite.pubsubMocks.On("Publish", ctx, "recommendation_updated", mock.Anything).Return(nil)
	err := suite.service.Update(ctx, updateRecommendationReqDto)
	suite.NoError(err)
}
