package main

import (
	"fmt"
    "os"
    "strconv"
	"bufio"
	"strings"
)

const day4InputFile = "inputs/day4_input.txt"

func main() {
	fmt.Printf("===== Day 4 Answers ===== \n")
	fmt.Printf("Part 1:\t\t%d\n", part1())
	fmt.Printf("Part 2:\t\t%d\n", part2())
}

func part1() int {
	f := openInputFile(day4InputFile)
    scanner := bufio.NewScanner(f)

	var containedRanges int = 0

	for scanner.Scan() {
		areas := convertAreaStringToIntSlice(scanner.Text())
		if (areas[0] >= areas[2] && areas[1] <= areas[3]) || (areas[2] >= areas[0] && areas[3] <= areas[1]) {
			containedRanges += 1
		}
	}
	return containedRanges
}

func part2() int {
	f := openInputFile(day4InputFile)
    scanner := bufio.NewScanner(f)

	var overlappingRanges int = 0

	for scanner.Scan() {
		areas := convertAreaStringToIntSlice(scanner.Text())
		if (areas[1] >= areas[2] && areas[1] <= areas[3]) || (areas[3] >= areas[0] && areas[3] <= areas[1]) {
			overlappingRanges += 1
		}
	}
	return overlappingRanges
}

// convertAreaStringToIntSlice takes a string of the form "a-b,c-d" (where a, b, c, and d can be converted 
// to integers) and returns a slice of integers [a b c d]
func convertAreaStringToIntSlice(lst string) []int {
	pair := strings.Split(lst, ",")
	areaStringSlice := append(strings.Split(pair[0], "-"), strings.Split(pair[1], "-")...)
	areaIntSlice := []int{}
	for _, i := range areaStringSlice {
		areaIntSlice = append(areaIntSlice, atoi(i))
	}
	return areaIntSlice
}

// openInputFile returns the input file
func openInputFile(inp string) *os.File {
	f, err := os.Open(inp)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v: %v", inp, err))
	}
	return f
}

// atoi converts an ascii value to an integer
func atoi(inp string) int {
	val, err := strconv.Atoi(inp)
	if err != nil {
		panic(fmt.Errorf("strconv.Atoi: %v", err))
	}
	return val
}
