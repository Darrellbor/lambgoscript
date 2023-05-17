import fs from "fs";
import Fibonacci from ".";

const runFibonacci = (): number => {
  const file = fs.readFileSync(`${__dirname}/fibonacci.txt`);
  const fileLine = file.toString().split(`\n`);

  const start = performance.now();
  for (let line of fileLine) {
    console.log(`\n`);

    const num = parseInt(line);
    console.log("Number before fibonacci: ", num);
    const fibonacciVal = Fibonacci(num);
    console.log(`Fibonacci of ${num} is ${fibonacciVal}`);
  }
  const duration = performance.now() - start;
  console.log("\n\nDuration of execution:", duration, "ms");

  return duration;
};

export const testFibonacci = () => {
  const testLen = 100000;
  let totalMs = 0;

  for (let i = 0; i < testLen; i++) {
    totalMs += runFibonacci();
  }

  const averageMs = totalMs / testLen;
  console.log(
    `The average of ${testLen} tests gives a duration of ${averageMs} ms`
  );
};

export default runFibonacci;
