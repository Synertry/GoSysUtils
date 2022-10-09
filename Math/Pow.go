package Math

// IntPow calculates n to the mth power. Since the result is an int, it is assumed that m is a positive power
func IntPow(base, exponent int) (result int) {
	if exponent == 0 {
		return 1
	}
	result = base
	for i := 2; i <= exponent; i++ {
		result *= base
	}
	return
}
