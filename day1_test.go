package main

import (
	"testing"
	"fmt"
)

var _, _ = fmt.Println("Testing Day1")

func BenchmarkDay1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day1()
	}
}

func BenchmarkDay1WithConcurrency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day1WithConcurrency()
	}
}