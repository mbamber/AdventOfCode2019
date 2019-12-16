package days

import (
	"aoc/opcode"
	"fmt"
	"strconv"
	"strings"
)

// Day7Part1 solves Day 7, Part 1
func Day7Part1(input []string) (string, error) {
	codeStrings := strings.Split(input[0], ",")
	codes := make([]int, len(codeStrings))
	for i, c := range codeStrings {
		x, err := strconv.Atoi(c)
		if err != nil {
			return "", err
		}
		codes[i] = x
	}

	numAmplifiers := 5

	phaseSettings := getPhaseSettings(numAmplifiers)
	maxOutput := 0
	for _, currPhaseSettings := range phaseSettings {
		// Make all the chans
		chans := map[int]struct {
			in  chan int
			out chan int
		}{}
		for amplifierID := 0; amplifierID < numAmplifiers; amplifierID++ {
			chans[amplifierID] = struct {
				in  chan int
				out chan int
			}{
				in:  make(chan int),
				out: make(chan int),
			}
		}

		lastOutput := 0
		for amplifierID := 0; amplifierID < numAmplifiers; amplifierID++ {
			codesCopy := append(codes[:0:0], codes...)

			phaseSetting := currPhaseSettings[amplifierID]

			in := chans[amplifierID].in
			out := chans[amplifierID].out

			go func() {
				opcode.Run(codesCopy, in, out)
			}()

			go func() {
				in <- phaseSetting
				in <- lastOutput
			}()

			output, open := <-out
			var prevOut int
			for open {
				prevOut = output
				output, open = <-out
			}
			lastOutput = prevOut
		}

		// Overwrite the max output if required
		if lastOutput > maxOutput {
			maxOutput = lastOutput
		}

	}

	return fmt.Sprintf("%d", maxOutput), nil
}

func getPhaseSettings(n int) [][]int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
