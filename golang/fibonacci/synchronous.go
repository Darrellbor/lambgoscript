package fibonacci

func Synchronous(n int) int {
	if n <= 1 {
		return n
	}
	return Synchronous(n-1) + Synchronous(n-2)
}

func SynchronousWithMemo(n int) int {
	if n <= 1 {
		return n
	}

	// Check if the result is already memoized
	if val, ok := memo[n]; ok {
		return val
	}

	// Calculate the Fibonacci number recursively
	fib := SynchronousWithMemo(n-1) + SynchronousWithMemo(n-2)

	// Memoize the result
	memo[n] = fib

	return fib
}
