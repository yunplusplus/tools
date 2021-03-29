package base64

import (
	"encoding/base64"

)

func Encode64(content string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(content))
	return encoded
}

func Decode64(content string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(content)
	return string(decoded), err
}
