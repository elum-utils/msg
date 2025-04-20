package msg

import "github.com/labstack/echo/v4"

// Predefined serialized error message indicating forbidden access.
var errorForbidden = serialize(errorData{
	Key:     "FORBIDDEN", // Identifier for the forbidden access error.
	Message: "Forbidden", // Descriptive message for forbidden action.
})

// Forbidden is a handler function that sends an error message when access is forbidden.
// It returns a pre-serialized JSON error response.
func Forbidden(ctx echo.Context) error {
	// Return the serialized forbidden error message as a JSON response.
	return ctx.JSONBlob(200, errorForbidden)
}
