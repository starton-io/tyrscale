package responses

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/starton-io/tyrscale/go-kit/pkg/httpcode"
)

// General is now a generic struct that can hold any type of data in the Data field.
type General[T any] struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"` // Data is now of generic type T
	Context T      `json:"context,omitempty"`
}

type Body[T any] struct {
	Items []T `json:"items"`
}

// JSON method remains the same but now works with the generic type.
func (g *General[T]) JSON(c *fiber.Ctx) error {
	return c.Status(g.Status).JSON(g)
}

// WithError method updates the message and returns the modified General struct.
func (g *General[T]) WithError(c *fiber.Ctx, err error) *General[T] {
	g.Message = g.Message + ": " + err.Error()
	return g
}

func (g *General[T]) WithMessage(message string) *General[T] {
	g.Message = message
	return g
}

// WithData method allows setting the Data field with a value of the generic type.
func (g *General[T]) WithData(data T) *General[T] {
	g.Data = data
	return g
}

// // Assuming this function is somewhere in your package
func DefaultErrorResponse[T any]() General[T] {
	return General[T]{
		Status:  http.StatusInternalServerError,
		Code:    httpcode.INTERNAL,
		Message: "An internal server error occurred",
	}
}

// BindingGeneral function creates a General struct from any data type using generics.
func BindingGeneral[T any](data T) General[T] {
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Assuming DefaultErrorResponse is of type General[T] and properly handles error cases
		return DefaultErrorResponse[T]()
	}
	var response General[T]
	if err := json.Unmarshal(jsonData, &response); err != nil {
		// Assuming DefaultErrorResponse is of type General[T] and properly handles error cases
		return DefaultErrorResponse[T]()
	}
	return response
}

type Option func(*General[any])

func WithContext(context any) Option {
	return func(g *General[any]) {
		if context != nil && !isNil(context) {
			g.Context = context
		}
	}
}

// Helper function to check if an interface is nil
func isNil(i interface{}) bool {
	return i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil())
}

func HandleServiceError(c *fiber.Ctx, err error, opts ...Option) error {
	var resp General[any]
	switch {
	case strings.Contains(err.Error(), "already") || strings.Contains(err.Error(), "associated") || strings.Contains(err.Error(), "conflict"):
		resp = ConflictResp.ToGeneral()
	case strings.Contains(err.Error(), "not found"):
		resp = NotFoundResp.ToGeneral()
	case strings.Contains(err.Error(), "invalid"):
		resp = BadRequestResp.ToGeneral()
	default:
		resp = InternalServerErrorResp.ToGeneral()
	}

	// Apply options
	for _, opt := range opts {
		opt(&resp)
	}

	return resp.WithError(c, err).JSON(c)
}
