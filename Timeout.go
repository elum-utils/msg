package msg

import "github.com/labstack/echo"

// Predefined serialized error message indicating that the request has timed out.
var errorTimeout = serialize(errorData{
    Key:     "TIMEOUT",                // Identifier for a timeout error.
    Message: "The request has timed out", // Message indicating that the request exceeded the time limit.
})

// Timeout is a handler function that sends an error message indicating
// that the request has exceeded the allowed time limit. It returns a pre-serialized JSON response.
func Timeout(ctx echo.Context) error {
    // Return the serialized timeout error message as a JSON response.
    return ctx.JSONBlob(200, errorTimeout)
}