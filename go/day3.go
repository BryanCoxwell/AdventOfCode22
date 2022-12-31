package main

import (
	"bufio"
	"fmt"
	"os"
)

const day3InputFile = "../inputs/day3_input.txt"

func main() {
	fmt.Printf("===== Day 3 Answers ===== \n")
	fmt.Printf("Part 1:\t\t%d\n", part1())
	fmt.Printf("Part 2:\t\t%d\n", part2())
}

func part1() int {
	f := openInputFile(day3InputFile)
	scanner := bufio.NewScanner(f)

	var prioritySum int
	/*
	   Here we:
	       - Read in a single line of text
	       - Split the line in half to make the two "compartments"
	       - Call findSharedItems on each compartment to get a string containing each item that occurs in both compartments
	       - Calculate the priority value of the shared item and add it to the sum.
	   Note: We can safely calculate the priority value of sharedItems[0] because the problem states that each rucksack
	   has exactly one value that occurs in both of its compartments.
	*/
	for scanner.Scan() {
		line := scanner.Text()
		sharedItems := findSharedItems(line[:len(line)/2], line[len(line)/2:])
		prioritySum += calculatePriority(sharedItems[0])
	}
	return prioritySum
}

func part2() int {
	f := openInputFile(day3InputFile)
	scanner := bufio.NewScanner(f)

	var group []string
	var prioritySum int

	/*
	   Here we:
	       - Read lines in until we have a group of 3 "rucksacks"
	       - Call findBadge on the group to find the one item each rucksack in the group has
	       - Calculate the priority of the badge and add it to prioritySum
	       - Empty the group to start again.
	*/
	for scanner.Scan() {
		group = append(group, scanner.Text())
		if len(group) < 3 {
			continue
		}
		badge := findBadge(group)
		prioritySum += calculatePriority(badge)
		group = []string{}
	}
	return prioritySum
}

// findSharedItems returns a string containing every character that occurs in both a and b
func findSharedItems(a, b string) string {
	var shared string
	for _, i := range a {
		for _, j := range b {
			if i == j {
				shared += string(i)
			}
		}
	}
	return shared
}

// calculatePriority returns the priority value of the given byte:
// bytes corresponding to characters a-z have priorities 1-26
// bytes corresponding to characters A-Z have priorities 27-52
func calculatePriority(b byte) int {
	if b >= 97 {
		return int(b) - 96
	}
	return int(b) - 38
}

// findBadge returns the only value that occurs in each element in a slice of slices
// Each group has 3 slices of bytes, so first we find all values shared between
// group[0] and group[1]. Then we can find which of those values is in group[2].
func findBadge(group []string) byte {
	shared := findSharedItems(group[0], group[1])
	return findSharedItems(shared, group[2])[0]
}

// openInputFile returns the input file
func openInputFile(inp string) *os.File {
	f, err := os.Open(inp)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v: %v", inp, err))
	}
	return f
}
