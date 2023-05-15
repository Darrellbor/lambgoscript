package main

import (
	"fmt"

	"github.com/Darrellbor/lambgoscript/sort"
)

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Unsorted array:", arr)
	sort.SynchronousBubbleSort(arr, sort.Descending)
	fmt.Println("Sorted array:", arr)
}
