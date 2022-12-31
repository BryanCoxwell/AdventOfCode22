package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const day5InputFile = "../inputs/day5_input.txt"

const SPACE = byte(32)
const STACK_COUNT = 9
const STACK_WIDTH = 4

/*
To solve the puzzle we start by reading in the starting arrangement from the
input file. Since the stacks of crates are essentially last in/first out
queues, we can represent them by defining a Stack type which has push and
pop methods. An unshift method is also needed for building the original
arrangement. The difference between part 1 and part 2 is that multiple
crates can be moved at once in part 2, so we can copy the starting arrangement
and then as we iterate over the instructions simply pass the part1 stacks
to the move() function and pass the part2 stacks to the moveMultiple() function.
After all instructions are executed, we can check what the top crates are by
calling pop on each stack.
*/
func main() {
	f := openInputFile(day5InputFile)
	scanner := bufio.NewScanner(f)
	part1Stacks := readStartingArrangement(scanner)
	part2Stacks := make([]Stack, len(part1Stacks))
	copy(part2Stacks, part1Stacks)

	for scanner.Scan() {
		count, from, to := readInstruction(scanner.Text())
		move(&part1Stacks[from-1], &part1Stacks[to-1], count)
		moveMultiple(&part2Stacks[from-1], &part2Stacks[to-1], count)
	}

	fmt.Printf("===== Day 5 Answers ===== \n")
	fmt.Printf("Part 1:\t\t%s\n", topCrates(part1Stacks))
	fmt.Printf("Part 2:\t\t%s\n", topCrates(part2Stacks))
}

// Stack represents a last in first out queue of integer values
type Stack []int

// push appends an integer to the stack
func (s *Stack) push(item int) {
	*s = append(*s, item)
}

// pop returns the integer on top of the stack
// returns -1 if the stack is empty
func (s *Stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

// unshift adds an integer to the beginning of the stack
func (s *Stack) unshift(item int) {
	*s = append([]int{item}, *s...)
}

// move pops `count` items from the fromStack and pushes them to the toStack
func move(fromStack, toStack *Stack, count int) {
	for i := 0; i < count; i++ {
		toStack.push(fromStack.pop())
	}
}

// moveMultiple is similar to move but can move more than one crate at once
func moveMultiple(fromStack, toStack *Stack, count int) {
	buf := &Stack{}
	for i := 0; i < count; i++ {
		buf.push(fromStack.pop())
	}
	for _, _ = range *buf {
		toStack.push(buf.pop())
	}
}

// readStartingArrangement parses the starting arrangement
// from the input text and returns the arrangment represented
// by a slice of Stacks
func readStartingArrangement(s *bufio.Scanner) []Stack {
	stacks := make([]Stack, STACK_COUNT)
	for s.Scan() {
		line := s.Bytes()
		if len(line) == 0 {
			break
		}
		for i := 1; i <= len(line); i += STACK_WIDTH {
			if line[i] == SPACE {
				continue
			}
			stacks[i/STACK_WIDTH].unshift(int(line[i]))
		}
	}
	return stacks
}

// readInstruction takes in a line of bytes and returns three
// values: the number of crates to move, the source stack, and the destination stack
func readInstruction(line string) (int, int, int) {
	lineSplit := strings.Fields(line)
	return atoi(lineSplit[1]), atoi(lineSplit[3]), atoi(lineSplit[5])
}

// topCrates returns a string representing the crates on top of each stack
func topCrates(stacks []Stack) string {
	tops := ""
	for _, s := range stacks {
		tops += string(s.pop())
	}
	return tops
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
