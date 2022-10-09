package Math

// IntPow calculates base to the expth power. Since the result is an int, it is assumed that exp is a positive power
// returns 0 if m is negative
func IntPow(base, exp int) int {
	if exp < 0 {
		return 0
	} else if exp == 0 {
		return 1
	}
	result := base
	for i := 2; i <= exp; i++ {
		result *= base
	}
	return result
}
