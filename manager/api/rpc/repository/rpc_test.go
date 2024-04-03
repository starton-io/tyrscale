package repository

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
	"github.com/redis/go-redis/v9"
	"github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv"
	mocks "github.com/starton-io/tyrscale/go-kit/pkg/infrastructure/kv/redismocks"
	"github.com/starton-io/tyrscale/manager/api/rpc/dto"
	pb "github.com/starton-io/tyrscale/manager/pkg/pb/rpc"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/proto"

	"github.com/stretchr/testify/suite"
)

type RPCRepositoryTestSuite struct {
	suite.Suite
	mockKV *mocks.IRedisStore

	repo        IRPCRepository
	mocksRedis  redismock.ClientMock
	redisClient *redis.Client
}

// MockRedisClient is a mock of Redis client
type MockRedisClient struct {
	mock.Mock
}

func (suite *RPCRepositoryTestSuite) SetupTest() {
	suite.mockKV = mocks.NewIRedisStore(suite.T())
	suite.redisClient, suite.mocksRedis = redismock.NewClientMock()
	suite.repo = NewRPCRepository(suite.mockKV)

}

func TestRPCRepositoryRunSuite(t *testing.T) {
	suite.Run(t, new(RPCRepositoryTestSuite))
}

func (suite *RPCRepositoryTestSuite) TestCreateRPCSuccess() {
	rpc := &pb.RpcModel{Uuid: "123", ChainId: 1, NetworkName: "test", Url: "test", Type: dto.RPCTypePublic.String(), Provider: "test", Collectors: []string{"test"}}
	rpcBytes, err := proto.Marshal(rpc)
	suite.Nil(err)
	ctx := context.Background()
	globalPrefix := "baseKey:"
	rpcKey := globalPrefix + "rpcs" // Assuming ':' is used as a delimiter in keys

	// Mocking kv pkg operations
	suite.mockKV.On("GetClient").Return(globalPrefix, suite.redisClient)

	// Mocking redis operations
	suite.mocksRedis.ExpectHExists(rpcKey, rpc.Uuid).SetVal(false)
	suite.mocksRedis.ExpectTxPipeline()
	suite.mocksRedis.ExpectHSet(rpcKey, rpc.Uuid, rpcBytes).SetVal(1)
	suite.mocksRedis.ExpectZAdd("baseKey:networks:test:rpcs", redis.Z{Score: 0, Member: rpc.Uuid}).SetVal(1)
	suite.mocksRedis.ExpectTxPipelineExec()

	// Call the method under test
	err = suite.repo.Create(ctx, rpc)
	suite.Nil(err) // Check no error returned

	// Assertions
	err = suite.mocksRedis.ExpectationsWereMet()
	suite.Nil(err)
}

func (suite *RPCRepositoryTestSuite) TestDeleteRPCSuccess() {
	rpc := &pb.RpcModel{Uuid: "123", ChainId: 1, NetworkName: "test", Url: "test", Type: "public", Provider: "test", Collectors: []string{"test"}}
	ctx := context.Background()
	globalPrefix := "baseKey:"
	rpcKey := globalPrefix + "rpcs" // Assuming ':' is used as a delimiter in keys

	// Mocking kv pkg operations
	suite.mockKV.On("GetClient").Return(globalPrefix, suite.redisClient)

	// Mocking redis operations
	suite.mocksRedis.ExpectTxPipeline()
	suite.mocksRedis.ExpectHDel(rpcKey, "123").SetVal(1)
	suite.mocksRedis.ExpectZRem("baseKey:networks:test:rpcs", "123").SetVal(1)
	suite.mocksRedis.ExpectTxPipelineExec()

	err := suite.repo.Delete(ctx, rpc)
	suite.Nil(err)
}

func (suite *RPCRepositoryTestSuite) TestListRPCs() {
	key := "rpcs"
	ctx := context.Background()
	filter := kv.NewParamsFilterPB[*pb.RpcModel](make(map[string]string), false, "", 0)
	rpc1 := &pb.RpcModel{Uuid: "123", ChainId: 1, NetworkName: "123", Url: "test", Type: dto.RPCTypePublic.String(), Provider: "test", Collectors: []string{"test"}}
	rpc1Bytes, err := proto.Marshal(rpc1)
	suite.Nil(err)
	rpc2 := &pb.RpcModel{Uuid: "124", ChainId: 1, NetworkName: "124", Url: "test", Type: dto.RPCTypePublic.String(), Provider: "test", Collectors: []string{"test"}}
	rpc2Bytes, err := proto.Marshal(rpc2)
	suite.Nil(err)
	expectedRpcs := [][]byte{
		rpc1Bytes,
		rpc2Bytes,
	}
	suite.mockKV.On("ScanHash", ctx, key, filter).Return(expectedRpcs, nil)

	rpcs, err := suite.repo.List(ctx, &dto.ListReq{ListSortReq: &dto.ListSortReq{SortBy: "network_name", SortDescending: false}})
	suite.Nil(err) // Check no error returned

	// check the correct order of rpcs (by networkName and URL)
	suite.Equal(2, len(rpcs))
	suite.Equal("123", rpcs[0].Uuid)
	suite.Equal(int64(1), rpcs[0].ChainId)
	suite.Equal("123", rpcs[0].NetworkName)
	suite.Equal("test", rpcs[0].Url)
	suite.Equal(dto.RPCTypePublic.String(), rpcs[0].Type)
	suite.Equal("test", rpcs[0].Provider)
	suite.Equal([]string{"test"}, rpcs[0].Collectors)
	suite.Equal("124", rpcs[1].Uuid)
	suite.Equal(int64(1), rpcs[1].ChainId)
	suite.Equal("124", rpcs[1].NetworkName)
	suite.Equal("test", rpcs[1].Url)
	suite.Equal(dto.RPCTypePublic.String(), rpcs[1].Type)
	suite.Equal("test", rpcs[1].Provider)
	suite.Equal([]string{"test"}, rpcs[1].Collectors)
}
