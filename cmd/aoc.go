package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"aoc/days"
	"aoc/utils"
)

var (
	day       int
	part      int
	input     []string
	inputPath string
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Run the Advent of Code 2019 Program",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Validate the arguments

		// Load the input
		var err error
		if inputPath == "" {
			input, err = utils.LoadInput(day)
		} else {
			input, err = utils.LoadInputFromPath(inputPath)
		}
		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		out, err := days.Run(day, part, input)
		if err != nil {
			return err
		}

		fmt.Println(out)
		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "Day to run")
	rootCmd.PersistentFlags().IntVarP(&part, "part", "p", 0, "Part to run")
	rootCmd.PersistentFlags().StringVarP(&inputPath, "input", "i", "", "Path to the input file")
}

// Execute executes the root Cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
