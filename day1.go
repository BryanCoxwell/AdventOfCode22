package main

import (
	"os"
	"strconv"
	"bufio"
)

func Day1() (int, error) {
	f, err := os.Open("files/day1.txt")
	if err != nil { return -1, err }

	scanner := bufio.NewScanner(f)

	highestCalorieCount := 0
	currentCalorieCount := 0

	for ;scanner.Scan(); {
		line := scanner.Text()
		if len(line) > 0 {
			calorieValue, err := strconv.Atoi(scanner.Text())
			if err != nil { return -1, err}
			currentCalorieCount += calorieValue
			continue
		}
		highestCalorieCount = max(currentCalorieCount, highestCalorieCount)
		currentCalorieCount = 0
	}
	return highestCalorieCount, nil
}

// Return the maximum of the two provided inputs
func max(a, b int) int {
	if a >= b { return a }
	return b
}
