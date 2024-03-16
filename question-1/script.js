function getTwoLargestInt(array) {
  const sortedArray = array.sort((a, b) => b - a);
  return sortedArray.slice(0, 2);
}

console.log(getTwoLargestInt([7, -2, 3, 4, 5, 1, 23, 31, 23]));
console.log(getTwoLargestInt([2, 1]));
console.log(getTwoLargestInt([100, 200, 200, 100]));
console.log(getTwoLargestInt([-1, -2, -3, -4, -5]));
console.log(getTwoLargestInt([5, 7, 9, 1, 3]));
console.log(getTwoLargestInt([0, 0, 0, 0, 0]));
console.log(
  getTwoLargestInt([
    5, 7, 9, 1, 3, 11, 15, 12, 18, 23, 44, 32, 49, 50, 31, 60, 70, 80, 90, 100,
    101, 102, 200, 500, 250,
  ])
);
console.log(
  getTwoLargestInt([
    100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400,
    1500, 1600, 1700, 1800, 1900, 2000, 2100, 2200, 2300, 2400, 2500,
  ])
);
