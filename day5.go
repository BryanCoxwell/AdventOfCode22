package main

import (
	"fmt"
	//"bufio"
	"os"
)

const day5InputFile = "inputs/day5_input.txt"

func main() {
	fmt.Printf("===== Day 5 Answers ===== \n")
	fmt.Printf("Part 1:\t\t%d\n", part1())
	// fmt.Printf("Part 2:\t\t%d\n", Day5Part2())
}

type Stack []int

func (s Stack) push(item int) {
	s = append(s, item)
}

func (s *Stack) pop(count int) []int {
	if count > len(*s) {
		fmt.Println("Can't pop more items than there are on the stack")
		return nil
	}
	popped := (*s)[len(*s)-count:]
	*s = (*s)[:len(*s)-count]
	return popped
}

func part1() int {
	s := Stack{1, 2, 5, 1}
	a := s.pop(1)
	fmt.Println(a)
	b := s.pop(1)
	fmt.Println(b)
	fmt.Println(s)

	return s[0]
}

// openInputFile returns the input file
func openInputFile(inp string) *os.File {
	f, err := os.Open(inp)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v: %v", inp, err))
	}
	return f
}
