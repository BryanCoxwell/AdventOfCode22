package main

import "log"

func main() {
	if err := run(); err != nil {
		log.Fatalf("encountered an error:\n%v", err)
	}
}

func run() error {
	err := Day1()

	return err
}