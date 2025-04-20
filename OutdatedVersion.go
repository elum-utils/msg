package msg

import "github.com/labstack/echo/v4"

// Predefined serialized error message for an outdated application version.
var errorOutdatedVersion = serialize(errorData{
	Key:     "OUTDATED_VERSION",                    // Identifier for the error related to outdated application version.
	Message: "Outdated version of the application", // Message indicating that the application version is outdated.
})

// OutdatedVersion is a handler function that sends an error message when the application version
// is outdated. It utilizes a pre-serialized JSON error response.
func OutdatedVersion(ctx echo.Context) error {
	// Return the serialized outdated version error message as a JSON response.
	return ctx.JSONBlob(200, errorOutdatedVersion)
}
