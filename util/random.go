package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString generates a random string of length
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // min -> max
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomAccountId generates a random account id
func RandomAccountId() int64 {
	return RandomInt(0, 1000)
}

func RanddomCurrency() string {
	currencies := []string{"USD", "COP", "BTC"}
	k := len(currencies)

	return currencies[rand.Intn(k)]
}
