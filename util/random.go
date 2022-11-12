package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generate a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generate random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generate a random name
func RandomName() string {
	return RandomString(6)
}

// RandomAmount generate a random amount
func RandomAmount() int64 {
	return RandomInt(1, 1000)
}

// RandomStatus generate a random status name
func RandomStatus() string {
	statuses := []string{"reserved", "confirmed", "canceled"}
	n := len(statuses)
	return statuses[rand.Intn(n)]
}
