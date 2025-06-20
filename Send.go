package msg

// Send is a handler function that serializes the given data into JSON format
// and sends it as a response. It accepts any data type that can be marshaled to JSON.
func Send(ctx any, data any, format ...string) (err error) {

	contentType := "json"
	if len(format) > 0 {
		switch format[0] {
		case "msgpack":
			contentType = "msgpack"
		case "json":
			contentType = "json"
		default:
			contentType = format[0]
		}
	}

	var messageJson []byte
	var messageMsgpack []byte

	if contentType == "json" {

		// Attempt to marshal the given data into JSON.
		messageJson, err = serialize(&responseMessage{Response: data}, contentType)
		if err != nil {
			// Return the error if marshalling fails.
			return err
		}

	}

	if contentType == "msgpack" {

		// Attempt to marshal the given data into JSON.
		messageMsgpack, err = serialize(&responseMessage{Response: data}, contentType)
		if err != nil {
			// Return the error if marshalling fails.
			return err
		}

	}

	if err != nil {
		return err
	}

	// Return the JSON serialized message as a response with the HTTP status code 200.
	return send(ctx, messageJson, messageMsgpack, format...)

}
