package e2024q01

import (
	"strconv"

	"github.com/merijnf/everybody-codes-go/pkg/input"
)

func SolveP1() string {
	input := input.ReadInput(2024, 1, 1)
	result := 0
	for _, creature := range input {
		_, potions := requiredPotions(creature)
		result = result + potions
	}
	return strconv.Itoa(result)
}

func SolveP2() string {
	input := input.ReadInput(2024, 1, 2)
	result := 0
	index := 0
	creatureCount := 0

	for _, char := range input {
		isCreature, potions := requiredPotions(char)
		if isCreature {
			creatureCount = creatureCount + 1
			result = result + potions
		}
		index++
		if index == 2 {
			if creatureCount == 2 {
				result = result + 2
			}
			index = 0
			creatureCount = 0
		}
	}
	return strconv.Itoa(result)
}

func SolveP3() string {
	input := input.ReadInput(2024, 1, 3)
	result := 0
	index := 0
	creatureCount := 0

	for _, char := range input {
		isCreature, potions := requiredPotions(char)
		if isCreature {
			creatureCount = creatureCount + 1
			result = result + potions
		}
		index++
		if index == 3 {
			if creatureCount == 2 {
				result = result + 2
			}
			if creatureCount == 3 {
				result = result + 6
			}
			index = 0
			creatureCount = 0
		}
	}
	return strconv.Itoa(result)
}

func requiredPotions(creature rune) (bool, int) {
	isCreature := false
	potionCount := 0
	switch creature {
	case 'A':
		isCreature = true
	case 'B':
		isCreature = true
		potionCount = 1
	case 'C':
		isCreature = true
		potionCount = 3
	case 'D':
		isCreature = true
		potionCount = 5
	}

	return isCreature, potionCount
}
