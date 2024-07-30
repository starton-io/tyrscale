package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/starton-io/tyrscale/manager/api/proxy/plugin/dto"
	"github.com/starton-io/tyrscale/manager/api/proxy/plugin/service"
)

type PluginHandler struct {
	pluginSvc *service.PluginService
	validator validation.Validation
}

func NewPluginHandler(pluginSvc *service.PluginService, validator validation.Validation) *PluginHandler {
	return &PluginHandler{pluginSvc: pluginSvc, validator: validator}
}

// ListPlugins godoc
//
//	@Id				listPlugins
//	@Summary		Get list plugins
//	@Description	Get list plugins
//	@Tags			plugins
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.DefaultSuccessResponse[plugin.ListPluginsResponse]
//	@Failure		400	{object}	responses.BadRequestResponse
//	@Failure		500	{object}	responses.InternalServerErrorResponse
//	@Router			/plugins [get]
func (h *PluginHandler) ListPlugins(c *fiber.Ctx) error {
	logger.Info("Fetch all plugins request")
	plugins, err := h.pluginSvc.List(c.UserContext())
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessResp.ToGeneral(plugins)
	return resp.JSON(c)
}

// ListPluginsFromRoute godoc
//
//	@Id				listPluginsFromRoute
//	@Summary		Get list plugins from route
//	@Description	Get list plugins from route
//	@Tags			plugins
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	true	"Route UUID"
//	@Success		200		{object}	responses.DefaultSuccessResponse[dto.Plugins]
//	@Failure		400		{object}	responses.BadRequestResponse
//	@Failure		500		{object}	responses.InternalServerErrorResponse
//	@Router			/routes/{uuid}/plugins [get]
func (h *PluginHandler) ListPluginsFromRoute(c *fiber.Ctx) error {
	logger.Info("Fetch all plugins request")
	if c.Params("uuid") == "" {
		logger.Error("Route uuid is required")
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("route uuid is required")).JSON(c)
	}
	logger.Infof("Route uuid: %s", c.Params("uuid"))
	plugins, err := h.pluginSvc.ListFromRoute(c.UserContext(), c.Params("uuid"))
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessResp.ToGeneral(plugins)
	return resp.JSON(c)
}

// AttachPlugin godoc
//
//	@Id				attachPlugin
//	@Summary		Attach plugin to route
//	@Description	Attach plugin to route
//	@Tags			plugins
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string				true	"Route UUID"
//	@Param			body	body		dto.AttachPluginReq	true	"Attach plugin request"
//	@Success		200		{object}	responses.DefaultSuccessResponseWithoutData
//	@Failure		400		{object}	responses.BadRequestResponse
//	@Failure		500		{object}	responses.InternalServerErrorResponse
//	@Router			/routes/{uuid}/attach-plugin [post]
func (h *PluginHandler) AttachPlugin(c *fiber.Ctx) error {
	logger.Info("Attach plugin request")
	req := &dto.AttachPluginReq{}
	if c.Params("uuid") == "" {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("route uuid is required")).JSON(c)
	}
	if err := c.BodyParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("failed to parse request")).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	err := h.pluginSvc.AttachPlugin(c.UserContext(), c.Params("uuid"), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessRespWithoutData.ToGeneral()
	return resp.JSON(c)
}

// DetachPlugin godoc
//
//	@Id				detachPlugin
//	@Summary		Detach plugin from route
//	@Description	Detach plugin from route
//	@Tags			plugins
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string				true	"Route UUID"
//	@Param			body	body		dto.DetachPluginReq	true	"Detach plugin request"
//	@Success		200		{object}	responses.DefaultSuccessResponseWithoutData
//	@Failure		400		{object}	responses.BadRequestResponse
//	@Failure		500		{object}	responses.InternalServerErrorResponse
//	@Router			/routes/{uuid}/detach-plugin [post]
func (h *PluginHandler) DetachPlugin(c *fiber.Ctx) error {
	logger.Info("Detach plugin request")
	req := &dto.DetachPluginReq{}
	if c.Params("uuid") == "" {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("route uuid is required")).JSON(c)
	}
	if err := c.BodyParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("failed to parse request")).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	err := h.pluginSvc.DetachPlugin(c.UserContext(), c.Params("uuid"), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessRespWithoutData.ToGeneral()
	return resp.JSON(c)
}
