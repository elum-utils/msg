package msg

import "github.com/labstack/echo"

// Predefined serialized error message indicating that an unknown error has occurred.
var errorUnknownError = serialize(errorData{
    Key:     "UNKNOWN_ERROR",                            // Identifier for an unknown error.
    Message: "An unknown error has occurred. Please try again later.", // Message indicating an unexpected issue.
})

// UnknownError is a handler function that sends an error message indicating
// an unknown issue has occurred. It returns a pre-serialized JSON response to the client.
func UnknownError(ctx echo.Context) error {
    // Return the serialized unknown error message as a JSON response.
    return ctx.JSONBlob(200, errorUnknownError)
}