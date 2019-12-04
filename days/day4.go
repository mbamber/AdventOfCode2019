package days

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Day4Part1 solves Day 4, Part 1
func Day4Part1(input []string) (string, error) {
	parts := strings.Split(input[0], "-")
	lower, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}
	upper, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	valid := []int{}
	for i := lower; i <= upper; i++ {
		digits := getDigits(i)
		if digitsNeverDecrease(digits) && existsSameDigitNextToEachother(digits) {
			valid = append(valid, i)
		}
	}

	return fmt.Sprintf("%d", len(valid)), nil
}

func getDigits(n int) (digits []int) {
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

func digitsNeverDecrease(digits []int) bool {
	prevD := digits[0]
	for _, d := range digits {
		if d < prevD {
			return false
		}
		prevD = d
	}
	return true
}

func existsSameDigitNextToEachother(digits []int) bool {
	prevD := -1
	for _, d := range digits {
		if d == prevD {
			return true
		}
		prevD = d
	}
	return false
}

func existsExactlyTwoDigitsNextToEachother(digits []int) bool {
	currD := digits[0]
	currCount := 1
	for i, d := range digits {
		if i == 0 {
			continue // Skip the first number because we already set that
		}

		if d != currD {
			if currCount == 2 {
				return true
			}
			currD = d
			currCount = 1
			continue
		}
		currCount++
	}
	return currCount == 2
}

// Day4Part2 solves Day 4, Part 2
func Day4Part2(input []string) (string, error) {
	parts := strings.Split(input[0], "-")
	lower, err := strconv.Atoi(parts[0])
	if err != nil {
		return "", err
	}
	upper, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", err
	}

	valid := []int{}
	for i := lower; i <= upper; i++ {
		digits := getDigits(i)
		if digitsNeverDecrease(digits) && existsExactlyTwoDigitsNextToEachother(digits) {
			valid = append(valid, i)
		}
	}

	return fmt.Sprintf("%d", len(valid)), nil
}
