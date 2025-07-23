package Slice

func RemoveIndex[T comparable](s []T, idx int) []T {
	if idx < 0 || idx >= len(s) {
		return s // Index out of bounds, return original slice
	}
	s[idx] = s[len(s)-1] // Move the last element to the index
	return s[:len(s)-1]  // Return the slice without the last element
}

func RemoveElement[T comparable](s []T, e T) []T {
	for i, v := range s {
		if v == e {
			s = RemoveIndex(s, i) // Use RemoveIndex to remove the element
		}
	}
	return s
}
