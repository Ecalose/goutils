package random

import "math/rand"

func Num(n int) string {
	charset := "0123456789"
	num := make([]byte, n)
	for i := range num {
		num[i] = charset[rand.Intn(len(charset))]
	}
	return string(num)
}
func Letter(n int) string {
	charset := "abcdefghijklmnopqrstuvwxyz"
	letter := make([]byte, n)
	for i := range letter {
		letter[i] = charset[rand.Intn(len(charset))]
	}
	return string(letter)
}

func Pwd(n int) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789()*&#?+-"
	letter := make([]byte, n)
	for i := range letter {
		letter[i] = charset[rand.Intn(len(charset))]
	}
	return string(letter)
}
