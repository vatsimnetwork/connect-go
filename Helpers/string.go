package helpers

import (
	"math/rand"
)

// RandString returns randomly generated string
func RandString(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	convertLength := make([]rune, length)
	for i := range convertLength {
		convertLength[i] = letters[rand.Intn(len(letters))]
	}
	return string(convertLength)
}
