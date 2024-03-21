package main

import (
	"fmt"
	"sort"
)

func getTwoLargestInteger(numbers []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	reduced := numbers[:2]
	return reduced
}

func main() {
	testCase := []int {7, 3, 7, 4, 2, 1, 5}
	result := getTwoLargestInteger(testCase)
	fmt.Println(result)

	testCase = []int {7, 3, 9, 4, 2, 1, 5}
	result = getTwoLargestInteger(testCase)
	fmt.Println(result)
	
	testCase = []int {1, 2, 3, 4, 5}
	result = getTwoLargestInteger(testCase)
	fmt.Println(result)
	
	testCase = []int {1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	result = getTwoLargestInteger(testCase)
	fmt.Println(result)
	
	testCase = []int {100, 101, 100, 100, 101}
	result = getTwoLargestInteger(testCase)
	fmt.Println(result)
	
	testCase = []int {1, 2, 3, 4, 5, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 5, 4, 3, 2, 1}
	result = getTwoLargestInteger(testCase)
	fmt.Println(result)
	
	testCase = []int {2, 3}
	result = getTwoLargestInteger(testCase)
	fmt.Println(result)
	
	testCase = []int {42, 13, 57, 81, 29, 66, 38, 95, 72, 20, 89, 10, 47, 63, 55, 12, 86, 51, 77,
		32}
	result = getTwoLargestInteger(testCase)
	fmt.Println(result)
	
	testCase = []int {15, 61, 25, 37, 52, 10, 47, 32, 73, 41, 88, 12, 64, 56, 29, 83, 98, 70, 90,
		21}
	result = getTwoLargestInteger(testCase)
	fmt.Println(result)
}
