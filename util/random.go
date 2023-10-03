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

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0 -> max-min
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k-1)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomName generates a random string, can be used for testing for name, created by, updated by
func RandomName() string {
	return RandomString(6)
}

// RandomPrice generates a random price for a menu item
func RandomPrice() int64 {
	return RandomInt(0, 1000)
}

// RandomDescription generates a sequence of random strings separated by space, ended by period.
func RandomDescription() string {
	return RandomString(6) + " " + RandomString(6) + " " + RandomString(6) + "."
}