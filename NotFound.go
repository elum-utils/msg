package msg

import "github.com/labstack/echo/v4"

// Predefined serialized error message indicating that the requested resource could not be found.
var errorNotFound = serialize(errorData{
	Key:     "NOT_FOUND",                                 // Identifier for a "not found" error.
	Message: "The requested resource could not be found", // Message indicating the resource is unavailable.
})

// NotFound is a handler function that sends an error message indicating
// that the requested resource does not exist. It returns a pre-serialized JSON response.
func NotFound(ctx echo.Context) error {
	// Return the serialized "not found" error message as a JSON response.
	return ctx.JSONBlob(200, errorNotFound)
}
