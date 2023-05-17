import { Handler } from "aws-lambda";
import runSort, { testSort } from "./sort/run";
import runFibonacci, { testFibonacci } from "./fibonacci/run";

// runSort();
// runFibonacci();

// testSort()
// testFibonacci()

export const sortHandler: Handler = async (event, context) => {
  testSort();
};

export const fibonacciHandler: Handler = async (event, context) => {
  testFibonacci();
};
