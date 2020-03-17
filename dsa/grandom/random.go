package grandom

import (
	"github.com/Pallinder/go-randomdata"
)

// Notice: it seems min & max is not included in random output
func RandomInt(min, max int) int {
	if min == max {
		return min
	}
	return randomdata.Number(min, max)
}

func RandomBool() bool {
	return RandomInt(0, 3) == 1
}

func RandomString(size int) string {
	return randomdata.RandStringRunes(size)
}

func RandomFloat(min, max, decimalPoint int) float64 {
	return randomdata.Decimal(min, max, decimalPoint)
}

func RandomSimplePassword(minLen, maxLen int) string {
	pwdchars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return RandomPassword(pwdchars, minLen, maxLen)
}

func RandomPassword(chars string, minLen, maxLen int) string {
	if len(chars) == 0 || maxLen <= 0 {
		return ""
	}
	if minLen <= 0 {
		minLen = 1
	}

	resultLen := RandomInt(minLen, maxLen)
	result := ""
	charsLen := len(chars)
	for i := 0; i < resultLen; i++ {
		idx := RandomInt(0, charsLen-1)
		result += string(chars[idx])
	}
	return result
}
