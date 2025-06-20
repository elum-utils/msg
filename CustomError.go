package msg

// CustomError sends a custom error response in specified format (JSON or MessagePack).
// It provides a unified interface for error handling across different web frameworks.
//
// Parameters:
//   - ctx: Context object compatible with various web frameworks (Echo, Gin, Fiber etc.)
//   - key: Unique error identifier/error code (e.g., "AUTH_REQUIRED", "NOT_FOUND")
//   - message: Human-readable error description
//   - format: Optional format specification:
//     - "json" for JSON response (default)
//     - "msgpack" for MessagePack binary format
//
// Returns:
//   - error: Returns serialization or sending error if any occurs
//
// Behavior:
// - Automatically sets appropriate Content-Type header
// - Supports multiple web frameworks through unified interface
// - Provides consistent error response structure
// - Defaults to JSON format if none specified
func CustomError(ctx any, key string, message string, format ...string) (err error) {
    // Determine response content type
    contentType := "json" // Default to JSON format
    if len(format) > 0 {
        switch format[0] {
        case "msgpack":
            contentType = "msgpack"
        case "json":
            contentType = "json"
        default:
            contentType = format[0] // Allow custom formats
        }
    }

    var (
        messageJson    []byte // JSON-formatted error response
        messageMsgpack []byte // MessagePack-formatted error response
    )

    // Prepare error data structure
    errData := errorData{
        Key:     key,     // Set error code/identifier
        Message: message, // Set error description
    }

    // Serialize to JSON when:
    // - JSON explicitly requested, or
    // - No format specified (default case)
    if contentType == "json" {
        messageJson, err = serialize(errData, contentType)
        if err != nil {
            return err
        }
    }

    // Serialize to MessagePack when explicitly requested
    if contentType == "msgpack" {
        messageMsgpack, err = serialize(errData, contentType)
        if err != nil {
            return err
        }
    }

    // Send prepared response using unified sender
    return send(ctx, messageJson, messageMsgpack, format...)
}