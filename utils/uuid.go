package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func GenerUUID() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("gener UUID failed: ", err)
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}
