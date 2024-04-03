package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/starton-io/tyrscale/manager/api/network/dto"
	"github.com/starton-io/tyrscale/manager/api/network/service/mocks"
	mockSvc "github.com/starton-io/tyrscale/manager/api/network/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type NetworkHandlerTestSuite struct {
	suite.Suite
	mockSvc  *mockSvc.INetworkService
	handler  *NetworkHandler
	fiberApp *fiber.App
}

func (suite *NetworkHandlerTestSuite) SetupTest() {
	_ = logger.InitLogger()
	suite.mockSvc = mocks.NewINetworkService(suite.T())
	suite.handler = NewNetworkHandler(suite.mockSvc, validation.New())
	suite.fiberApp = fiber.New()
}

func TestNetworkHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(NetworkHandlerTestSuite))
}

type serviceMock struct {
	res    any
	method string
}

type TestCreateNetwork struct {
	name       string
	input      any
	wantRes    *dto.CreateNetworkRes
	wantStatus int
	mockRes    *serviceMock
	skip       bool
}

func (suite *NetworkHandlerTestSuite) TestHandlersNetwork() {
	// Mock request and expected data
	req := &dto.Network{ChainId: 123, Name: "test", Blockchain: "test"}

	suite.fiberApp.Post("/networks", suite.handler.CreateNetwork)
	// Marshal body data
	tests := []TestCreateNetwork{
		{
			name:       "Create_Network_Success",
			input:      req,
			wantRes:    &dto.CreateNetworkRes{ChainId: 123},
			wantStatus: fiber.StatusCreated,
			mockRes:    &serviceMock{res: nil, method: "Create"},
		},
		{
			name:       "Create_Network_Failed",
			input:      &dto.Network{Name: "test", Blockchain: "test"},
			wantRes:    &dto.CreateNetworkRes{},
			wantStatus: fiber.StatusBadRequest,
			mockRes:    &serviceMock{res: nil, method: "Create"},
			skip:       true,
		},
		{
			name:       "Create_Network_Failed_2",
			input:      "string",
			wantRes:    &dto.CreateNetworkRes{},
			wantStatus: fiber.StatusBadRequest,
			mockRes:    &serviceMock{res: nil, method: "Create"},
			skip:       true,
		},
		{
			name:       "Create_Network_Failed_3",
			input:      req,
			wantRes:    &dto.CreateNetworkRes{},
			wantStatus: fiber.StatusConflict,
			mockRes:    &serviceMock{res: fmt.Errorf("network with the same name already exist"), method: "Create"},
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var got responses.CreatedSuccessResponse[dto.CreateNetworkRes]
			if !tt.skip {
				suite.mockSvc.On(tt.mockRes.method, mock.Anything, tt.input).Return(tt.mockRes.res).Times(1)
			}
			reqBody, _ := json.Marshal(tt.input)
			reqServer := httptest.NewRequest("POST", "/networks", bytes.NewReader(reqBody))
			reqServer.Header.Set("Content-Type", "application/json")
			resp, _ := suite.fiberApp.Test(reqServer, -1)
			//suite.T().Log(err)
			_ = json.NewDecoder(resp.Body).Decode(&got)
			assert.Equal(suite.T(), tt.wantStatus, resp.StatusCode)
			//assert.Equal(suite.T(), tt.wantErr, err)
			assert.Equal(suite.T(), tt.wantRes, got.Data)
		})
	}

}

type TestDeleteNetwork struct {
	name       string
	wantRes    *dto.DeleteNetworkRes
	path       string
	input      string
	wantStatus int
	mockRes    *serviceMock
}

func (suite *NetworkHandlerTestSuite) TestDeleteNetwork() {
	suite.fiberApp.Delete("/networks/:name", suite.handler.DeleteNetwork)
	tests := []TestDeleteNetwork{
		{
			name:       "Delete_Network_Success",
			path:       "/networks/eth-mainnet",
			input:      "eth-mainnet",
			wantRes:    &dto.DeleteNetworkRes{Name: "eth-mainnet"},
			wantStatus: fiber.StatusOK,
			mockRes:    &serviceMock{res: nil, method: "Delete"},
		},
		{
			name:       "Delete_Network_Failed_1",
			path:       "/networks/eth-mainnet",
			input:      "eth-mainnet",
			wantRes:    &dto.DeleteNetworkRes{},
			wantStatus: fiber.StatusNotFound,
			mockRes:    &serviceMock{res: fmt.Errorf("network not found"), method: "Delete"},
		},
		//{
		//	name:       "Delete_Network_Failed_2",
		//	path:       "/networks",
		//	input:      "",
		//	wantRes:    &dto.DeleteNetworkRes{},
		//	wantStatus: fiber.StatusBadRequest,
		//	mockRes:    &serviceMock{res: fmt.Errorf("network not found"), method: "Delete"},
		//},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			var got responses.General[dto.DeleteNetworkRes]
			suite.mockSvc.On(tt.mockRes.method, mock.Anything, tt.input).Return(tt.mockRes.res).Times(1)
			reqBody, _ := json.Marshal(tt.input)
			reqServer := httptest.NewRequest("DELETE", tt.path, bytes.NewReader(reqBody))
			reqServer.Header.Set("Content-Type", "application/json")
			resp, _ := suite.fiberApp.Test(reqServer, -1)
			_ = json.NewDecoder(resp.Body).Decode(&got)
			assert.Equal(suite.T(), tt.wantStatus, resp.StatusCode)
			assert.Equal(suite.T(), tt.wantRes, got.Data)
		})
	}

}
