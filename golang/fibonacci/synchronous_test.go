package fibonacci

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestSynchronousFibonacci(t *testing.T) {
	cases := []struct{
		name string
		testVal int
		expectedVal int
	} {
		{
			name: "Negative value fibonacci",
			testVal: -20,
			expectedVal: -20,
		},
		{
			name: "Regular small value fibonacci",
			testVal: 20,
			expectedVal: 6765,
		},
		{
			name: "Regular high value fibonacci",
			testVal: 45,
			expectedVal: 1134903170,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			start := time.Now()

			fibonacciVal := Synchronous(tc.testVal)
			duration := time.Since(start)
			require.Equal(t, tc.expectedVal, fibonacciVal, "Computation error as fibonacci value does not equal expected value")
			fmt.Println("Duration of execution: ", duration)
		})
	}
}