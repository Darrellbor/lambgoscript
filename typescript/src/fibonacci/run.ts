import fs from 'fs';
import Fibonacci from ".";

const runFibonacci = () => {
  const file = fs.readFileSync(`${__dirname}/../../../input/fibonacci.txt`)
  const fileLine = file.toString().split(`\n`)

  const start = performance.now();
  for(let line of fileLine) {
      console.log(`\n`)

      const num = parseInt(line);
      console.log('Number before fibonacci: ', num)
      const fibonacciVal = Fibonacci(num);
      console.log(`Fibonacci of ${num} is ${fibonacciVal}`)
  }
  const duration = performance.now() - start;
  console.log("\n\nDuration of execution:", duration, "ms");
};

export default runFibonacci;
