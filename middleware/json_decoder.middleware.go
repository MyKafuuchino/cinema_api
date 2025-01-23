package middleware

import (
	"bytes"
	"encoding/json"
)

func JSONDecoder() func(data []byte, v interface{}) error {
	return func(data []byte, v interface{}) error {
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.DisallowUnknownFields()
		return decoder.Decode(v)
	}
}
