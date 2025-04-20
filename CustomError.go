package msg

import "github.com/labstack/echo/v4"

// CustomError is a handler function to send a custom error response in JSON format.
// It accepts a context for the HTTP request, an error key, and an optional message.
func CustomError(ctx echo.Context, key string, message ...string) error {

	// Initialize error message as an empty string.
	errorMessage := ""

	// If a message is provided, use the first one in the variadic arguments.
	if len(message) > 0 {
		errorMessage = message[0]
	}

	// Serialize the error data and return it as a JSON response using JSONBlob.
	return ctx.JSONBlob(200, serialize(errorData{
		Key:     key,          // Unique error identifier.
		Message: errorMessage, // Custom error message.
	}))
}
