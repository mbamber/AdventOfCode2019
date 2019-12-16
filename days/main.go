package days

import "fmt"

type partMap map[int]func(input []string) (string, error)

// Run calls the correct function for the given day and part
func Run(day, part int, input []string) (string, error) {
	// Build a map containing all the functions
	funcList := map[int]partMap{
		1: partMap{
			1: Day1Part1,
			2: Day1Part2,
		},
		2: partMap{
			1: Day2Part1,
			2: Day2Part2,
		},
		3: partMap{
			1: Day3Part1,
			2: Day3Part2,
		},
		4: partMap{
			1: Day4Part1,
			2: Day4Part2,
		},
		5: partMap{
			1: Day5Part1,
			2: Day5Part2,
		},
		6: partMap{
			1: Day6Part1,
			2: Day6Part2,
		},
		7: partMap{
			1: Day7Part1,
			2: Day7Part2,
		},
		8: partMap{
			1: Day8Part1,
			2: Day8Part2,
		},
	}

	f, ok := funcList[day][part]
	if !ok {
		return "", fmt.Errorf("No function defined for Day %d, Part %d", day, part)
	}
	return f(input)
}
