package errors

import (
	"github.com/starton-io/tyrscale/go-kit/pkg/httpcode"
)

type Error struct {
	Status  int    `json:"-"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) Error() string {
	return err.Message
}

var (
	ErrInternalServer = &Error{
		Status:  500,
		Code:    httpcode.INTERNAL,
		Message: "Internal server error",
	}

	ErrBadRequest = &Error{
		Status:  400,
		Code:    httpcode.INVALID_ARGUMENT,
		Message: "Bad request",
	}

	ErrPermissionDenied = &Error{
		Status:  403,
		Code:    httpcode.PERMISSION_DENIED,
		Message: "Permission denied",
	}

	ErrNotFound = &Error{
		Status:  404,
		Code:    httpcode.NOT_FOUND,
		Message: "Not found",
	}

	ErrAlreadyExists = &Error{
		Status:  409,
		Code:    httpcode.ALREADY_EXISTS,
		Message: "Already exists",
	}

	ErrUnauthenticated = &Error{
		Status:  401,
		Code:    httpcode.UNAUTHENTICATED,
		Message: "Unauthorized",
	}
)
