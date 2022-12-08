package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)
var inputFile = "files/day1.txt"

func Day1() int {
	f, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v: %v", inputFile, err))
	}

	scanner := bufio.NewScanner(f)

	maxCalorieCount := 0
	currentCalorieCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			calorieValue, err := strconv.Atoi(scanner.Text())
			if err != nil { 
				panic(fmt.Errorf("strconv.Atoi: %v", err))
			 }
			currentCalorieCount += calorieValue
			continue
		}
		maxCalorieCount = max(currentCalorieCount, maxCalorieCount)
		currentCalorieCount = 0
	}
	return maxCalorieCount
}

// Day1WithConcurrency works more or less the same as
// Day1, but employs goroutines for better performance.
// This ends up being more than twice as slow, but it's a good exercise. 
func Day1WithConcurrency() int{
	f, err := os.Open(inputFile)
	if err != nil {
		panic(fmt.Errorf("failed to open file %v: %v", inputFile, err))
	}

	/*
	Create a buffered channel:
	This allows scanIntoChannel to keep working without
	having to wait for items to be received from the channel,
	as sending and receiving to/from channels are blocking operations
	when not buffered.
	*/
	ch := make(chan string, 100)
	go scanIntoChannel(f, ch)

	// Create a WaitGroup so we can block until we're ready to return
	var wg sync.WaitGroup
	wg.Add(1)

	maxCalorieCount := 0

	go func() {
		currentCalorieCount := 0
		for line := range ch{
			if len(line) != 0 {
				calorieValue, err := strconv.Atoi(line)
				if err != nil { 
					panic(fmt.Errorf("strconv.Atoi: %v", err))
				 }
				currentCalorieCount += calorieValue
				continue
			}
			maxCalorieCount = max(maxCalorieCount, currentCalorieCount)
			currentCalorieCount = 0
		}
		wg.Done()
	}()

	wg.Wait()
	return maxCalorieCount
}

// scanIntoChannel pushes text from a file onto a channel
func scanIntoChannel(f *os.File, ch chan string) error {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	close(ch)
	return nil
}

// max returns the maximum of the two inputs
func max(a, b int) int {
	if a >= b { return a }
	return b
}

