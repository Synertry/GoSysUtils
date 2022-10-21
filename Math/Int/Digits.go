package Int

// CountDigits returns the number of digits in a number
func CountDigits(num int) (count int) {
	num = Abs(num)
	for num > 0 {
		num /= 10
		count++
	}
	return
}
