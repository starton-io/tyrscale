package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/starton-io/tyrscale/manager/api/recommendation/dto"
	"github.com/starton-io/tyrscale/manager/api/recommendation/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RecommendationHandlerTestSuite struct {
	suite.Suite
	svcMock  *mocks.IRecommendationService
	handler  *RecommendationHandler
	fiberApp *fiber.App
}

func (suite *RecommendationHandlerTestSuite) SetupTest() {
	_ = logger.InitLogger()
	suite.svcMock = mocks.NewIRecommendationService(suite.T())
	suite.handler = NewHandler(suite.svcMock, validation.New())
	suite.fiberApp = fiber.New()
}

func TestRecommendationHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(RecommendationHandlerTestSuite))
}

type serviceMock struct {
	res    any
	method string
}

type testCreateRecommendation struct {
	name       string
	input      any
	wantRes    *dto.CreateRecommendationRes
	wantStatus int
	resMock    *serviceMock
	skipMock   bool
}

func (suite *RecommendationHandlerTestSuite) TestCreateRecommendation() {

	req := &dto.CreateRecommendationReq{
		RouteUuid:   "c3edbfad-69ff-4078-bb12-5e39fbddd988",
		Schedule:    "* * * * *",
		NetworkName: "eth-mainnet",
		Strategy:    "STRATEGY_CUSTOM",
	}
	suite.fiberApp.Post("/recommendations", suite.handler.CreateRecommendation)
	tests := []testCreateRecommendation{
		{
			name:  "Create_Recommendation_Success",
			input: req,
			wantRes: &dto.CreateRecommendationRes{
				RouteUuid: "test",
			},
			wantStatus: http.StatusCreated,
			resMock:    &serviceMock{res: nil, method: "Create"},
		},
		{
			name:       "Create_Recommendation_BadRequest_BodyParser",
			input:      "wrongbody",
			wantRes:    &dto.CreateRecommendationRes{},
			skipMock:   true,
			wantStatus: http.StatusBadRequest,
			resMock:    &serviceMock{res: nil, method: "Create"},
		},
		{
			name: "Create_Recommendation_BadRequest_Validator",
			input: &dto.CreateRecommendationReq{
				Schedule:    "unknown",
				NetworkName: "eth-mainnet",
				Strategy:    "strategy-highest-block",
			},
			wantRes:    &dto.CreateRecommendationRes{},
			skipMock:   true,
			wantStatus: http.StatusBadRequest,
			resMock:    &serviceMock{res: nil, method: "Create"},
		},
		{
			name:       "Create_Recommendation_Conflict",
			input:      req,
			wantRes:    &dto.CreateRecommendationRes{},
			wantStatus: http.StatusConflict,
			resMock:    &serviceMock{res: fmt.Errorf("recommendation with the same uuid already exist"), method: "Create"},
		},
		{
			name:       "Create_Recommendation_NotFound",
			input:      req,
			wantRes:    &dto.CreateRecommendationRes{},
			wantStatus: http.StatusNotFound,
			resMock:    &serviceMock{res: fmt.Errorf("network not found"), method: "Create"},
		},
	}
	for _, tt := range tests {
		suite.T().Run(tt.name, func(t *testing.T) {
			var got responses.CreatedSuccessResponse[dto.CreateRecommendationRes]
			if !tt.skipMock {
				suite.svcMock.On(tt.resMock.method, mock.Anything, tt.input).Return(tt.wantRes, tt.resMock.res).Times(1)
			}
			reqBody, _ := json.Marshal(tt.input)
			reqServer := httptest.NewRequest(http.MethodPost, "/recommendations", bytes.NewReader(reqBody))
			reqServer.Header.Set("Content-Type", "application/json")
			resp, _ := suite.fiberApp.Test(reqServer, -1)
			_ = json.NewDecoder(resp.Body).Decode(&got)
			//t.log(got)
			assert.Equal(suite.T(), tt.wantStatus, resp.StatusCode)
			assert.Equal(suite.T(), tt.wantRes, got.Data)
		})
	}
}
