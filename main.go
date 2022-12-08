package main

import (
	"fmt"
	// "github.com/pkg/profile"
)

func main() {
	day1Ans := Day1()
	day1ConcurrencyAns := Day1WithConcurrency()
	fmt.Printf("Day 1 Answer:\t\t%v\n", day1Ans)
	fmt.Printf("Day 1 (v2) Answer:\t%v\n", day1ConcurrencyAns)
}
