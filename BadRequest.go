package msg

import "github.com/labstack/echo"

// BadRequest is a handler function designed to send a bad request error response.
// It takes an HTTP request context and a custom error message.
func BadRequest(ctx echo.Context, message string) error {
    // Invoke the CustomError function with a predefined error key "BAD_REQUEST"
    // and the provided message to generate and send a JSON error response.
    return CustomError(ctx, "BAD_REQUEST", message)
}