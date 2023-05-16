package fibonacci

import "sync"

var (
	memo = make(map[int]int)
	mu   sync.Mutex
)

func processFibonacci(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Check if the result is already memoized
	mu.Lock()
	if val, ok := memo[n]; ok {
		mu.Unlock()
		ch <- val
		return
	}
	mu.Unlock()

	// Calculate the Fibonacci number recursively
	var fib int
	if n <= 1 {
		fib = n
	} else {
		wg.Add(2)
		ch1 := make(chan int)
		ch2 := make(chan int)

		go processFibonacci(n-1, ch1, wg)
		go processFibonacci(n-2, ch2, wg)

		fib = <-ch1 + <-ch2
	}

	// Memoize the result
	mu.Lock()
	memo[n] = fib
	mu.Unlock()

	ch <- fib
}

func Concurrent(n int) int {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go processFibonacci(n, ch, &wg)

	fib := <-ch
	wg.Wait()

	return fib
}
