package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().Unix())
}

// returns a random integer between min and max
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// return a random string with length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabets)
	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// return random owner with 6 length
func RandomOwner() string {
	return RandomString(6)
}

// return random money 0 - 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// return random currency
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "INR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
