package days

import (
	"fmt"
)

// Day8Part1 solves Day 8, Part 1
func Day8Part1(input []string) (string, error) {

	imData := input[0]
	imWidth, imHeight := 25, 6

	layers := generateLayers(imData, imWidth, imHeight)

	var layerWithFewestZeros [][]int
	fewestZeros := imWidth * imHeight
	for _, layer := range layers {

		zeros := 0
		for y := 0; y < len(layer); y++ {
			for x := 0; x < len(layer[y]); x++ {
				if layer[y][x] == 0 {
					zeros++
				}
			}
		}

		if zeros < fewestZeros {
			fewestZeros = zeros
			layerWithFewestZeros = layer
		}
	}

	ones, twos := 0, 0
	for y := 0; y < len(layerWithFewestZeros); y++ {
		for x := 0; x < len(layerWithFewestZeros[y]); x++ {
			switch layerWithFewestZeros[y][x] {
			case 1:
				ones++
			case 2:
				twos++
			}
		}
	}

	return fmt.Sprintf("%d", ones*twos), nil
}

func generateLayers(imData string, width, height int) [][][]int {
	layerSize := width * height

	numLayers := len(imData) / layerSize

	layerStrings := []string{}
	for i := 0; i < numLayers; i++ {
		layerStrings = append(layerStrings, imData[i*layerSize:(i+1)*layerSize])
	}

	layers := [][][]int{}
	for _, layerString := range layerStrings {

		layer := [][]int{}
		for y := 0; y < height; y++ {
			row := []int{}
			for x := 0; x < width; x++ {
				row = append(row, int(layerString[(width*y)+x]-'0'))
			}
			layer = append(layer, row)
		}
		layers = append(layers, layer)
	}

	return layers
}

// Day8Part2 solves Day 8, Part 2
func Day8Part2(input []string) (string, error) {

	imData := input[0]
	imWidth, imHeight := 25, 6

	layers := generateLayers(imData, imWidth, imHeight)

	message := ""
	for y := 0; y < imHeight; y++ {
		for x := 0; x < imWidth; x++ {
			message += getMessagePixel(layers, x, y)
		}
		message += "\n"
	}

	return message, nil
}

func getMessagePixel(layers [][][]int, x, y int) string {
	for _, layer := range layers {
		pix := layer[y][x]
		switch pix {
		case 0:
			return " "
		case 1:
			return "X"
		default:
			continue
		}
	}
	return " "
}
