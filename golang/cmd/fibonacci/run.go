package main

import (
	"fmt"

	"github.com/Darrellbor/lambgoscript/fibonacci"
)

func main() {
	testFibonacci := 20
	testConcurrentFibonacci := 2000
	testSynchronousWithMemoFibonacci := 20000

	fmt.Println("\nFibonacci to calculate: ", testFibonacci)
	fmt.Printf("Fibonnaci of %d is %v ", testFibonacci, fibonacci.Synchronous(testFibonacci))

	fmt.Println("\n\nFibonacci to calculate: ", testConcurrentFibonacci)
	fmt.Printf("Fibonnaci of %d is %v ", testConcurrentFibonacci, fibonacci.Concurrent(testConcurrentFibonacci))

	fmt.Println("\n\nFibonacci to calculate: ", testSynchronousWithMemoFibonacci)
	fmt.Printf("Fibonnaci of %d is %v ", testSynchronousWithMemoFibonacci, fibonacci.SynchronousWithMemo(testSynchronousWithMemoFibonacci))
}
