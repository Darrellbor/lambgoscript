package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Darrellbor/lambgoscript/sort"
)

func main() {
	testLen := 1000
	var totalMs time.Duration

	for i := 0; i < testLen; i++ {
		totalMs += sortMain()
	}

	averageMs := totalMs / time.Duration(testLen)
	fmt.Printf("The average of %d tests gives a duration of %v", testLen, averageMs)
}

func sortMain() time.Duration {
	filePath := "../input/sort.txt"
	file, err := os.ReadFile(filePath)
	ErrCheck(err)

	fileContent := string(file)
	fileLines := strings.Split(fileContent, "\n")

	start := time.Now()
	for _, line := range fileLines {
		fmt.Printf("\n\n")
		numStringArr := strings.Split(line, " ")
		numArr := make([]int, 0, len(numStringArr))

		for _, numString := range numStringArr {
			num, err := strconv.Atoi(numString)
			if err != nil {
				fmt.Println("Error converting number:", err)
				continue
			}
			numArr = append(numArr, num)
		}
		synchronousCall(numArr)
	}
	duration := time.Since(start)
	fmt.Println("Duration of execution:", duration)

	return duration
}

func synchronousCall(numArr []int) {
	fmt.Println("Number array before sorted:", numArr)
	sort.SynchronousBubbleSort(numArr, sort.Ascending)
	fmt.Println("Number array after sorted (Ascending):", numArr)
	sort.SynchronousBubbleSort(numArr, sort.Descending)
	fmt.Println("Number array after sorted (Descending):", numArr)
}

func concurrentCall(numArr []int) {
	fmt.Println("Number array before sorted:", numArr)
	sort.ConcurrentBubbleSort(numArr, sort.Ascending)
	fmt.Println("Number array after sorted (Ascending-Concurrent):", numArr)
	sort.ConcurrentBubbleSort(numArr, sort.Descending)
	fmt.Println("Number array after sorted (Descending-Concurrent):", numArr)
}

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}
