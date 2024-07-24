package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/manager/api/proxy/plugin/service"
)

type PluginHandler struct {
	pluginSvc *service.PluginService
}

func NewPluginHandler(pluginSvc *service.PluginService) *PluginHandler {
	return &PluginHandler{pluginSvc: pluginSvc}
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
