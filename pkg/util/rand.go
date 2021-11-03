package util

import (
	"math/rand"
	"time"
)

var (
	alphabets  = []rune("abcdefghijklmonpqrstuvwxyzABCDEFGHIJKLMONPQRSTUVWXYZ")
	digits     = []rune("0123456789")
	characters = []rune("!@#$%&/()=?*~")
)

// RandomKey returns a random key
// You can pass lower limit and upper limit of the key length
// A key can container alpha numeric character
// A key always starts with an alphabet
func RandomKey(l, u int) string {
	chars := append(alphabets, digits...)
	rand.Seed(time.Now().UnixNano())
	// create a rune slice of random length
	o := make([]rune, l+rand.Intn(u-l))
	// Make the first character an alphabet
	o[0] = alphabets[rand.Intn(len(alphabets))]
	for i := 1; i < len(o); i++ {
		index := rand.Intn(len(chars))
		o[i] = chars[index]
	}
	return string(o)
}

// RandomValue returns a random value
// You can pass lower limit and upper limit of the value length
// A value can contain alphabets, digits and special characters
func RandomValue(l, u int) string {
	chars := append(alphabets, append(digits, characters...)...)
	rand.Seed(time.Now().UnixNano())
	// create a rune slice of random length
	o := make([]rune, l+rand.Intn(u-l))
	for i := 0; i < len(o); i++ {
		index := rand.Intn(len(chars))
		o[i] = chars[index]
	}
	return string(o)
}
