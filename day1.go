package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

const day1InputFile = "inputs/day1_input.txt"

func Day1() {
	fmt.Printf("===== Day 1 Answers ===== \n")
	fmt.Printf("Part 1:\t\t%d\n", Day1Part1())
	fmt.Printf("Part 2:\t\t%d\n", Day1Part2())
}

/*
Part1 reads in the values from Inputs/day1_input.txt, sums
the values in each sublist (sublists are separated by empty spaces)
and returns the highest sum.
*/
func Day1Part1() int {
	f := openInputFile(day1InputFile)
	scanner := bufio.NewScanner(f)

	maxCalorieCount := 0
	currentCalorieCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			currentCalorieCount += atoi(line)
			continue
		}
		maxCalorieCount = max(currentCalorieCount, maxCalorieCount)
		currentCalorieCount = 0
	}
	return maxCalorieCount
}

/*
Part2 is similar to Part1, but now we return the sum of the 
sublists that have the top three highest sums.
*/
func Day1Part2() int {
	f := openInputFile(day1InputFile)
	scanner := bufio.NewScanner(f)
	sums := []int{}

	currentCalorieCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			currentCalorieCount += atoi(line)
			continue
		}
		sums = append(sums, currentCalorieCount)
		currentCalorieCount = 0
	}
	sort.Ints(sums)
	maxIdx := len(sums) - 1
	return sums[maxIdx] + sums[maxIdx - 1] + sums[maxIdx - 2]
}

// max returns the maximum of two ints
func max(a, b int) int {
	if a >= b { return a }
	return b
}

// atoi converts an ascii value to an integer
func atoi(inp string) int {
	val, err := strconv.Atoi(inp)
	if err != nil {
		panic(fmt.Errorf("strconv.Atoi: %v", err))
	}
	return val
}

// openInputFile returns the input file
func openInputFile(inp string) *os.File {
	f, err := os.Open(inp)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v: %v", inp, err))
	}
	return f
}