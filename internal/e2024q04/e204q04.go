package e2024q04

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/merijnf/everybody-codes-go/pkg/input"
)

func SolveP1() string {
	input := input.ReadInput(2024, 4, 1)
	lines := strings.Split(input, "\r\n")
	depths := make([]int, len(lines))
	smallest := math.MaxInt
	for i, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		depths[i] = depth
		if depth < smallest {
			smallest = depth
		}
	}
	result := 0
	for _, depth := range depths {
		hits := depth - smallest
		fmt.Println(depth, smallest, hits)
		result = result + hits
	}
	return strconv.Itoa(result)
}

func SolveP2() string {
	input := input.ReadInput(2024, 4, 2)
	lines := strings.Split(input, "\r\n")
	depths := make([]int, len(lines))
	smallest := math.MaxInt
	for i, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		depths[i] = depth
		if depth < smallest {
			smallest = depth
		}
	}
	result := 0
	for _, depth := range depths {
		hits := depth - smallest
		fmt.Println(depth, smallest, hits)
		result = result + hits
	}
	return strconv.Itoa(result)
}

func SolveP3() string {
	input := input.ReadInput(2024, 4, 3)
	lines := strings.Split(input, "\r\n")
	depths := make([]int, len(lines))
	// convert to ints
	for i, line := range lines {
		depth, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		depths[i] = depth
	}
	// sort
	slices.Sort(depths)
	// find median
	median := depths[len(depths)/2]
	fmt.Println("Median:", median)
	result := 0
	for _, depth := range depths {
		if depth < median {
			result = result + (median - depth)
		} else {
			result = result + (depth - median)
		}
	}

	return strconv.Itoa(result)
}
