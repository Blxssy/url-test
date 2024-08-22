package random

import (
	"crypto/rand"
	"math/big"
)

func NewRandomString(size int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789")

	b := make([]rune, size)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		b[i] = chars[num.Int64()]
	}

	return string(b)
}
