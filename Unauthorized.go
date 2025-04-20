package msg

import "github.com/labstack/echo/v4"

// Predefined serialized error message indicating that the user is unauthorized.
var errorUnauthorized = serialize(errorData{
	Key:     "UNAUTHORIZED", // Identifier for unauthorized access errors.
	Message: "Unauthorized", // Message indicating lack of authentication or permission.
})

// Unauthorized is a handler function that sends an error message when a user
// attempts to access a resource they are not authorized to access. It responds with a
// pre-serialized JSON message.
func Unauthorized(ctx echo.Context) error {
	// Return the serialized unauthorized error message as a JSON response.
	return ctx.JSONBlob(200, errorUnauthorized)
}
