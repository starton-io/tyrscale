package http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/logger"
	"github.com/starton-io/tyrscale/go-kit/pkg/ptr"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
	"github.com/starton-io/tyrscale/go-kit/pkg/utils"
	"github.com/starton-io/tyrscale/go-kit/pkg/validation"
	"github.com/starton-io/tyrscale/manager/api/rpc/dto"
	"github.com/starton-io/tyrscale/manager/api/rpc/service"
)

type RPCHandler struct {
	validator validation.Validation
	service   service.IRPCService
}

func NewRPCHandler(service service.IRPCService, validator validation.Validation) *RPCHandler {
	return &RPCHandler{service: service, validator: validator}
}

// createRPC godoc
//
//	@Id				createRpc
//	@Summary		Create a new Rpc
//	@Description	Create a new Rpc
//	@Tags			rpcs
//	@Accept			json
//	@Produce		json
//	@Param			rpc	body		dto.CreateRpcReq	true	"Rpc request"
//	@Success		201	{object}	responses.CreatedSuccessResponse[dto.CreateRpcRes]
//	@Failure		400	{object}	responses.BadRequestResponse
//	@Failure		409	{object}	responses.ConflictResponse[dto.CreateRpcCtx]
//	@Failure		500	{object}	responses.InternalServerErrorResponse
//	@Router			/rpcs [post]
func (h *RPCHandler) CreateRPC(c *fiber.Ctx) error {
	req := new(dto.CreateRpcReq)
	if err := c.BodyParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	res, ctxErr, err := h.service.Create(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err, responses.WithContext(ctxErr))
	}
	resp := responses.CreatedSuccessResp.ToGeneral(res)
	return resp.JSON(c)
}

// listRPCs godoc
//
//	@Id				listRPCs
//	@Summary		List RPCs
//	@Description	List RPCs
//	@Tags			rpcs
//	@Accept			json
//	@Produce		json
//	@Param			uuid			query		string	false	"UUID"
//	@Param			chain_id		query		string	false	"Chain ID"
//	@Param			provider		query		string	false	"provider"
//	@Param			type			query		string	false	"type"
//	@Param			network_name	query		string	false	"network_name"
//	@Param			sort_by			query		string	false	"sort_by"
//	@Param			sort_ascending	query		bool	false	"sort_ascending"
//
//	@Success		200				{object}	responses.DefaultSuccessResponse[dto.ListRpcRes]
//	@Router			/rpcs [get]
func (h *RPCHandler) ListRPCs(c *fiber.Ctx) error {
	req := new(dto.ListReq)
	if err := c.QueryParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	// validate struct
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}

	rpcs, err := h.service.List(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	var rpcResp dto.ListRpcRes
	utils.Copy(&rpcResp.RPCs, &rpcs)
	resp := responses.DefaultSuccessResp.ToGeneral(rpcResp)
	return resp.JSON(c)
}

// deleteRPC godoc
//
//	@Id				deleteRPC
//	@Summary		Delete a RPC
//	@Description	Delete a RPC
//	@Tags			rpcs
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	true	"UUID"
//	@Success		200		{object}	responses.DefaultSuccessResponseWithoutData
//	@Failure		404		{object}	responses.NotFoundResponse
//	@Failure		500		{object}	responses.InternalServerErrorResponse
//	@Router			/rpcs/{uuid} [delete]
func (h *RPCHandler) DeleteRPC(c *fiber.Ctx) error {
	req := new(dto.DeleteRpcOptReq)
	uuid := c.Params("uuid")
	if uuid == "" {
		logger.Warn("Failed to get uuid from path")
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, fmt.Errorf("failed to get uuid from path")).JSON(c)
	}
	if c.Get("Content-Type") != "" {
		if err := c.BodyParser(req); err != nil {
			logger.Warn("Failed to parse request body", "error", err)
			resp := responses.BadRequestResp.ToGeneral()
			return resp.WithError(c, err).JSON(c)
		}
	}
	if req.CascadeDeleteUpstream == nil {
		req.CascadeDeleteUpstream = ptr.Bool(false)
	}
	deleteReq := &dto.DeleteRpcReq{
		UUID:                  uuid,
		CascadeDeleteUpstream: *req.CascadeDeleteUpstream,
	}
	fmt.Println(deleteReq)
	err := h.service.Delete(c.UserContext(), deleteReq)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessRespWithoutData.ToGeneral()
	return resp.JSON(c)
}

// updateRpc godoc
//
//	@Id				updateRPC
//	@Summary		Update a RPC
//	@Description	Update a RPC
//	@Tags			rpcs
//	@Accept			json
//	@Produce		json
//	@Param			rpc	body		dto.Rpc	true	"RPC Object request"
//	@Success		200	{object}	responses.DefaultSuccessResponseWithoutData
//	@Failure		400	{object}	responses.BadRequestResponse
//	@Failure		500	{object}	responses.InternalServerErrorResponse
//	@Router			/rpcs [put]
func (h *RPCHandler) UpdateRPC(c *fiber.Ctx) error {
	req := new(dto.UpdateRpcReq)
	if err := c.BodyParser(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	if err := h.validator.ValidateStruct(req); err != nil {
		resp := responses.BadRequestResp.ToGeneral()
		return resp.WithError(c, err).JSON(c)
	}
	err := h.service.Update(c.UserContext(), req)
	if err != nil {
		return responses.HandleServiceError(c, err)
	}
	resp := responses.DefaultSuccessRespWithoutData.ToGeneral()
	return resp.JSON(c)
}
