package msg

import "github.com/labstack/echo"

// Predefined serialized error message indicating that technical work is underway.
var errorServiceWork = serialize(errorData{
    Key:     "SERVICE_WORK",              // Identifier for the service work status.
    Message: "Technical work is underway", // Message indicating ongoing technical work.
})

// ServiceWork is a handler function that sends an error message indicating
// that the service is undergoing technical work. It returns a pre-serialized JSON response.
func ServiceWork(ctx echo.Context) error {
    // Return the serialized service work message as a JSON response.
    return ctx.JSONBlob(200, errorServiceWork)
}