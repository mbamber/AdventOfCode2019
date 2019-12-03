package days

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
	}

	return funcList[day][part](input)
}
