package msg

import (
    jsoniter "github.com/json-iterator/go"
)

// json is an instance of json-iterator configured to be compatible with the standard library.
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// serialize is a function that takes an errorData struct and serializes it into a JSON byte slice.
// It returns an empty byte slice if serialization fails.
func serialize(data errorData) []byte {
    // Marshal the errorMessage struct containing the provided errorData into a JSON byte slice.
    message, err := json.Marshal(&errorMessage{Error: data})
    if err != nil {
        // Return an empty byte slice if an error occurs during JSON serialization.
        return []byte{}
    }
    // Return the serialized JSON message.
    return message
}