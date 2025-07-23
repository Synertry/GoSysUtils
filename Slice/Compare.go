package Slice

// Contains searches for an element in a slice and returns true if found
// Source: https://stackoverflow.com/a/70802740/5516320
//
// Deprecated: As of Go 1.21, you can use the stdlib slices package,
// which was promoted from the experimental package:
// https://stackoverflow.com/a/71181131/5516320
//
// This makes this function redundant, instead use
// slices.Contains()
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
