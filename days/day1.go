package days

import (
	"fmt"
	"math"
	"strconv"
)

// Day1Part1 solves Day 1, Part 1
func Day1Part1(input []string) (output string, err error) {
	total := 0
	for _, val := range input {
		i, err := strconv.Atoi(val)
		if err != nil {
			return "", err
		}
		total += calculateFuelRequirement(i)
	}

	return fmt.Sprintf("%d", total), nil
}

func calculateFuelRequirement(mass int) int {
	return int(math.Floor(float64(mass/3))) - 2
}

// Day1Part2 solves Day 1, Part 2
func Day1Part2(input []string) (output string, err error) {
	total := 0
	for _, val := range input {
		i, err := strconv.Atoi(val)
		if err != nil {
			return "", err
		}
		total += calculateFuelRequirementWithAdditionalMass(i)
	}

	return fmt.Sprintf("%d", total), nil
}

func calculateFuelRequirementWithAdditionalMass(mass int) int {
	total := int(math.Floor(float64(mass/3))) - 2
	if total <= 0 {
		return 0
	}
	return total + calculateFuelRequirementWithAdditionalMass(total)
}
