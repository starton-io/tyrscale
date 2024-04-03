package errors

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/responses"
)

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var msg responses.General[string]

	// trieve the custom status code if it's an fiber.*Error
	var e *fiber.Error
	if errors.As(err, &e) {
		msg.Status = e.Code
		msg.Code = e.Code
		msg.Message = e.Message
	}
	var customErr *Error
	if errors.As(err, &customErr) {
		msg = responses.BindingGeneral[string](customErr.Error())
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return msg.JSON(ctx)
}
