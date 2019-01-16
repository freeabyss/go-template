package util

import (
	"fmt"
	"encoding/base64"
)

func Base64(data string) string {
	fmt.Println("base64前的信息: ", data)
	bytes := []byte(data)
	result := base64.StdEncoding.EncodeToString(bytes)
	fmt.Println("base64后的信息: ", result)
	return result
}
