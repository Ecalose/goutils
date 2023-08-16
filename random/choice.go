package random

import "math/rand"

func Choice[T string | int | int8 | int16 | int32 | int64 | float32 | float64 | uint | uint8 | uint16 | uint32 | uint64 | bool](s []T) T {
	return s[rand.Intn(len(s))]
}
