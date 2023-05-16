import { bubbleSort, Order } from ".";

describe("SynchronousBubbleSort", () => {
  test("Sort With Negatives", () => {
    const testArr = [-20, -5, 34, 98, 4, 190, 1309, -900, 45, 76, 0];
    const expectedArr = [-900, -20, -5, 0, 4, 34, 45, 76, 98, 190, 1309];

    const start = performance.now();
    bubbleSort(testArr, Order.Ascending);
    const duration = performance.now() - start;

    expect(testArr).toEqual(expectedArr);
    console.log("Duration of execution:", duration, "ms");
  });

  test("Sort Ascending", () => {
    const testArr = [20, 5, 34, 98, 4, 190, 1309, 900, 45, 76, 0];
    const expectedArr = [0, 4, 5, 20, 34, 45, 76, 98, 190, 900, 1309];

    const start = performance.now();
    bubbleSort(testArr, Order.Ascending);
    const duration = performance.now() - start;

    expect(testArr).toEqual(expectedArr);
    console.log("Duration of execution:", duration, "ms");
  });

  test("Sort Descending", () => {
    const testArr = [20, 5, 34, 98, 4, 190, 1309, 900, 45, 76, 0];
    const expectedArr = [1309, 900, 190, 98, 76, 45, 34, 20, 5, 4, 0];

    const start = performance.now();
    bubbleSort(testArr, Order.Descending);
    const duration = performance.now() - start;

    expect(testArr).toEqual(expectedArr);
    console.log("Duration of execution:", duration, "ms");
  });

  test("Sort With Repeated Values", () => {
    const testArr = [20, 5, 0, 34, 98, 4, 190, 1309, 900, 4, 45, 76, 0];
    const expectedArr = [1309, 900, 190, 98, 76, 45, 34, 20, 5, 4, 4, 0, 0];

    const start = performance.now();
    bubbleSort(testArr, Order.Descending);
    const duration = performance.now() - start;

    expect(testArr).toEqual(expectedArr);
    console.log("Duration of execution:", duration, "ms");
  });
});
