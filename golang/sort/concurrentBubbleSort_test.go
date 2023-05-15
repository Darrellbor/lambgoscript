package sort

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConcurrentBubbleSort(t *testing.T) {

	cases := []struct{
		name string
		testArr []int
		expectedArr []int
		order int
	}{
		{
			name: "Sort With Negatives",
			testArr: []int{-20, -5, 34, 98, 4, 190, 1309, -900, 45, 76, 0},
			expectedArr: []int{-900, -20, -5, 0, 4, 34, 45, 76, 98, 190, 1309},
			order: Ascending,
		},
		{
			name: "Sort Ascending",
			testArr: []int{20, 5, 34, 98, 4, 190, 1309, 900, 45, 76, 0},
			expectedArr: []int{0, 4, 5, 20, 34, 45, 76, 98, 190, 900, 1309},
			order: Ascending,
		},
		{
			name: "Sort Descending",
			testArr: []int{20, 5, 34, 98, 4, 190, 1309, 900, 45, 76, 0},
			expectedArr: []int{1309, 900, 190, 98, 76, 45, 34, 20, 5, 4, 0},
			order: Descending,
		},
		{
			name: "Sort With Repeated Values",
			testArr: []int{20, 5, 0, 34, 98, 4, 190, 1309, 900, 4,  45, 76, 0},
			expectedArr: []int{1309, 900, 190, 98, 76, 45, 34, 20, 5, 4, 4, 0, 0},
			order: Descending,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()
			ConcurrentBubbleSort(tc.testArr, tc.order)
			duration := time.Since(start)
			require.Equal(t, tc.expectedArr, tc.testArr, "Test array did not match the expected array")
			fmt.Println("Duration of execution: ", duration)
		})
	}
}

