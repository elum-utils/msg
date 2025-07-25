// Code generated by generate.go; DO NOT EDIT.
// This file was automatically generated by generate.go
// Manual edits will be overwritten on next generation.

package msg

// Pre-serialized JSON representation of "InvalidFields" error.
// Contains byte array representing JSON structure:
// {
//   "error": {
//     "key": "INVALID_FIELDS",
//     "message": "Required fields are missing or their type does not match the declared one"
//   }
// }
var errorInvalidFieldsJson = []byte{123, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 73, 78, 86, 65, 76, 73, 68, 95, 70, 73, 69, 76, 68, 83, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 82, 101, 113, 117, 105, 114, 101, 100, 32, 102, 105, 101, 108, 100, 115, 32, 97, 114, 101, 32, 109, 105, 115, 115, 105, 110, 103, 32, 111, 114, 32, 116, 104, 101, 105, 114, 32, 116, 121, 112, 101, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 109, 97, 116, 99, 104, 32, 116, 104, 101, 32, 100, 101, 99, 108, 97, 114, 101, 100, 32, 111, 110, 101, 34, 125, 125}

// Pre-serialized MessagePack representation of "InvalidFields" error.
// Binary-optimized format containing the same data as JSON version.
var errorInvalidFieldsMsgpack = []byte{129, 165, 101, 114, 114, 111, 114, 130, 163, 75, 101, 121, 174, 73, 78, 86, 65, 76, 73, 68, 95, 70, 73, 69, 76, 68, 83, 167, 77, 101, 115, 115, 97, 103, 101, 217, 73, 82, 101, 113, 117, 105, 114, 101, 100, 32, 102, 105, 101, 108, 100, 115, 32, 97, 114, 101, 32, 109, 105, 115, 115, 105, 110, 103, 32, 111, 114, 32, 116, 104, 101, 105, 114, 32, 116, 121, 112, 101, 32, 100, 111, 101, 115, 32, 110, 111, 116, 32, 109, 97, 116, 99, 104, 32, 116, 104, 101, 32, 100, 101, 99, 108, 97, 114, 101, 100, 32, 111, 110, 101}

// InvalidFields sends a pre-serialized error message to the client.
//
// Parameters:
//   - ctx: Request context (compatible with Fiber, Echo, Gin or other frameworks)
//   - format: Optional format parameter ("json" or "msgpack"). Defaults to JSON.
//
// Returns:
//   - error: Returns an error if there were problems sending the message
//
// Behavior:
// - Automatically sets appropriate Content-Type header
// - Supports multiple frameworks through unified interface
// - Uses pre-serialized data for maximum performance
// - Defaults to JSON format if none specified
func InvalidFields(ctx any, format ...string) error {
    return send(ctx, errorInvalidFieldsJson, errorInvalidFieldsMsgpack, format...)
}
