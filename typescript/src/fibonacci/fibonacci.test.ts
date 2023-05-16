import fibonacci from "./fibonacci";

test("Test Fibonacci", () => {
  const start = performance.now();
  const testVal = 200;
  const expectedVal = 2.8057117299251016e41;

  const fibonacciVal = fibonacci(testVal);
  const duration = performance.now() - start;

  expect(fibonacciVal).toEqual(expectedVal);
  console.log("Duration of execution:", duration, "ms");
});
