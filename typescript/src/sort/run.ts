import { bubbleSort, Order } from './sort'

const arr = [5, 3, 8, 2, 1, 4];
console.log("Original array:", arr);

bubbleSort(arr, Order.Ascending);
console.log("Sorted array in ascending order:", arr);

bubbleSort(arr, Order.Descending);
console.log("Sorted array in descending order:", arr);