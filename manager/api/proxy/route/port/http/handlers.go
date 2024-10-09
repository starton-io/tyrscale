package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/gateway/pkg/plugin"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	pluginDto "github.com/starton-io/tyrscale/manager/api/proxy/plugin/dto"
	pluginService "github.com/starton-io/tyrscale/manager/api/proxy/plugin/service"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/dto"
	"github.com/starton-io/tyrscale/manager/api/proxy/route/service"
	upstreamDto "github.com/starton-io/tyrscale/manager/api/proxy/upstream/dto"
	upstreamService "github.com/starton-io/tyrscale/manager/api/proxy/upstream/service"
)

type RouteHandler struct {
	validator       validation.Validation
	service         service.IRouteService
	upstreamService upstreamService.IUpstreamService
	pluginService   pluginService.IPluginService
}

func NewRouteHandler(service service.IRouteService, upstreamService upstreamService.IUpstreamService, pluginService pluginService.IPluginService, validator validation.Validation) *RouteHandler {
	return &RouteHandler{service: service, validator: validator, upstreamService: upstreamService, pluginService: pluginService}
}

// CreateRoute godoc
//
//	@Id				createRoute
//	@Summary		Create a route
//	@Description	Create a route
//	@Tags			routes
//	@Accept			json
//	@Produce		json
//
//	@Param			route	body		dto.CreateRouteReq	true	"Route request"
//	@Success		201		{object}	responses.CreatedSuccessResponse[dto.CreateRouteRes]
//	@Failure		400		{object}	responses.BadRequestResponse			"Bad Request"
//	@Failure		500		{object}	responses.InternalServerErrorResponse	"Internal Server Error"
//	@Router			/routes [post]
func (h *RouteHandler) CreateRoute(c *fiber.Ctx) error {
	req := new(dto.CreateRouteReq)
	if err := c.BodyParser(&req); c.Body() == nil || err != nil {
		logger.Warnf("Failed to parse body: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	if err := h.validator.ValidateStruct(req); err != nil {
		logger.Warnf("Validation failed for create network: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	routeResp, err := h.service.Create(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.CreatedSuccessResp.ToGeneral(*routeResp)
	return resp.JSON(c)
}

// UpdateRoute godoc
//
//	@Id				updateRoute
//	@Summary		Update a route
//	@Description	Update a route
//	@Tags			routes
//	@Accept			json
//	@Produce		json
//
//	@Param			uuid	path		string				true	"UUID"
//	@Param			route	body		dto.UpdateRouteReq	true	"Route request"
//	@Success		200		{object}	responses.DefaultSuccessResponseWithoutData
//	@Failure		400		{object}	responses.BadRequestResponse			"Bad Request"
//	@Failure		500		{object}	responses.InternalServerErrorResponse	"Internal Server Error"
//	@Router			/routes/{uuid} [patch]
func (h *RouteHandler) UpdateRoute(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		logger.Error("Failed to get uuid from path")
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("failed to get uuid from path")).JSON(c)
	}

	req := new(dto.UpdateRouteReq)
	if err := c.BodyParser(&req); c.Body() == nil || err != nil {
		logger.Warnf("Failed to parse body: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	if err := h.validator.ValidateStruct(req); err != nil {
		logger.Warnf("Validation failed for create network: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	err := h.service.Update(c.UserContext(), uuid, req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessRespWithoutData.ToGeneral()
	return resp.JSON(c)
}

// ListRoutes godoc
//
//	@Id				listRoutes
//	@Summary		Get list routes
//	@Description	Get list routes
//	@Tags			routes
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.DefaultSuccessResponse[dto.ListRouteRes]
//	@Failure		400	{object}	responses.BadRequestResponse
//	@Failure		500	{object}	responses.InternalServerErrorResponse
//	@Router			/routes [get]
//	@Param			_	query	dto.ListRouteReq	false	"comment"
func (h *RouteHandler) ListRoutes(c *fiber.Ctx) error {
	var req dto.ListRouteReq
	if err := c.QueryParser(&req); err != nil {
		logger.Warnf("Failed to parse query: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(&req); err != nil {
		logger.Warnf("Validation failed for list routes: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	logger.Info("Fetch all routes request", "query", req)
	routeResp, err := h.service.List(c.UserContext(), &req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	var res dto.ListRouteRes
	utils.Copy(&res.Routes, &routeResp)
	logger.Info("Fetch all routes response", "body", res)
	resp := responses.DefaultSuccessResp.ToGeneral(res)
	return resp.JSON(c)
}

// DeleteRoute godoc
//
//	@Id				deleteRoute
//	@Summary		Delete a route
//	@Description	Delete a route
//	@Tags			routes
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	true	"Route UUID"
//	@Success		200		{object}	responses.DefaultSuccessResponseWithoutData
//	@Failure		400		{object}	responses.BadRequestResponse
//	@Failure		500		{object}	responses.InternalServerErrorResponse
//	@Router			/routes/{uuid} [delete]
func (h *RouteHandler) DeleteRoute(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		logger.Error("Failed to get uuid from path")
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("failed to get route uuid from path")).JSON(c)
	}
	logger.Debugf("Delete route request: %v", uuid)

	// list upstreams
	upstreams, err := h.upstreamService.List(c.UserContext(), &upstreamDto.ListUpstreamReq{RouteUuid: uuid})
	if err != nil {
		return responses.HandleServiceError(c, err)
	}

	// delete upstreams
	for _, upstream := range upstreams {
		err = h.upstreamService.Delete(c.UserContext(), &upstreamDto.UpstreamDeleteReq{Uuid: upstream.Uuid, RouteUuid: uuid})
		if err != nil {
			return responses.HandleServiceError(c, err)
		}
	}

	// detach plugins
	plugins, err := h.pluginService.ListFromRoute(c.UserContext(), uuid)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	for _, pluginValue := range plugins.Middleware {
		err = h.pluginService.DetachPlugin(c.UserContext(), uuid, &pluginDto.DetachPluginReq{Name: pluginValue.Name, Type: plugin.PluginTypeMiddleware})
		if err != nil {
			return responses.HandleServiceError(c, err)
		}
	}
	for _, pluginValue := range plugins.InterceptorRequest {
		err = h.pluginService.DetachPlugin(c.UserContext(), uuid, &pluginDto.DetachPluginReq{Name: pluginValue.Name, Type: plugin.PluginTypeRequestInterceptor})
		if err != nil {
			return responses.HandleServiceError(c, err)
		}
	}
	for _, pluginValue := range plugins.InterceptorResponse {
		err = h.pluginService.DetachPlugin(c.UserContext(), uuid, &pluginDto.DetachPluginReq{Name: pluginValue.Name, Type: plugin.PluginTypeResponseInterceptor})
		if err != nil {
			return responses.HandleServiceError(c, err)
		}
	}

	// delete route
	err = h.service.Delete(c.UserContext(), &dto.DeleteRouteReq{Uuid: uuid})
	if err != nil {
		return responses.HandleServiceError(c, err)
	}

	resp := responses.DefaultSuccessRespWithoutData.ToGeneral()
	return resp.JSON(c)
}
