package e2024q01

import (
	"fmt"
	"strconv"

	input "github.com/merijnf/everybody-codes-go/pkg"
)

func SolveP1() string {
	input := input.ReadInput(2024, 1, 1)
	result := 0
	for _, creature := range input {
		if creature == 'A' {
		}
		if creature == 'B' {
			result = result + 1
		}
		if creature == 'C' {
			result = result + 3
		}
	}
	return strconv.Itoa(result)
}

func SolveP2() string {
	input := input.ReadInput(2024, 1, 2)
	result := 0
	var pair [2]rune
	for i, char := range input {
		if i%2 == 0 {
			pair[0] = char
		} else {
			pair[1] = char
			asPair := true
			subResult := 0
			for _, creature := range pair {

				if creature == 'A' {

				}
				if creature == 'B' {
					subResult = subResult + 1
				}
				if creature == 'C' {
					subResult = subResult + 3
				}
				if creature == 'D' {
					subResult = subResult + 5
				}
				if creature == 'x' {
					asPair = false
				}
			}
			if asPair {
				subResult = subResult + 2
			}
			fmt.Printf("%c%c %d\n", pair[0], pair[1], subResult)
			result = result + subResult
		}

	}
	return strconv.Itoa(result)
}

func SolveP3() string {
	input := input.ReadInput(2024, 1, 3)
	result := 0
	var set [3]rune
	index := 0
	for _, char := range input {
		set[index] = char
		index++
		if index == 3 {
			index = 0
			creatureCount := 0
			subResult := 0
			for _, creature := range set {
				if creature == 'A' {
					creatureCount = creatureCount + 1
				}
				if creature == 'B' {
					subResult = subResult + 1
					creatureCount = creatureCount + 1
				}
				if creature == 'C' {
					subResult = subResult + 3
					creatureCount = creatureCount + 1
				}
				if creature == 'D' {
					subResult = subResult + 5
					creatureCount = creatureCount + 1
				}
			}
			if creatureCount == 3 {
				subResult = subResult + 6
			}
			if creatureCount == 2 {
				subResult = subResult + 2
			}

			fmt.Printf("%c%c%c %d\n", set[0], set[1], set[2], subResult)
			result = result + subResult
		}

	}
	return strconv.Itoa(result)
}
