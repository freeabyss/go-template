package util

import (
	"bytes"
	"fmt"
	"encoding/json"
	"testing"
	"strings"
)

func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func Concat(strList ...string) string {
	var buffer bytes.Buffer
	for _, v := range strList {
		buffer.WriteString(v)
	}
	result := buffer.String()
	return result
}

func ToString(src interface{}) string {
	str, err := json.Marshal(src)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(str)
}

func CanLogedBody(body string) string {
	rs := []rune(body)
	if len(rs) <= 1000 {
		return body
	}
	return string(rs[0:1000])
}

