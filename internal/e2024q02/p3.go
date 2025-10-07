package e2024q02

import (
	"strconv"
	"strings"

	"github.com/merijnf/everybody-codes-go/pkg/input"
)

func SolveP3() string {
	input := input.ReadInput(2024, 2, 3)
	lines := strings.Split(input, "\r\n")
	words := strings.Split(strings.Split(lines[0], ":")[1], ",")
	grid := make([][]rune, len(lines[2:]))
	for y, line := range lines[2:] {
		grid[y] = []rune(line)
	}

	visited := make([][]bool, len(grid))
	for y := range grid {
		visited[y] = make([]bool, len(grid[y]))
	}

	for y, gridLine := range grid {
		for x := range gridLine {
			for _, word := range words {
				left := true
				right := true
				up := true
				down := true

				for i, char := range word {
					if i == 0 {
						if grid[y][x] != char {
							left = false
							right = false
							up = false
							down = false
							break
						}
					} else {
						if left {
							newX := x - i
							if newX < 0 {
								newX = len(gridLine) + newX
							}
							if grid[y][newX] != char {
								left = false
							}
						}
						if right {
							newX := x + i
							if newX >= len(gridLine) {
								newX = newX - len(gridLine)
							}
							if grid[y][newX] != char {
								right = false
							}
						}
						if up {
							newY := y - i
							if newY < 0 {
								up = false
							} else {
								if grid[newY][x] != char {
									up = false
								}
							}

						}
						if down {
							newY := y + i
							if newY >= len(grid) {
								down = false
							} else {
								if grid[newY][x] != char {
									down = false
								}
							}
						}

						if !left && !right && !up && !down {
							break
						}
					}
				}

				if left {
					for i := range len(word) {
						newX := x - i
						if newX < 0 {
							newX = len(gridLine) + newX
						}
						visited[y][newX] = true
					}
				}
				if right {
					for i := range len(word) {
						newX := x + i
						if newX >= len(gridLine) {
							newX = newX - len(gridLine)
						}
						visited[y][newX] = true
					}
				}
				if up {
					for i := range len(word) {
						newY := y - i
						visited[newY][x] = true
					}
				}
				if down {
					for i := range len(word) {
						newY := y + i
						visited[newY][x] = true
					}
				}
			}
		}
	}

	result := 0
	for y := range visited {
		for x := range visited[y] {
			if visited[y][x] {
				result++
			}
		}
	}
	return strconv.Itoa(result)
}
