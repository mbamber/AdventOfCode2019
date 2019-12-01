package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// LoadInput loads the input file for the given day
func LoadInput(day int) ([]string, error) {
	filePath := fmt.Sprintf("inputs/day%d.input", day)
	return LoadInputFromPath(filePath)
}

// LoadInputFromPath loads the input from the given path
func LoadInputFromPath(path string) ([]string, error) {
	out, err := ioutil.ReadFile(path)
	parts := strings.Split(string(out), "\n")
	return parts[:len(parts)-1], err
}
