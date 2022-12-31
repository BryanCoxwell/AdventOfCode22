package main

import (
	"bufio"
	"fmt"
	"os"
)

const day2InputFile = "../inputs/day2_input.txt"

func main() {
	f := openInputFile(day2InputFile)
	scanner := bufio.NewScanner(f)
	part1Points := 0
	part2Points := 0

	for scanner.Scan() {
		b := scanner.Bytes()
		part1Points += part1(b)
		part2Points += part2(b)
	}

	fmt.Printf("===== Day 2 Answers ===== \n")
	fmt.Printf("Part 1:\t\t%d\n", part1Points)
	fmt.Printf("Part 2:\t\t%d\n", part2Points)
}

// part1Rules calculates the points for a single rock, paper, scissors game
// based on the rules for Part 1
func part1(b []byte) int {
	selectionPoints := asciiToSelectionPoint(b[2])
	score := scoreMap[[2]byte{b[0], b[2]}]
	gameOutcome := selectionPoints + score
	return gameOutcome
}

// part2Rules calculates the points for a single rock, paper, scissors game
// based on the rules for Part 2
func part2(b []byte) int {
	selectionPoints := selectionPointMap[[2]byte{b[0], b[2]}]
	score := asciiToScore(b[2])
	gameOutcome := score + selectionPoints
	return gameOutcome
}

// asciiToSelectionPoint takes an ASCII move value and returns its corresponding selection point value
func asciiToSelectionPoint(move byte) int {
	return int(move) - 87
}

// asciiToScore takes an ASCII outcome value and returns its corresponding score
func asciiToScore(outcome byte) int {
	return (int(outcome) - 88) * 3
}

// selectionPointMap maps {outcome, opponentMove} to a selection point value
var selectionPointMap = map[[2]byte]int{
	{65, 88}: 3,
	{66, 88}: 1,
	{67, 88}: 2,
	{65, 89}: 1,
	{66, 89}: 2,
	{67, 89}: 3,
	{65, 90}: 2,
	{66, 90}: 3,
	{67, 90}: 1,
}

// scoreMap maps {opponentMove, playerMove} to a score value
var scoreMap = map[[2]byte]int{
	{65, 88}: 3,
	{66, 88}: 0,
	{67, 88}: 6,
	{65, 89}: 6,
	{66, 89}: 3,
	{67, 89}: 0,
	{65, 90}: 0,
	{66, 90}: 6,
	{67, 90}: 3,
}

// openInputFile returns the input file
func openInputFile(inp string) *os.File {
	f, err := os.Open(inp)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v: %v", inp, err))
	}
	return f
}
