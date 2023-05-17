import fs from "fs";
import { bubbleSort, Order } from ".";

export const runSort = (): number => {
  const file = fs.readFileSync(`${__dirname}/sort.txt`);
  const fileLine = file.toString().split(`\n`);

  const start = performance.now();
  for (let line of fileLine) {
    console.log(`\n\n`);
    const numStringArr = line.split(" ");
    const numArr = numStringArr.map((num) => parseInt(num));
    console.log("Number array before sorted: ", numArr);
    bubbleSort(numArr, Order.Ascending);
    console.log("Number array after sorted (Ascending): ", numArr);
    bubbleSort(numArr, Order.Descending);
    console.log("Number array after sorted (Descending): ", numArr);
  }
  const duration = performance.now() - start;
  console.log("Duration of execution:", duration, "ms");
  return duration;
};

export const testSort = () => {
  const testLen = 100000;
  let totalMs = 0;

  for (let i = 0; i < testLen; i++) {
    totalMs += runSort();
  }

  const averageMs = totalMs / testLen;
  console.log(
    `The average of ${testLen} tests gives a duration of ${averageMs} ms`
  );
};

export default runSort;
