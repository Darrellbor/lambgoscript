package sort

const (
	Ascending  = iota
	Descending = iota
)

func SynchronousBubbleSort(arr []int, order int) {
	n := len(arr)

	if order == Ascending {
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
	} else if order == Descending {
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
	}
}
