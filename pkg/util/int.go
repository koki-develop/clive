package util

import "strconv"

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Digits(n int) int {
	return len(strconv.Itoa(n))
}
