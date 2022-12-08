package main

import (
	"os"
	"fmt"
)

type Elf struct {
	TotalCalories int
	Calories []int
	RawList [][]byte
}

type Elves []Elf

func Day1() error {
	inputPath := "files/day1.txt"
	data, err := os.ReadFile(inputPath)
	if err != nil {return fmt.Errorf("Day1:\n%v", err)}
	return nil
}


func asciiArrayToInt(line []byte) int {
	return 0
}

func isNewLine(b byte) bool {
	return b == 10
}