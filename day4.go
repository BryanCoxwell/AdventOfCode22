package main

import (
	"fmt"
	"bufio"
	"strings"
)

const day4InputFile = "inputs/day4_input.txt"

func Day4() {
	fmt.Printf("===== Day 3 Answers ===== \n")
	fmt.Printf("Part 1:\t\t%d\n", Day4Part1())
	fmt.Printf("Part 2:\t\t%d\n", Day4Part2())
}

func Day4Part1() int {
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

func Day4Part2() int {
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