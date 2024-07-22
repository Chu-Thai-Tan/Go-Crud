package util

import (
	"math/rand/v2"
	"strings"
)

func GenerateUID() string {
	var s [36]string
	const hexDigits = "0123456789abcdef"
	runes := []rune(hexDigits)

	for i := range s {
		var a = rand.IntN(16)
		s[i] = string(runes[a])
	}

	s[14] = "4"
	s[19] = string(runes[([]byte(s[19])[0]&0x3)|0x8])
	s[8] = "-"
	s[13] = "-"
	s[18] = "-"
	s[23] = "-"

	return strings.Join(s[:], "")
}
