package main

import (
	"fmt"

	"github.com/merijnf/everybody-codes-go/internal/e2024q01"
	"github.com/merijnf/everybody-codes-go/internal/e2024q02"
	"github.com/merijnf/everybody-codes-go/internal/e2024q03"
)

type QuestPart struct {
	quest int
	part  int
}

func main() {
	questPart := getQuestAndPart()
	var result string
	switch questPart {
	case QuestPart{1, 1}:
		result = e2024q01.SolveP1()
	case QuestPart{1, 2}:
		result = e2024q01.SolveP2()
	case QuestPart{1, 3}:
		result = e2024q01.SolveP3()
	case QuestPart{2, 1}:
		result = e2024q02.SolveP1()
	case QuestPart{2, 2}:
		result = e2024q02.SolveP2()
	case QuestPart{2, 3}:
		result = e2024q02.SolveP3()
	case QuestPart{3, 1}:
		result = e2024q03.SolveP1()
	case QuestPart{3, 2}:
		result = e2024q03.SolveP2()
	case QuestPart{3, 3}:
		result = e2024q03.SolveP3()
	default:
		result = "Not implemented"
	}

	fmt.Println("Solution:", result)
}

func getQuestAndPart() QuestPart {
	fmt.Println("Enter quest number (1-20):")
	var quest int
	_, err := fmt.Scan(&quest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Enter part number (1-3):")
	var part int
	_, err = fmt.Scan(&part)
	if err != nil {
		panic(err)
	}
	fmt.Println("Solving quest", quest, "part", part)
	return QuestPart{quest, part}
}
