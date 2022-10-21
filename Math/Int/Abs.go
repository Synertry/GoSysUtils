package Int

// Abs returns the absolute value of an integer
func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
