package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/starton-io/tyrscale/manager/api/proxy/upstream/dto"
	"github.com/starton-io/tyrscale/manager/api/proxy/upstream/service"
)

type UpstreamHandler struct {
	validator validation.Validation
	service   service.IUpstreamService
}

func NewUpstreamHandler(service service.IUpstreamService, validator validation.Validation) *UpstreamHandler {
	return &UpstreamHandler{service: service, validator: validator}
}

// UpsertUpstream godoc
//
//	@Id				upsertUpstream
//	@Summary		Create or update a upstream
//	@Description	Create or update a upstream
//	@Tags			upstreams
//	@Accept			json
//	@Produce		json
//
//	@Param			upstream	body		dto.Upstream	true	"Upstream request"
//	@Success		201			{object}	responses.CreatedSuccessResponse[dto.UpstreamUpsertRes]
//	@Failure		400			{object}	responses.BadRequestResponse			"Bad Request"
//	@Failure		500			{object}	responses.InternalServerErrorResponse	"Internal Server Error"
//	@Router			/routes/{route_uuid}/upstreams [put]
//
//	@Param			route_uuid	path	string	true	"Route UUID"
func (h *UpstreamHandler) UpsertUpstream(c *fiber.Ctx) error {
	req := new(dto.Upstream)
	routeUuid := c.Params("route_uuid")
	if routeUuid == "" {
		logger.Error("Failed to get route_uuid from path")
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("failed to get route_uuid from path")).JSON(c)
	}
	if err := c.BodyParser(&req); c.Body() == nil || err != nil {
		logger.Warnf("Failed to parse body: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		logger.Warnf("Validation failed for create upstream: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	upstreamResp, err := h.service.Upsert(c.UserContext(), routeUuid, req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.CreatedSuccessResp.ToGeneral(*upstreamResp)
	return resp.JSON(c)
}

// ListUpstreams godoc
//
//	@Id				listUpstreams
//	@Summary		Get list upstreams
//	@Description	Get list upstreams
//	@Tags			upstreams
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.DefaultSuccessResponse[dto.ListUpstreamRes]
//	@Failure		400	{object}	responses.BadRequestResponse
//	@Failure		500	{object}	responses.InternalServerErrorResponse
//	@Router			/routes/{route_uuid}/upstreams [get]
//	@Param			route_uuid	path	string	true	"Route UUID"
func (h *UpstreamHandler) ListUpstreams(c *fiber.Ctx) error {
	routeUuid := c.Params("route_uuid")
	if routeUuid == "" {
		logger.Error("Failed to get route_uuid from path")
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("failed to get route_uuid from path")).JSON(c)
	}
	req := new(dto.ListUpstreamReq)
	if err := c.QueryParser(req); err != nil {
		logger.Warnf("Failed to parse query: %v", err)
		//return responses.DefaultBadRequestResponse.WithError(c, err).JSON(c)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	req.RouteUuid = routeUuid
	if err := h.validator.ValidateStruct(req); err != nil {
		logger.Warnf("Validation failed for list upstreams: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	logger.Info("Fetch all upstreams request", "query", req)
	upstreamResp, err := h.service.List(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	logger.Info("Fetch all upstreams response", "body", upstreamResp)
	var res dto.ListUpstreamRes
	utils.Copy(&res.Upstreams, &upstreamResp)
	logger.Info("Fetch all upstreams response", "body", res)
	resp := responses.DefaultSuccessResp.ToGeneral(res)
	return resp.JSON(c)
}

func (h *UpstreamHandler) getParam(c *fiber.Ctx, param string) (string, error) {
	value := c.Params(param)
	if value == "" {
		logger.Errorf("Failed to get %s from path", param)
		resp := responses.BadRequestResp.ToGeneral()
		return "", resp.WithError(c, fmt.Errorf("failed to get %s from path", param)).JSON(c)
	}
	return value, nil
}

// DeleteUpstream godoc
//
//	@Id				deleteUpstream
//	@Summary		Delete a upstream
//	@Description	Delete a upstream
//	@Tags			upstreams
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.DefaultSuccessResponse[dto.UpstreamUpsertRes]
//	@Failure		400	{object}	responses.BadRequestResponse
//	@Failure		500	{object}	responses.InternalServerErrorResponse
//	@Router			/routes/{route_uuid}/upstreams/{uuid} [delete]
//
//	@Param			route_uuid	path	string	true	"Route UUID"
//	@Param			uuid		path	string	true	"Upstream UUID"
func (h *UpstreamHandler) DeleteUpstream(c *fiber.Ctx) error {
	routeUuid, err := h.getParam(c, "route_uuid")
	if err != nil {
		return err
	}
	uuid, err := h.getParam(c, "uuid")
	if err != nil {
		return err
	}
	logger.Debugf("Delete upstream request: %v", uuid)
	req := &dto.ListUpstreamReq{RouteUuid: routeUuid, Uuid: &uuid}

	//validate request
	if err := h.validator.ValidateStruct(req); err != nil {
		logger.Warnf("Validation failed for delete upstream: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	// list upstreams
	upstreams, err := h.service.List(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	if len(upstreams) == 0 {
		return responses.HandleServiceError(c, fmt.Errorf("upstream not found"))
	}
	reqDelete := &dto.UpstreamDeleteReq{
		Uuid:      uuid,
		RouteUuid: routeUuid,
	}
	err = h.service.Delete(c.UserContext(), reqDelete)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}

	upstreamResp := &dto.UpstreamUpsertRes{Uuid: uuid}
	resp := responses.DefaultSuccessResp.ToGeneral(*upstreamResp)
	return resp.JSON(c)
}
