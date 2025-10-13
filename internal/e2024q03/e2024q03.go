package e2024q03

import (
	"strconv"
	"strings"

	"github.com/merijnf/everybody-codes-go/pkg/input"
	"github.com/merijnf/everybody-codes-go/pkg/print"
)

type Coordinate struct {
	X int
	Y int
}

func SolveP1() string {
	input := input.ReadInput(2024, 3, 1)
	lines := strings.Split(input, "\r\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	dugPositions := make([]Coordinate, 0)

	print.PrintGrid(grid)
	blocksDug := 0
	// dig first layer
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				dugPositions = append(dugPositions, Coordinate{X: x, Y: y})
				grid[y][x] = '1'
				blocksDug++
			}
		}
	}

	newDugPositions := make([]Coordinate, 0)
	layer := 2
	prevLayerStr := '1'
	layerStr := '2'
	print.PrintGrid(grid)
	for len(dugPositions) > 0 {
		for _, pos := range dugPositions {
			// check neighbors

			if (pos.X > 0 && (grid[pos.Y][pos.X-1] == prevLayerStr || (grid[pos.Y][pos.X-1] == layerStr))) &&
				(pos.X < len(grid[0])-1 && (grid[pos.Y][pos.X+1] == prevLayerStr || grid[pos.Y][pos.X+1] == layerStr)) &&
				(pos.Y > 0 && (grid[pos.Y-1][pos.X] == prevLayerStr || grid[pos.Y-1][pos.X] == layerStr)) &&
				(pos.Y < len(grid)-1 && (grid[pos.Y+1][pos.X] == prevLayerStr || grid[pos.Y+1][pos.X] == layerStr)) {
				// all neighbors dug, dig this one too
				newDugPositions = append(newDugPositions, Coordinate{X: pos.X, Y: pos.Y})
				grid[pos.Y][pos.X] = layerStr
				blocksDug++
			}
		}
		dugPositions = newDugPositions
		newDugPositions = make([]Coordinate, 0)
		prevLayerStr = rune(strconv.Itoa(layer)[0])
		layer++
		layerStr = rune(strconv.Itoa(layer)[0])
		print.PrintGrid(grid)
	}

	return strconv.Itoa(blocksDug)
}

func SolveP2() string {
	input := input.ReadInput(2024, 3, 2)
	lines := strings.Split(input, "\r\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	dugPositions := make([]Coordinate, 0)

	print.PrintGrid(grid)
	blocksDug := 0
	// dig first layer
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				dugPositions = append(dugPositions, Coordinate{X: x, Y: y})
				grid[y][x] = '1'
				blocksDug++
			}
		}
	}

	newDugPositions := make([]Coordinate, 0)
	layer := 2
	prevLayerStr := '1'
	layerStr := '2'
	print.PrintGrid(grid)
	for len(dugPositions) > 0 {
		for _, pos := range dugPositions {
			// check neighbors

			if (pos.X > 0 && (grid[pos.Y][pos.X-1] == prevLayerStr || (grid[pos.Y][pos.X-1] == layerStr))) &&
				(pos.X < len(grid[0])-1 && (grid[pos.Y][pos.X+1] == prevLayerStr || grid[pos.Y][pos.X+1] == layerStr)) &&
				(pos.Y > 0 && (grid[pos.Y-1][pos.X] == prevLayerStr || grid[pos.Y-1][pos.X] == layerStr)) &&
				(pos.Y < len(grid)-1 && (grid[pos.Y+1][pos.X] == prevLayerStr || grid[pos.Y+1][pos.X] == layerStr)) {
				// all neighbors dug, dig this one too
				newDugPositions = append(newDugPositions, Coordinate{X: pos.X, Y: pos.Y})
				grid[pos.Y][pos.X] = layerStr
				blocksDug++
			}
		}
		dugPositions = newDugPositions
		newDugPositions = make([]Coordinate, 0)
		prevLayerStr = rune(strconv.Itoa(layer)[0])
		layer++
		layerStr = rune(strconv.Itoa(layer)[0])
		print.PrintGrid(grid)
	}

	return strconv.Itoa(blocksDug)
}

func SolveP3() string {
	input := input.ReadInput(2024, 3, 3)
	lines := strings.Split(input, "\r\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	dugPositions := make([]Coordinate, 0)

	print.PrintGrid(grid)
	blocksDug := 0
	// dig first layer
	for y, row := range grid {
		for x, cell := range row {
			if cell == '#' {
				dugPositions = append(dugPositions, Coordinate{X: x, Y: y})
				grid[y][x] = '1'
				blocksDug++
			}
		}
	}

	newDugPositions := make([]Coordinate, 0)
	layer := 2
	prevLayerStr := '1'
	layerStr := '2'
	print.PrintGrid(grid)
	for len(dugPositions) > 0 {
		for _, pos := range dugPositions {
			// check neighbors

			if (pos.X > 0 && (grid[pos.Y][pos.X-1] == prevLayerStr || (grid[pos.Y][pos.X-1] == layerStr))) &&
				(pos.X < len(grid[0])-1 && (grid[pos.Y][pos.X+1] == prevLayerStr || grid[pos.Y][pos.X+1] == layerStr)) &&
				(pos.Y > 0 && (grid[pos.Y-1][pos.X] == prevLayerStr || grid[pos.Y-1][pos.X] == layerStr)) &&
				(pos.Y < len(grid)-1 && (grid[pos.Y+1][pos.X] == prevLayerStr || grid[pos.Y+1][pos.X] == layerStr)) &&
				(pos.X > 0 && pos.Y > 0 && (grid[pos.Y-1][pos.X-1] == prevLayerStr || (grid[pos.Y-1][pos.X-1] == layerStr))) &&
				(pos.X > 0 && pos.Y < len(grid)-1 && (grid[pos.Y+1][pos.X-1] == prevLayerStr || grid[pos.Y+1][pos.X-1] == layerStr)) &&
				(pos.X < len(grid[0])-1 && pos.Y > 0 && (grid[pos.Y-1][pos.X+1] == prevLayerStr || grid[pos.Y-1][pos.X+1] == layerStr)) &&
				(pos.X < len(grid[0])-1 && pos.Y < len(grid)-1 && (grid[pos.Y+1][pos.X+1] == prevLayerStr || grid[pos.Y+1][pos.X+1] == layerStr)) {
				// all neighbors dug, dig this one too
				newDugPositions = append(newDugPositions, Coordinate{X: pos.X, Y: pos.Y})
				grid[pos.Y][pos.X] = layerStr
				blocksDug++
			}
		}
		dugPositions = newDugPositions
		newDugPositions = make([]Coordinate, 0)
		prevLayerStr = rune(strconv.Itoa(layer)[0])
		layer++
		layerStr = rune(strconv.Itoa(layer)[0])
		print.PrintGrid(grid)
	}

	return strconv.Itoa(blocksDug)
}
