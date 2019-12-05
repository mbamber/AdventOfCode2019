package utils

import "math"

// GetDigits gets each individual digit from a number n, returning them in the order of the number
func GetDigits(n int) (digits []int) {
	backwards := []int{}
	for n != 0 {
		rem := int(math.Mod(float64(n), float64(10)))
		backwards = append(backwards, rem)
		n /= 10
	}

	digits = make([]int, len(backwards))
	for i := len(backwards); i > 0; i-- {
		digits[len(backwards)-i] = backwards[i-1]
	}
	return digits
}
