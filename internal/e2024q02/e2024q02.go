package e2024q02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/merijnf/everybody-codes-go/pkg/input"
	"github.com/merijnf/everybody-codes-go/pkg/stringutil"
)

func SolveP1() string {
	input := input.ReadInput(2024, 2, 1)
	lines := strings.Split(input, "\r\n")
	words := strings.Split(strings.Split(lines[0], ":")[1], ",")
	inscription := lines[2]
	fmt.Println("Words:", words)

	result := 0
	for _, word := range words {
		result += strings.Count(inscription, word)
	}
	return strconv.Itoa(result)
}

func SolveP2() string {
	input := input.ReadInput(2024, 2, 2)
	lines := strings.Split(input, "\r\n")
	originalWords := strings.Split(strings.Split(lines[0], ":")[1], ",")
	words := make([]string, len(originalWords)*2)
	copy(words, originalWords)
	for i, word := range originalWords {
		words[i+len(originalWords)] = stringutil.Reverse(word)
	}
	inscriptions := lines[2:]

	fmt.Println("Words:", words)

	result := 0
	for _, inscription := range inscriptions {
		runeAtIndex := make(map[int]bool)
		for _, word := range words {
			index := strings.Index(inscription, word)
			for index != -1 {
				for i := range word {
					runeAtIndex[index+i] = true
				}

				subindex := strings.Index(inscription[index+1:], word)
				if subindex == -1 {
					index = -1
				} else {
					index = subindex + index + 1
				}

			}
		}
		result += len(runeAtIndex)
	}
	return strconv.Itoa(result)
}

func SolveP3() string {
	return "not implemented"
}
