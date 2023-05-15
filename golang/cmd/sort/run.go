package main

import (
	"fmt"

	"github.com/Darrellbor/lambgoscript/sort"
)

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	arr2 := []int{84, 23, 98, 25, 9, 12, 54}
	fmt.Println("\nUnsorted array:", arr)
	sort.SynchronousBubbleSort(arr, sort.Descending)
	fmt.Println("Sorted array (with synchronous bubble sort):", arr)

	fmt.Println("\n\nUnsorted array 2:", arr2)
	sort.ConcurrentBubbleSort(arr2, sort.Ascending)
	fmt.Println("Sorted array (with concurrent bubble sort):", arr2)
}
