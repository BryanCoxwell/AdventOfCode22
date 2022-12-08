package main

import (
	"fmt"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("encountered an error:\n%v", err)
	}
}

func run() error {
	ans, err := Day1()
	if err != nil { return err }
	fmt.Printf("Day 1 answer: %v\n", ans)


	return err
}