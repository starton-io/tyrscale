package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/starton-io/tyrscale/manager/api/recommendation/dto"
	"github.com/starton-io/tyrscale/manager/api/recommendation/service"
)

type RecommendationHandler struct {
	recommendationService service.IRecommendationService
	validator             validation.Validation
}

func NewHandler(recommendationService service.IRecommendationService, validator validation.Validation) *RecommendationHandler {
	return &RecommendationHandler{
		recommendationService: recommendationService,
		validator:             validator,
	}
}

// createRecommendation godoc
//
//	@Id				createRecommendation
//
//	@Summary		Create a recommendation
//	@Description	Create a recommendation
//	@Tags			recommendation
//	@Accept			json
//	@Produce		json
//
//	@Param			recommendation	body		dto.CreateRecommendationReq	true	"Recommendation Object request"
//
//	@Success		201				{object}	responses.CreatedSuccessResponse[dto.CreateRecommendationRes]
//	@Failure		400				{object}	responses.BadRequestResponse			"Bad Request"
//	@Failure		500				{object}	responses.InternalServerErrorResponse	"Internal Server Error"
//	@Router			/recommendations [post]
func (h *RecommendationHandler) CreateRecommendation(c *fiber.Ctx) error {
	req := new(dto.CreateRecommendationReq)
	if err := c.BodyParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	recommendationResp, err := h.recommendationService.Create(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.CreatedSuccessResp.ToGeneral(recommendationResp)
	return resp.JSON(c)
}

// ListRecommendation godoc
//
//	@Id				listRecommendations
//	@Summary		List recommendation
//	@Description	List recommendation
//	@Tags			recommendation
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.DefaultSuccessResponse[dto.ListRecommendationRes]
//	@Failure		400	{object}	responses.BadRequestResponse			"Bad Request"
//	@Failure		500	{object}	responses.InternalServerErrorResponse	"Internal Server Error"
//	@Router			/recommendations [get]
func (h *RecommendationHandler) ListRecommendation(c *fiber.Ctx) error {
	req := new(dto.ListRecommendationReq)
	if err := c.QueryParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	recommendations, err := h.recommendationService.List(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	var recommendationResp dto.ListRecommendationRes
	utils.Copy(&recommendationResp.Recommendations, recommendations)
	resp := responses.DefaultSuccessResp.ToGeneral(recommendationResp)
	return resp.JSON(c)
}

// DeleteRecommendation godoc
//
//	@Id				deleteRecommendation
//	@Summary		Delete a recommendation
//	@Description	Delete a recommendation
//	@Tags			recommendation
//	@Accept			json
//	@Produce		json
//	@Param			route_uuid	path		string	true	"Route UUID"
//	@Success		200			{object}	responses.DefaultSuccessResponseWithoutData
//	@Failure		400			{object}	responses.BadRequestResponse			"Bad Request"
//	@Failure		404			{object}	responses.NotFoundResponse				"Not Found"
//	@Failure		409			{object}	responses.ConflictResponse				"Conflict"
//	@Failure		500			{object}	responses.InternalServerErrorResponse	"Internal Server Error"
//	@Router			/recommendations/{route_uuid} [delete]
func (h *RecommendationHandler) DeleteRecommendation(c *fiber.Ctx) error {
	req := new(dto.DeleteRecommendationReq)
	req.RouteUuid = c.Params("route_uuid")
	if err := c.QueryParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	err := h.recommendationService.Delete(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessRespWithoutData.ToGeneral()
	return resp.JSON(c)
}

// UpdateRecommendation godoc
//
//	@Id				updateRecommendation
//	@Summary		Update a recommendation
//	@Description	Update a recommendation
//	@Tags			recommendation
//	@Accept			json
//	@Produce		json
//	@Param			recommendation	body		dto.UpdateRecommendationReq	true	"Recommendation Object request"
//	@Success		200				{object}	responses.DefaultSuccessResponseWithoutData
//	@Failure		400				{object}	responses.BadRequestResponse			"Bad Request"
//	@Failure		404				{object}	responses.NotFoundResponse				"Not Found"
//	@Failure		500				{object}	responses.InternalServerErrorResponse	"Internal Server Error"
//	@Router			/recommendations [put]
func (h *RecommendationHandler) UpdateRecommendation(c *fiber.Ctx) error {
	req := new(dto.UpdateRecommendationReq)
	if err := c.BodyParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	err := h.recommendationService.Update(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessRespWithoutData.ToGeneral()
	return resp.JSON(c)
}
