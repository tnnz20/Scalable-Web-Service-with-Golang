package helpers

import "math/rand"

type RandSlice struct {
	num1 int
	num2 int
}

func RandomNum(min, max int) (result *RandSlice) {
	result = &RandSlice{
		num1: rand.Intn(max-min) + min,
		num2: rand.Intn(max-min) + min,
	}
	return
}
