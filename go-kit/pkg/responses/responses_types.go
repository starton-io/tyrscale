package responses

import (
	"net/http"

	"github.com/starton-io/tyrscale/go-kit/pkg/httpcode"
)

var (
	CreatedSuccessResp = CreatedSuccessResponse[any]{
		Status:  http.StatusCreated,
		Code:    httpcode.OK,
		Message: "Resource Created Successfully",
	}
	DefaultSuccessResp = DefaultSuccessResponse[any]{
		Status:  http.StatusOK,
		Code:    httpcode.OK,
		Message: "Success",
	}
	DefaultSuccessRespWithoutData = DefaultSuccessResponseWithoutData[any]{
		Status:  http.StatusOK,
		Code:    httpcode.OK,
		Message: "Success",
	}
	InternalServerErrorResp = InternalServerErrorResponse[any]{
		Status:  http.StatusInternalServerError,
		Code:    httpcode.INTERNAL,
		Message: "Internal server error",
	}
	BadRequestResp = BadRequestResponse[any]{
		Status:  http.StatusBadRequest,
		Code:    httpcode.INVALID_ARGUMENT,
		Message: "Invalid input parameters",
	}
	ConflictResp = ConflictResponse[any]{
		Status:  http.StatusConflict,
		Code:    httpcode.ALREADY_EXISTS,
		Message: "Resource Conflict",
	}
	NotFoundResp = NotFoundResponse[any]{
		Status:  http.StatusNotFound,
		Code:    httpcode.NOT_FOUND,
		Message: "Resource Not found",
	}
)

type CreatedSuccessResponse[T any] struct {
	Status  int    `json:"status" example:"201"`
	Code    int    `json:"code" example:"0"`
	Message string `json:"message" example:"Created Success"`
	Data    T      `json:"data,omitempty"`
}

func (resp *CreatedSuccessResponse[T]) ToGeneral(data T) General[T] {
	return General[T]{
		Status:  resp.Status,
		Code:    resp.Code,
		Message: resp.Message,
		Data:    data,
	}
}

type DefaultSuccessResponse[T any] struct {
	Status  int    `json:"status" example:"200"`
	Code    int    `json:"code" example:"0"`
	Message string `json:"message" example:"Created Success"`
	Data    T      `json:"data,omitempty"`
}

type DefaultSuccessResponseWithoutData[T any] struct {
	Status  int    `json:"status" example:"200"`
	Code    int    `json:"code" example:"0"`
	Message string `json:"message" example:"Success"`
}

func (resp *DefaultSuccessResponseWithoutData[T]) ToGeneral() General[T] {
	return General[T]{
		Status:  resp.Status,
		Code:    resp.Code,
		Message: resp.Message,
	}
}

func (resp *DefaultSuccessResponse[T]) ToGeneral(data T) General[T] {
	return General[T]{
		Status:  resp.Status,
		Code:    resp.Code,
		Message: resp.Message,
		Data:    data,
	}
}

type BadRequestResponse[T any] struct {
	Status  int    `json:"status" example:"400"`
	Code    int    `json:"code" example:"3"`
	Message string `json:"message" example:"Invalid input parameters"`
}

func (resp *BadRequestResponse[T]) ToGeneral() General[T] {
	return General[T]{
		Status:  resp.Status,
		Code:    resp.Code,
		Message: resp.Message,
	}
}

type ConflictResponse[T any] struct {
	Status  int    `json:"status" example:"409"`
	Code    int    `json:"code" example:"5"`
	Message string `json:"message" example:"Conflict"`
	Context T      `json:"context,omitempty"`
}

type ConflictResponseWithoutContext struct {
	Status  int    `json:"status" example:"409"`
	Code    int    `json:"code" example:"5"`
	Message string `json:"message" example:"Conflict"`
}

func (resp *ConflictResponse[T]) ToGeneral() General[T] {
	return General[T]{
		Status:  resp.Status,
		Code:    resp.Code,
		Message: resp.Message,
	}
}

func (resp *ConflictResponse[T]) ToGeneralWithContext(context T) General[T] {
	return General[T]{
		Status:  resp.Status,
		Code:    resp.Code,
		Message: resp.Message,
		Context: context,
	}
}

type InternalServerErrorResponse[T any] struct {
	Status  int    `json:"status" example:"500"`
	Code    int    `json:"code" example:"13"`
	Message string `json:"message" example:"Internal server error"`
}

func (resp *InternalServerErrorResponse[T]) ToGeneral() General[T] {
	return General[T]{
		Status:  resp.Status,
		Code:    resp.Code,
		Message: resp.Message,
	}
}

type NotFoundResponse[T any] struct {
	Status  int    `json:"status" example:"404"`
	Code    int    `json:"code" example:"6"`
	Message string `json:"message" example:"Not found"`
}

func (resp *NotFoundResponse[T]) ToGeneral() General[T] {
	return General[T]{
		Status:  resp.Status,
		Code:    resp.Code,
		Message: resp.Message,
	}
}
