// Code generated by generate.go; DO NOT EDIT.
// This file was automatically generated by generate.go
// Manual edits will be overwritten on next generation.

package msg

// Pre-serialized JSON representation of "Unauthorized" error.
// Contains byte array representing JSON structure:
// {
//   "error": {
//     "key": "UNAUTHORIZED",
//     "message": "Unauthorized"
//   }
// }
var errorUnauthorizedJson = []byte{123, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 107, 101, 121, 34, 58, 34, 85, 78, 65, 85, 84, 72, 79, 82, 73, 90, 69, 68, 34, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 85, 110, 97, 117, 116, 104, 111, 114, 105, 122, 101, 100, 34, 125, 125}

// Pre-serialized MessagePack representation of "Unauthorized" error.
// Binary-optimized format containing the same data as JSON version.
var errorUnauthorizedMsgpack = []byte{129, 165, 101, 114, 114, 111, 114, 130, 163, 75, 101, 121, 172, 85, 78, 65, 85, 84, 72, 79, 82, 73, 90, 69, 68, 167, 77, 101, 115, 115, 97, 103, 101, 172, 85, 110, 97, 117, 116, 104, 111, 114, 105, 122, 101, 100}

// Unauthorized sends a pre-serialized error message to the client.
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
func Unauthorized(ctx any, format ...string) error {
    return send(ctx, errorUnauthorizedJson, errorUnauthorizedMsgpack, format...)
}
