package utils

import (
	"crypto/md5"
	"strings"
)

// Md5Str 生成sessionID
func Md5Str(bt [md5.Size]byte) string {
	result := ""
	for _, i := range bt {
		result += string(i)
	}
	return strings.ToLower(result)
}
