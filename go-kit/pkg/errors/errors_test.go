package errors

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// mockError is a custom error type for testing purposes.
type mockError struct {
	message string
}

func (e *mockError) Error() string {
	return e.message
}

// TestCustomErrorHandler tests the CustomErrorHandler function.
func TestCustomErrorHandler(t *testing.T) {
	// Initialize Fiber app with the custom error handler.
	app := fiber.New(fiber.Config{
		ErrorHandler: CustomErrorHandler,
	})

	// Test route that triggers a fiber error.
	app.Get("/fiber-error", func(c *fiber.Ctx) error {
		return fiber.ErrBadRequest
	})

	// Test route that triggers a custom error.
	app.Get("/custom-error", func(c *fiber.Ctx) error {
		//return &mockError{message: "custom error occurred"}
		return fiber.ErrInternalServerError
	})

	tests := []struct {
		route       string
		expectedMsg string
		statusCode  int
	}{
		{"/fiber-error", fiber.ErrBadRequest.Message, fiber.ErrBadRequest.Code},
		{"/custom-error", "custom error occurred", fiber.StatusInternalServerError}, // Assuming your custom error defaults to 500
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.route, nil)
		resp, err := app.Test(req)
		if err != nil {
			t.Errorf("Failed to send request to %s: %v", tt.route, err)
			continue
		}

		if resp.StatusCode != tt.statusCode {
			t.Errorf("Expected status code %d for route %s, got %d", tt.statusCode, tt.route, resp.StatusCode)
		}

		// Additional response body checks can be added here if necessary.
		// For example, you might want to decode the JSON response and verify the message or other fields.
	}
}
