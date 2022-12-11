package main

import (
	"fmt"
	"bufio"
)

const day2InputFile = "inputs/day2_input.txt"

func Day2() {
	fmt.Printf("===== Day 2 Answers ===== \n")
	fmt.Printf("Part 1:\t\t%d\n", Day2Part1())
	// fmt.Printf("Part 2:\t\t%d\n", Part2())
}

func Day2Part1() int {
	f := openInputFile(day2InputFile)
	scanner := bufio.NewScanner(f)

	totalPoints := 0
	for scanner.Scan() {
		game := NewGame(scanner.Bytes())
		totalPoints += game.Outcome
	}
	return totalPoints
}

type Game struct {
	MoveA byte
	MoveB byte
	Outcome int
}

func NewGame(l []byte) *Game {
	selectionPoints := int(l[2] - 87)
	score := outcomes[[2]byte{l[2], l[0]}]
	return &Game{
		MoveA: l[0],
		MoveB: l[2],
		Outcome: selectionPoints + score,
	}
}

var outcomes = map[[2]byte] int {
	{88, 65}: 3,
	{88, 66}: 0,
	{88, 67}: 6,
	{89, 65}: 6,
	{89, 66}: 3,
	{89, 67}: 0,
	{90, 65}: 0,
	{90, 66}: 6,
	{90, 67}: 3,
}