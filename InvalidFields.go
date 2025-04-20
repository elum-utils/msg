package msg

import "github.com/labstack/echo/v4"

// Predefined serialized error message for invalid fields.
var errorInvalidFields = serialize(errorData{
	Key:     "INVALID_FIELDS",                                                            // Identifier for the invalid fields error type.
	Message: "Required fields are missing or their type does not match the declared one", // Detailed message describing the error condition.
})

// InvalidFields is a handler function that sends an error message when required fields
// are missing or their types are incorrect. It uses a pre-serialized JSON error response.
func InvalidFields(ctx echo.Context) error {
	// Return the serialized invalid fields error message as a JSON response.
	return ctx.JSONBlob(200, errorInvalidFields)
}
