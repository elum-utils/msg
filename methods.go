package msg

import (
	"errors"

	jsoniter "github.com/json-iterator/go"
	"github.com/vmihailenco/msgpack/v5"
)

// ByteSender defines the interface for sending raw byte data.
// It's used as a fallback for custom context types not natively supported.
type ByteSender interface {
	Send([]byte) error
}

// FiberContext represents the Fiber framework's request context interface.
type FiberContext interface {
	Set(string, string) // Set response header
	Status(int) any     // Set HTTP status code
	Write([]byte) error // Write raw response
}

// EchoContext represents the Echo framework's request context interface.
type EchoContext interface {
	Response() EchoResponse // Get response writer
}

// EchoResponse represents Echo's response writer interface.
type EchoResponse interface {
	Header() Header            // Get header setter
	WriteHeader(int)           // Write status code
	Write([]byte) (int, error) // Write response body
}

// GinContext represents the Gin framework's request context interface.
type GinContext interface {
	Writer() GinWriter // Get response writer
	Status(int) any    // Set HTTP status code
}

// GinWriter represents Gin's response writer interface.
type GinWriter interface {
	Header() Header            // Get header setter
	Write([]byte) (int, error) // Write response body
}

// Header provides common interface for setting HTTP headers.
type Header interface {
	Set(string, string) // Set header value
}

var (
	ErrorUnsupportedContextType = errors.New("unsupported context type")
)

// json is a high-performance JSON serializer compatible with standard library.
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// serialize converts data to specified format (JSON or MessagePack).
//
// Parameters:
//   - data: The data structure to serialize (must be serializable)
//   - format: Target format ("json" or "msgpack"), defaults to JSON
//
// Returns:
//   - []byte: Serialized data in requested format
//   - error: Serialization error if marshaling fails
//
// Note:
// - Uses json-iterator for JSON serialization (faster than stdlib)
// - Uses vmihailenco/msgpack for MessagePack serialization
func serialize(data any, format string) ([]byte, error) {
	switch format {
	case "msgpack":
		message, err := msgpack.Marshal(data)
		if err != nil {
			return nil, err
		}
		return message, nil

	default: // Default to JSON
		message, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		return message, nil
	}
}

// send delivers a pre-serialized response through supported context types.
//
// Parameters:
//   - ctx: Request context (Fiber, Echo, Gin or ByteSender compatible)
//   - json: Pre-serialized JSON response data
//   - msgpack: Pre-serialized MessagePack response data
//   - format: Optional response format specification
//
// Returns:
//   - error: Error if response sending fails
//
// Behavior:
// - Automatically detects context type
// - Sets appropriate Content-Type header
// - Supports multiple web frameworks through interfaces
// - Defaults to JSON format if none specified
// - Uses ByteSender as fallback for custom context types
//
// Content-Type handling:
// - "application/json" for JSON responses
// - "application/msgpack" for MessagePack
// - Custom content types if explicitly specified
func send(ctx any, json []byte, msgpack []byte, format ...string) error {
	contentType := "application/json"
	content := json

	if len(format) > 0 {
		switch format[0] {
		case "msgpack":
			contentType = "application/msgpack"
			content = msgpack
		case "json":
			contentType = "application/json"
		default:
			contentType = format[0] // Allow custom content types
		}
	}

	switch c := ctx.(type) {
	case FiberContext:
		c.Set("Content-Type", contentType)
		c.Status(200)
		return c.Write(content)

	case EchoContext:
		resp := c.Response()
		resp.Header().Set("Content-Type", contentType)
		resp.WriteHeader(200)
		_, err := resp.Write(content)
		return err

	case GinContext:
		writer := c.Writer()
		writer.Header().Set("Content-Type", contentType)
		c.Status(200)
		_, err := writer.Write(content)
		return err

	case ByteSender:
		return c.Send(content)

	default:
		return ErrorUnsupportedContextType
	}
}
