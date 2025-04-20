package msg

// errorMessage is a struct that wraps error data for JSON serialization.
// The JSON object will have an "error" field containing the error details.
type errorMessage struct {
    Error errorData `json:"error"` // Contains an errorData struct representing the error information.
}

// responseMessage is a struct that wraps a generic response for JSON serialization.
// The JSON object will have a "response" field containing the actual response data.
type responseMessage struct {
    Response any `json:"response"` // Uses "any" to allow any data type as the response, providing flexibility.
}

// errorData is a struct that defines the structure of error information.
// It includes a key for identification and a message for additional context.
type errorData struct {
    Key     string `json:"key"`     // Unique key identifying the type of error.
    Message string `json:"message"` // Descriptive message providing details about the error.
}