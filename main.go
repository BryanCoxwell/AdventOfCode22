package main

import (
	"fmt"
	"flag"

	"github.com/BryanCoxwell/AdventOfCode22/Day1"
)

var problems = map[int]func() {
	1: Day1.All,
}

func main() {
	dayPtr := flag.Int("day", 1, "day")
	flag.Parse()

	problem, ok := problems[*dayPtr]
	if !ok {
		fmt.Printf("invalid entry: %d\n", *dayPtr)
		fmt.Printf("valid entries are:\n")
		for k := range problems {
			fmt.Printf("[%d]:\tDay %d\n", k, k)
		}
		return
	}

	problem()
	
}
