const fibonacci = (n: number, memo: number[] = []): number => {
    if (n <= 1) {
      return n;
    }
  
    if (memo[n]) {
      return memo[n];
    }
  
    const fib = fibonacci(n - 1, memo) + fibonacci(n - 2, memo);
    memo[n] = fib;
  
    return fib;
  };
  
  export default fibonacci;
  