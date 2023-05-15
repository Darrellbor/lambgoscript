package main

import (
	"fmt"

	"github.com/Darrellbor/lambgoscript/fibonacci"
)

func main() {
	testFibonacci := 20

	fmt.Println("Fibonacci to calculate: ", testFibonacci)
	fmt.Printf("Fibonnaci of %d is %v ", testFibonacci, fibonacci.Synchronous(testFibonacci))
}