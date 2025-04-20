package msg

import "github.com/labstack/echo"

// Predefined serialized error message indicating that an unexpected internal server error occurred.
var errorInternalError = serialize(errorData{
	Key:     "INTERNAL_ERROR",                               // Identifier for an internal server error.
	Message: "An unexpected internal server error occurred", // Message describing the server error.
})

// InternalError is a handler function that sends an error message indicating
// an unexpected internal server issue. It returns a pre-serialized JSON response.
func InternalError(ctx echo.Context) error {
	// Return the serialized internal server error message as a JSON response.
	return ctx.JSONBlob(200, errorInternalError)
}
