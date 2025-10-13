package e2024q03

import (
	"strconv"
	"strings"

	"github.com/merijnf/everybody-codes-go/pkg/input"
	"github.com/merijnf/everybody-codes-go/pkg/print"
)

type Vector struct {
	X int
	Y int
}

var allDirections = []Vector{
	{X: 0, Y: -1},  // up
	{X: 1, Y: -1},  // up-right
	{X: 1, Y: 0},   // right
	{X: 1, Y: 1},   // down-right
	{X: 0, Y: 1},   // down
	{X: -1, Y: 1},  // down-left
	{X: -1, Y: 0},  // left
	{X: -1, Y: -1}, // up-left
}

var cardinalDirections = []Vector{
	{X: 0, Y: -1}, // up
	{X: 1, Y: 0},  // right
	{X: 0, Y: 1},  // down
	{X: -1, Y: 0}, // left
}

func dig(part int, checkDirections []Vector) string {
	input := input.ReadInput(2024, 3, part)
	lines := strings.Split(input, "\r\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	dugPositions := make([]Vector, 0)

	print.PrintGrid(grid)
	blocksDug := 0
	// dig first layer
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				dugPositions = append(dugPositions, Vector{X: x, Y: y})
				grid[y][x] = '1'
				blocksDug++
			}
		}
	}

	newDugPositions := make([]Vector, 0)
	layer := 2
	prevLayerStr := '1'
	layerStr := '2'
	print.PrintGrid(grid)
	for len(dugPositions) > 0 {
		for _, pos := range dugPositions {
			// check neighbors
			ok := true
			for _, dir := range checkDirections {
				nx := pos.X + dir.X
				ny := pos.Y + dir.Y
				if nx >= 0 && nx <= len(grid[0])-1 && ny >= 0 && ny <= len(grid)-1 {
					// in bounds
					posCell := grid[ny][nx]
					if !(posCell == prevLayerStr || posCell == layerStr) {
						ok = false
						break
					}
				} else {
					// out of bounds
					ok = false
					break
				}
			}
			if ok {
				// all neighbors dug, dig this one too
				newDugPositions = append(newDugPositions, Vector{X: pos.X, Y: pos.Y})
				grid[pos.Y][pos.X] = layerStr
				blocksDug++
			}
		}
		dugPositions = newDugPositions
		newDugPositions = make([]Vector, 0)
		prevLayerStr = rune(strconv.Itoa(layer)[0])
		layer++
		layerStr = rune(strconv.Itoa(layer)[0])
		print.PrintGrid(grid)
	}

	return strconv.Itoa(blocksDug)
}

func SolveP1() string {
	return dig(1, cardinalDirections)
}

func SolveP2() string {
	return dig(2, cardinalDirections)
}

func SolveP3() string {
	return dig(3, allDirections)
}
