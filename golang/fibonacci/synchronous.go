package fibonacci

func Synchronous(n int) int {
	if n <= 1 {
		return n
	}
	return Synchronous(n-1) + Synchronous(n-2)
}