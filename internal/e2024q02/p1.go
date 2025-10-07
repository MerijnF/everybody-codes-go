package e2024q02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/merijnf/everybody-codes-go/pkg/input"
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
