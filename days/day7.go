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
	phaseSettings := getPhaseSettings(5)

	maxOutput := 0
	chans := map[int]struct {
		in  chan int
		out chan int
	}{}
	for _, currPhaseSettings := range phaseSettings {
		// First setup all the chans for these phase settings
		for amplifierID := 0; amplifierID < numAmplifiers; amplifierID++ {
			chans[amplifierID] = struct {
				in  chan int
				out chan int
			}{
				in:  make(chan int, 20),
				out: make(chan int, 20),
			}
		}

		// Keep track of the output for these phase settings
		thisOutput := 0

		// Send the initial phase settings to the correct channels
		for amplifierID, phaseSetting := range currPhaseSettings {
			chans[amplifierID].in <- phaseSetting
		}

		// Send 0 to the first amplifier
		chans[0].in <- 0

		// Start the amplifiers
		for amplifierID := 0; amplifierID < numAmplifiers; amplifierID++ {
			go opcode.Run(append(codes[:0:0], codes...), chans[amplifierID].in, chans[amplifierID].out)
		}

		// Link the amplifiers
		for amplifierID := 0; amplifierID < numAmplifiers-1; amplifierID++ {
			sendsTo := amplifierID + 1

			go func(id, to int) {
				var output int
				for output = range chans[id].out {
					chans[to].in <- output
				}
			}(amplifierID, sendsTo)
		}

		thisOutput = <-chans[numAmplifiers-1].out // Block on the final amplifier output
		if thisOutput > maxOutput {
			maxOutput = thisOutput
		}
	}

	return fmt.Sprintf("%d", maxOutput), nil
}

func getPhaseSettings(n int) [][]int {
	return getPhaseSettingsBetween(0, n)
}

func getPhaseSettingsBetween(lower, upper int) [][]int {
	arr := make([]int, upper-lower)
	for i := lower; i < upper; i++ {
		arr[i-lower] = i
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

// Day7Part2 solves Day 7, Part 2
func Day7Part2(input []string) (string, error) {
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
	phaseSettings := getPhaseSettingsBetween(5, 10)

	maxOutput := 0
	chans := map[int]struct {
		in  chan int
		out chan int
	}{}
	for _, currPhaseSettings := range phaseSettings {
		// First setup all the chans for these phase settings
		for amplifierID := 0; amplifierID < numAmplifiers; amplifierID++ {
			chans[amplifierID] = struct {
				in  chan int
				out chan int
			}{
				in:  make(chan int, 20),
				out: make(chan int, 20),
			}
		}
		done := make(chan bool)

		// Keep track of the output for these phase settings
		thisOutput := 0

		// Send the initial phase settings to the correct channels
		for amplifierID, phaseSetting := range currPhaseSettings {
			chans[amplifierID].in <- phaseSetting
		}

		// Send 0 to the first amplifier
		chans[0].in <- 0

		// Start the amplifiers
		for amplifierID := 0; amplifierID < numAmplifiers; amplifierID++ {
			go opcode.Run(append(codes[:0:0], codes...), chans[amplifierID].in, chans[amplifierID].out)
		}

		// Link the amplifiers
		for amplifierID := 0; amplifierID < numAmplifiers; amplifierID++ {
			sendsTo := amplifierID + 1
			if sendsTo >= numAmplifiers {
				sendsTo = 0
			}

			go func(id, to int) {
				var output int
				for output = range chans[id].out {
					chans[to].in <- output
				}

				if id == numAmplifiers-1 {
					thisOutput = output
					done <- true
				}
			}(amplifierID, sendsTo)
		}

		<-done // Block on completion of the sequence

		if thisOutput > maxOutput {
			maxOutput = thisOutput
		}
	}

	return fmt.Sprintf("%d", maxOutput), nil
}
