package util

func Contains[T comparable](slice []T, r T) bool {
	for _, l := range slice {
		if l == r {
			return true
		}
	}
	return false
}
