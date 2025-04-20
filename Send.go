package msg

import "github.com/labstack/echo/v4"

// Send is a handler function that serializes the given data into JSON format
// and sends it as a response. It accepts any data type that can be marshaled to JSON.
func Send(ctx echo.Context, data any) error {

	// Attempt to marshal the given data into JSON.
	message, err := json.Marshal(&responseMessage{Response: data})
	if err != nil {
		// Return the error if marshalling fails.
		return err
	}

	// Return the JSON serialized message as a response with the HTTP status code 200.
	return ctx.JSONBlob(200, message)
}
