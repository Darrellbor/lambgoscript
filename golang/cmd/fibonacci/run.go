package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Darrellbor/lambgoscript/fibonacci"
)

func main() {
	testLen := 1000
	var totalMs time.Duration

	for i := 0; i < testLen; i++ {
		totalMs += fibonacciMain()
	}

	averageMs := totalMs / time.Duration(testLen)
	fmt.Printf("The average of %d tests gives a duration of %v", testLen, averageMs)
}

func fibonacciMain() time.Duration {
	filePath := "../input/fibonacci.txt"
	file, err := os.ReadFile(filePath)
	ErrCheck(err)

	fileContent := string(file)
	fileLines := strings.Split(fileContent, "\n")

	start := time.Now()
	for _, line := range fileLines {
		fmt.Printf("\n\n")

		num, err := strconv.Atoi(line)
		ErrCheck(err)
		synchronousCallWithMemo(num)
	}
	duration := time.Since(start)
	fmt.Println("\n\nDuration of execution:", duration)

	return duration
}

func synchronousCall(num int) {
	fmt.Println("Number before convertion:", num)
	fibVal := fibonacci.Synchronous(num)
	fmt.Printf("Fibonacci for number %d is %v:", num, fibVal)
}

func concurrentCall(num int) {
	fmt.Println("Number before convertion:", num)
	fibVal := fibonacci.Concurrent(num)
	fmt.Printf("Fibonacci for number %d is %v:", num, fibVal)
}

func synchronousCallWithMemo(num int) {
	fmt.Println("Number before convertion:", num)
	fibVal := fibonacci.SynchronousWithMemo(num)
	fmt.Printf("Fibonacci for number %d is %v:", num, fibVal)
}

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}
