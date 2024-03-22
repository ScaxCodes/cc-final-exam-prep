package main

import "fmt"

func oddSumMaxPair(numbers []int) []int {
	reduced := numbers[:2]
	return reduced
}

func main() {
	testCase := []int {7, 3, 9, 4, 2, 1, 5}
	result := oddSumMaxPair(testCase)
	fmt.Println(result)
}