package random

import (
	"math/rand"
	"strconv"
)

func PwdRange(min, max int) string {
	n := rand.Intn(max-min) + min
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789()*&#?+-"
	letter := make([]byte, n)
	for i := range letter {
		letter[i] = charset[rand.Intn(len(charset))]
	}
	return string(letter)
}
func NumRange(min, max int) string {
	num := rand.Intn(max-min+1) + min
	return strconv.Itoa(num)
}
func LetterRange(min, max int) string {
	n := rand.Intn(max-min) + min
	charset := "abcdefghijklmnopqrstuvwxyz0123456789"
	letter := make([]byte, n)
	for i := range letter {
		letter[i] = charset[rand.Intn(len(charset))]
	}
	return string(letter)
}
