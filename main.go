package main

import (
	"advent-of-code24/solutions"
	"advent-of-code24/utils"
	"fmt"
)

func main() {
	day := 1 // Change this to run a different day
	inputFile := fmt.Sprintf("inputs/day%02d.txt", day)
	input, err := utils.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	switch day {
	case 1:
		solutions.SolveDay01(input)
	default:
		fmt.Println("Day not implemented")
	}
}
