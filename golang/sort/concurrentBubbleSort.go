package sort

import "sync"

func bubbleSortingAscending(arr []int, wg *sync.WaitGroup) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	wg.Done()
}

func bubbleSortingDescending(arr []int, wg *sync.WaitGroup) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	wg.Done()
}

func ConcurrentBubbleSort(arr []int, order int) {
	n := len(arr)
	var wg sync.WaitGroup
	wg.Add(n - 1)

	for i := 0; i < n-1; i++ {
		if order == Ascending {
			go bubbleSortingAscending(arr[i:], &wg)
		} else if order == Descending {
			go bubbleSortingDescending(arr[i:], &wg)
		}
	}

	wg.Wait()
}
