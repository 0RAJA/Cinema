package utils

import (
	"math/rand"
	"time"
)

const Length = 6

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes 生成邀请码
func RandStringRunes() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, Length)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
