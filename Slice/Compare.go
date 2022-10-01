package Slice

// Contains searches for an element in a slice and returns true if found
// Source: https://stackoverflow.com/a/70802740/5516320
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
