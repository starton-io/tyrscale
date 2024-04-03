package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/starton-io/tyrscale/manager/api/network/dto"
	"github.com/starton-io/tyrscale/manager/api/network/service"
)

type NetworkHandler struct {
	validator validation.Validation
	service   service.INetworkService
}

func NewNetworkHandler(service service.INetworkService, validator validation.Validation) *NetworkHandler {
	return &NetworkHandler{service: service, validator: validator}
}

// CreateNetwork godoc
//
//	@Id				createNetwork
//	@Summary		Create a network
//	@Description	Create a network
//	@Tags			networks
//	@Accept			json
//	@Produce		json
//
//	@Param			network	body		dto.Network	true	"Network request"
//	@Success		201		{object}	responses.CreatedSuccessResponse[dto.CreateNetworkRes]
//	@Failure		400		{object}	responses.BadRequestResponse			"Bad Request"
//	@Failure		500		{object}	responses.InternalServerErrorResponse	"Internal Server Error"
//	@Router			/networks [post]
func (h *NetworkHandler) CreateNetwork(c *fiber.Ctx) error {
	req := new(dto.Network)
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

	err := h.service.Create(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	var networkResp = &dto.CreateNetworkRes{ChainId: req.ChainId}
	//resp := responses.DefaultSuccessCreatedResponse[dto.CreateNetworkRes](*networkResp)
	resp := responses.CreatedSuccessResp.ToGeneral(*networkResp)
	return resp.JSON(c)
}

// ListNetworks godoc
//
//	@Id				listNetworks
//	@Summary		Get list networks
//	@Description	Get list networks
//	@Tags			networks
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	responses.DefaultSuccessResponse[dto.ListNetworkRes]
//	@Failure		400	{object}	responses.BadRequestResponse
//	@Failure		500	{object}	responses.InternalServerErrorResponse
//	@Router			/networks [get]
//	@Param			_	query	dto.ListNetworkReq	false	"comment"
func (h *NetworkHandler) ListNetworks(c *fiber.Ctx) error {
	var req dto.ListNetworkReq
	if err := c.QueryParser(&req); err != nil {
		logger.Warnf("Failed to parse query: %v", err)
		//return responses.DefaultBadRequestResponse.WithError(c, err).JSON(c)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(&req); err != nil {
		logger.Warnf("Validation failed for list networks: %v", err)
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	logger.Info("Fetch all networks request", "query", req)
	networkResp, err := h.service.List(c.UserContext(), &req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	var res responses.Body[dto.Network]
	utils.Copy(&res.Items, &networkResp)
	logger.Info("Fetch all networks response", "body", res)
	resp := responses.DefaultSuccessResp.ToGeneral(res)
	return resp.JSON(c)
}

// DeleteNetwork godoc
//
//	@Id				deleteNetwork
//	@Summary		Delete a network
//	@Description	Delete a network
//	@Tags			networks
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"Network Name"
//	@Success		200		{object}	responses.DefaultSuccessResponse[dto.DeleteNetworkRes]
//	@Failure		400		{object}	responses.BadRequestResponse
//	@Failure		500		{object}	responses.InternalServerErrorResponse
//	@Router			/networks/{name} [delete]
func (h *NetworkHandler) DeleteNetwork(c *fiber.Ctx) error {
	name := c.Params("name")
	if name == "" {
		logger.Error("Failed to get name from path")
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("failed to get chainID from path")).JSON(c)
	}
	logger.Debugf("Delete network request: %v", name)
	err := h.service.Delete(c.UserContext(), name)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}

	networkResp := &dto.DeleteNetworkRes{Name: name}
	resp := responses.DefaultSuccessResp.ToGeneral(*networkResp)
	return resp.JSON(c)
}
