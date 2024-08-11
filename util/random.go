package util

import (
	"math/rand"
	"strings"
	"time"
)

var rng *rand.Rand

var alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rng.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{USD, EUR, INR}
	n := len(currencies)
	return currencies[rng.Intn(n)]
}

func RandomCountryCode() int32 {
	return int32(RandomInt(1, 10000))
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}
