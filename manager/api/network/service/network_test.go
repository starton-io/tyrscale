package service

import (
	"context"
	"testing"

	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/manager/api/network/dto"
	"github.com/starton-io/tyrscale/manager/api/network/repository/mocks"
	dtoRPC "github.com/starton-io/tyrscale/manager/api/rpc/dto"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/network"
	rpcPb "github.com/starton-io/tyrscale/manager/pkg/pb/rpc"

	mocksPubSub "github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub/mocks"
	recommendationDto "github.com/starton-io/tyrscale/manager/api/recommendation/dto"
	mocksRecommendationSvc "github.com/starton-io/tyrscale/manager/api/recommendation/service/mocks"
	mocksSvcRPC "github.com/starton-io/tyrscale/manager/api/rpc/service/mocks"
	recommendationPb "github.com/starton-io/tyrscale/manager/pkg/pb/recommendation"
	"github.com/stretchr/testify/suite"
)

type NetworkServiceTestSuite struct {
	suite.Suite
	mockNetRepo           *mocks.INetworkRepository
	mockSvcRPC            *mocksSvcRPC.IRPCService
	mockPubSub            *mocksPubSub.IPub
	mockRecommendationSvc *mocksRecommendationSvc.IRecommendationService
	service               INetworkService
}

func (suite *NetworkServiceTestSuite) SetupTest() {
	_ = logger.InitLogger()
	suite.mockNetRepo = mocks.NewINetworkRepository(suite.T())
	suite.mockPubSub = mocksPubSub.NewIPub(suite.T())
	suite.mockSvcRPC = mocksSvcRPC.NewIRPCService(suite.T())
	suite.mockRecommendationSvc = mocksRecommendationSvc.NewIRecommendationService(suite.T())
	suite.service = NewNetworkService(suite.mockNetRepo, suite.mockSvcRPC, suite.mockRecommendationSvc)
}

func TestNetworkServiceTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkServiceTestSuite))
}

func (suite *NetworkServiceTestSuite) TestCreateNetworkSucess() {
	ctx := context.Background()
	req := &dto.Network{
		ChainId:    1,
		Name:       "eth-mainnet",
		Blockchain: "ethereum",
	}

	suite.mockNetRepo.On("Create", ctx, &pb.NetworkModel{
		ChainId:    1,
		Name:       "eth-mainnet",
		Blockchain: "ethereum",
	}).Return(nil).Times(1)

	// suite list
	suite.mockNetRepo.On("List", ctx, &dto.ListNetworkReq{
		Name: "eth-mainnet",
	}).Return([]*pb.NetworkModel{}, nil)

	err := suite.service.Create(ctx, req)
	suite.Nil(err)
}

func (suite *NetworkServiceTestSuite) TestDeleteNetworkSucess() {
	ctx := context.Background()

	// suite delete
	suite.mockNetRepo.On("Delete", ctx, "eth-mainnet").Return(nil)
	res := []*rpcPb.RpcModel{
		{
			Uuid:        "cacc3b1a-11a4-43ad-b466-91950a53ade0",
			ChainId:     1,
			Url:         "https://eth-mainnet.gw.starton.io",
			NetworkName: "eth-mainnet",
			Type:        dtoRPC.RPCTypePrivate.String(),
			Collectors:  []string{"eth_block_number"},
		},
	}
	recommandationRes := []*recommendationPb.RecommendationModel{}

	// suite list
	suite.mockSvcRPC.On("List", ctx, &dtoRPC.ListReq{
		ListFilterReq: dtoRPC.ListFilterReq{
			NetworkName: "eth-mainnet",
		},
	}).Return(res, nil)

	for _, network := range res {
		suite.mockSvcRPC.On("Delete", ctx, &dtoRPC.DeleteRpcReq{
			UUID: network.Uuid,
		}).Return(nil)
	}

	suite.mockRecommendationSvc.On("List", ctx, &recommendationDto.ListRecommendationReq{
		NetworkName: "eth-mainnet",
	}).Return(recommandationRes, nil)

	//for _, recommendation := range recommandationRes {
	//	suite.mockRecommendationSvc.On("Delete", ctx, &recommendationDto.DeleteRecommendationReq{
	//		RouteUuid: recommendation.RouteUuid,
	//	}).Return(nil)
	//}

	suite.mockNetRepo.On("Delete", ctx, "eth-mainnet").Return(nil)
	err := suite.service.Delete(ctx, "eth-mainnet")
	suite.Nil(err)
}

func (suite *NetworkServiceTestSuite) TestListNetworkSucess() {
	ctx := context.Background()
	req := &dto.ListNetworkReq{
		Name:       "eth-mainnet",
		Blockchain: "ethereum",
	}
	res := []*pb.NetworkModel{
		{
			ChainId:    1,
			Name:       "eth-mainnet",
			Blockchain: "ethereum",
		},
	}
	suite.mockNetRepo.On("List", ctx, req).Return(res, nil)

	_, err := suite.service.List(ctx, req)
	suite.Nil(err)
}
