package Day1 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
var inputFile = "Inputs/day1_input.txt"

func All() {
	fmt.Printf("== Day 1 Answers == \n")
	Part1()
}

func Part1() {
	f, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v: %v\n", inputFile, err))
	}

	scanner := bufio.NewScanner(f)

	maxCalorieCount := 0
	currentCalorieCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			calorieValue, err := strconv.Atoi(scanner.Text())
			if err != nil { 
				panic(fmt.Errorf("strconv.Atoi: %v\n", err))
			 }
			currentCalorieCount += calorieValue
			continue
		}
		maxCalorieCount = max(currentCalorieCount, maxCalorieCount)
		currentCalorieCount = 0
	}
	fmt.Printf("Part 1: \t%d\n", maxCalorieCount)
}

// max returns the maximum of the two inputs
func max(a, b int) int {
	if a >= b { return a }
	return b
}

