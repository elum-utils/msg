package msg

// BadRequest sends a standardized 400 Bad Request error response.
// It provides a consistent way to handle validation and malformed request errors.
//
// Parameters:
//   - ctx: Request context compatible with multiple frameworks (Echo, Gin, Fiber, etc.)
//   - message: Human-readable error description detailing what was invalid
//   - format: Optional response format specification:
//     - "json" for JSON format (default)
//     - "msgpack" for binary MessagePack format
//     - Other custom formats if supported by the sender
//
// Returns:
//   - error: Returns an error if serialization or response sending fails
//
// Behavior:
// - Sets HTTP status code implicitly through CustomError
// - Uses standardized error key "BAD_REQUEST"
// - Formats response according to requested format
// - Maintains consistent error response structure:
//   {
//     "error": {
//       "key": "BAD_REQUEST",
//       "message": "[provided message]"
//     }
//   }
//
// Example:
//   err := msg.BadRequest(c, "Invalid email format", "json")
//   if err != nil {
//       log.Printf("Failed to send error: %v", err)
//   }
func BadRequest(ctx any, message string, format ...string) error {
    return CustomError(ctx, "BAD_REQUEST", message, format...)
}