package repository

import (
	"context"
	"sort"
	"testing"

	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	mocks "github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv/redismocks"

	"github.com/starton-io/tyrscale/manager/api/network/dto"
	"github.com/starton-io/tyrscale/manager/api/network/model"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/network"
	"github.com/stretchr/testify/suite"
	"google.golang.org/protobuf/proto"
)

type NetworkRepositoryTestSuite struct {
	suite.Suite
	mockV *mocks.IRedisStore
	repo  INetworkRepository
}

func (suite *NetworkRepositoryTestSuite) SetupTest() {
	suite.mockV = mocks.NewIRedisStore(suite.T())
	suite.repo = NewNetworkRepository(suite.mockV)
}

// func suite run
func TestNetworkRepositoryRunSuite(t *testing.T) {
	suite.Run(t, new(NetworkRepositoryTestSuite))
}

func (suite *NetworkRepositoryTestSuite) TestCreateNetworkSuccess() {
	network := &pb.NetworkModel{ChainId: 1, Name: "test", Blockchain: "test"}
	key := BASE_KEY
	ctx := context.Background()
	serializedValue, err := proto.Marshal(network)
	suite.Nil(err)

	// Adjust mock expectations to include context
	suite.mockV.On("StoreHash", ctx, key, network.Name, serializedValue).Return(nil)

	// Test
	err = suite.repo.Create(ctx, network)
	suite.Nil(err)
}

func (suite *NetworkRepositoryTestSuite) TestDeleteNetwork() {
	newNetwork := &model.Network{ChainID: 1, Name: "test", Blockchain: "test"}
	key := BASE_KEY
	ctx := context.Background()
	suite.mockV.On("DeleteHash", ctx, key, newNetwork.Name).Return(nil)

	// Test
	err := suite.repo.Delete(ctx, newNetwork.Name)
	suite.Nil(err)
}

func (suite *NetworkRepositoryTestSuite) TestListNetworksPB() {
	key := BASE_KEY
	ctx := context.Background()
	filter := kv.NewParamsFilterPB[*pb.NetworkModel](map[string]string{}, false, "", 0)

	result1 := &pb.NetworkModel{ChainId: 1137, Name: "test", Blockchain: "test"}
	result1Bytes, err := proto.Marshal(result1)
	suite.Nil(err)
	result2 := &pb.NetworkModel{ChainId: 2, Name: "test", Blockchain: "test"}
	result2Bytes, err := proto.Marshal(result2)
	suite.Nil(err)
	result := [][]byte{result1Bytes, result2Bytes}
	suite.mockV.On("ScanHash", ctx, key, filter).Return(result, nil)

	// Test
	networks, err := suite.repo.List(ctx, &dto.ListNetworkReq{})
	// sorted by chain_id
	sort.Slice(networks, func(i, j int) bool {
		return networks[i].ChainId < networks[j].ChainId
	})
	suite.Nil(err)
	suite.Equal(2, len(networks))
	suite.Equal(int64(2), networks[0].ChainId)
	suite.Equal("test", networks[0].Name)
	suite.Equal("test", networks[0].Blockchain)
	suite.Equal(int64(1137), networks[1].ChainId)
	suite.Equal("test", networks[1].Name)
	suite.Equal("test", networks[1].Blockchain)
}
