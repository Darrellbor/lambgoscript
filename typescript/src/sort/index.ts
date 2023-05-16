export enum Order {
  Ascending,
  Descending,
}

export const bubbleSort = (arr: number[], order: Order): number[] => {
  const n = arr.length;

  if (order === Order.Ascending) {
    for (let i = 0; i < n - 1; i++) {
      let swapped = false;
      for (let j = 0; j < n - i - 1; j++) {
        if (arr[j] > arr[j + 1]) {
          [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
          swapped = true;
        }
      }
      if (!swapped) {
        break;
      }
    }
  } else if (order === Order.Descending) {
    for (let i = 0; i < n - 1; i++) {
      let swapped = false;
      for (let j = 0; j < n - i - 1; j++) {
        if (arr[j] < arr[j + 1]) {
          [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
          swapped = true;
        }
      }
      if (!swapped) {
        break;
      }
    }
  }

  return arr;
};

export default bubbleSort