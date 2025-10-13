package print

import "fmt"

func PrintGrid(grid [][]rune) {
	fmt.Println()
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}
