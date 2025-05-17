package util

import "slices"

func Contains[T comparable](slice []T, r T) bool {
	return slices.Contains(slice, r)
}
