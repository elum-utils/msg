package msg

import "github.com/labstack/echo/v4"

// Predefined serialized error message for "Too Many Requests".
var errorManyRequest = serialize(errorData{
	Key:     "MANY_REQUEST",    // Identifier for the specific error type.
	Message: "To many request", // Descriptive error message.
})

// ManyRequest is a handler that responds with a predefined JSON error message
// when the client makes too many requests.
func ManyRequest(ctx echo.Context) error {
	return ctx.JSONBlob(200, errorManyRequest) // Return the serialized error message as a JSON response.
}
