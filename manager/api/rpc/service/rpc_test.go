package service

import (
	"context"
	"testing"

	"github.com/jinzhu/copier"
	mocksPubSub "github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/pubsub/mocks"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/tracer"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	dtoNet "github.com/starton-io/tyrscale/manager/api/network/dto"
	mocksNet "github.com/starton-io/tyrscale/manager/api/network/repository/mocks"
	mocksUpstream "github.com/starton-io/tyrscale/manager/api/proxy/upstream/service/mocks"
	"github.com/starton-io/tyrscale/manager/api/rpc/dto"
	"github.com/starton-io/tyrscale/manager/api/rpc/repository/mocks"
	netPb "github.com/starton-io/tyrscale/manager/pkg/pb/network"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/rpc"
	"github.com/starton-io/tyrscale/manager/pkg/pb/upstream"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RPCServiceTestSuite struct {
	suite.Suite
	service       IRPCService
	mocks         *mocks.IRPCRepository
	mocksNet      *mocksNet.INetworkRepository
	mocksUpstream *mocksUpstream.IUpstreamService
	mocksPubSub   *mocksPubSub.IPub
}

func (suite *RPCServiceTestSuite) SetupTest() {
	_ = logger.InitLogger()
	tracer.InitTracer("test")
	suite.mocks = mocks.NewIRPCRepository(suite.T())
	suite.mocksNet = mocksNet.NewINetworkRepository(suite.T())
	suite.mocksUpstream = mocksUpstream.NewIUpstreamService(suite.T())
	suite.mocksPubSub = mocksPubSub.NewIPub(suite.T())

	suite.service = NewRPCService(suite.mocks, suite.mocksNet, suite.mocksUpstream, suite.mocksPubSub)
}

func TestRPCServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RPCServiceTestSuite))
}

// test create rpc
func (suite *RPCServiceTestSuite) TestCreateRPC() {
	ctx := context.Background()
	var rpcModel pb.RpcModel
	rpc := &dto.CreateRpcReq{
		UUID:        "test-rpc",
		NetworkName: "test-network",
		Collectors:  []string{"eth_block_number"},
		Type:        "public",
		URL:         "https://test-rpc.com",
		Provider:    "test-provider",
	}
	netRes := []*netPb.NetworkModel{
		{
			ChainId:    1,
			Name:       "test-network",
			Blockchain: "ethereum",
		},
	}

	utils.Copy(&rpcModel, rpc)
	rpcModel.ChainId = netRes[0].ChainId
	suite.mocksNet.On("List", mock.Anything, &dtoNet.ListNetworkReq{
		Name: "test-network",
	}).Return(netRes, nil)

	suite.mocks.On("Create", mock.Anything, &rpcModel).Return(nil)
	suite.mocksPubSub.On("Publish", mock.Anything, "rpc_created", mock.Anything).Return(nil)
	err := suite.service.Create(ctx, rpc)
	suite.Nil(err)
}

func (suite *RPCServiceTestSuite) TestDeleteRPC() {
	ctx := context.Background()

	listRes := []*pb.RpcModel{
		{
			Uuid:        "test-rpc",
			NetworkName: "test-network",
			Collectors:  []string{"eth_block_number"},
			Type:        dto.RPCTypePublic.String(),
			Url:         "https://test-rpc.com",
			Provider:    "test-provider",
			ChainId:     1,
		},
	}
	dtoList := &dto.ListReq{
		ListFilterReq: dto.ListFilterReq{
			UUID: "test-rpc",
		},
	}
	dtoDelete := &dto.DeleteRpcReq{
		UUID: "test-rpc",
	}
	suite.mocks.On("List", mock.Anything, dtoList).Return(listRes, nil)
	suite.mocks.On("ListAssociatedUpstreamByRPCUuid", mock.Anything, listRes[0].Uuid).Return([]*upstream.UpstreamRPCRouteAssociation{}, nil)
	suite.mocks.On("Delete", mock.Anything, listRes[0]).Return(nil)
	suite.mocksPubSub.On("Publish", mock.Anything, "rpc_deleted", mock.Anything).Return(nil)
	err := suite.service.Delete(ctx, dtoDelete)
	suite.Nil(err)
}

func (suite *RPCServiceTestSuite) TestUpdateRPC() {
	ctx := context.Background()

	dtoList := &dto.ListReq{
		ListFilterReq: dto.ListFilterReq{
			UUID: "test-rpc",
		},
	}
	listRes := []*pb.RpcModel{
		{
			Uuid:        "test-rpc",
			NetworkName: "test-network",
			Collectors:  []string{"eth_block_number"},
			Type:        dto.RPCTypePublic.String(),
			Url:         "https://test-rpc.com",
			Provider:    "test-provider",
			ChainId:     1,
		},
	}

	private := dto.RPCTypePrivate
	dtoUpdate := &dto.UpdateRpcReq{
		UUID: "test-rpc",
		Type: &private,
	}
	updateRpc := listRes[0]
	_ = copier.CopyWithOption(updateRpc, dtoUpdate, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})

	suite.mocks.On("List", mock.Anything, dtoList).Return(listRes, nil)
	suite.mocks.On("Update", mock.Anything, updateRpc).Return(nil)
	suite.mocksPubSub.On("Publish", mock.Anything, "rpc_updated", mock.Anything).Return(nil)
	err := suite.service.Update(ctx, dtoUpdate)
	suite.Nil(err)
}
