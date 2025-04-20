package msg

import "github.com/labstack/echo"

// Predefined serialized error message for token expiration.
var errorExpiration = serialize(errorData{
    Key:     "EXPIRATION",        // Identifier for the error type related to expiration.
    Message: "Token expiration",  // Descriptive message indicating the token has expired.
})

// Expiration is a handler function that sends an error message indicating token expiration.
// It uses a pre-serialized JSON error response.
func Expiration(ctx echo.Context) error {
    // Return the serialized token expiration error message as a JSON response.
    return ctx.JSONBlob(200, errorExpiration)
}