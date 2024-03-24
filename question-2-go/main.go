package main

import "fmt"

func oddSumMaxPair(numbers []int) []int {
	maxPair := make([]int, 2)
	maxPairSum := 0

	for index, value := range numbers {
		for i := index; i < len(numbers); i++ {
			currentSum := value + numbers[i]
			if currentSum % 2 != 0 && (currentSum > maxPairSum || maxPairSum == 0) {
					// fmt.Printf("%d + %d is odd and has a higher sum!\n", value, numbers[i])
					maxPair[0] = value
					maxPair[1] = numbers[i]
					maxPairSum = maxPair[0] + maxPair[1]
			}
		}
	}
	if maxPair[0] == 0 && maxPair[1] == 0 { return nil }
	return maxPair
}

func main() {
	testCase := []int {10, 15, 3, 7, 8, 2, 12, 17}
	result := oddSumMaxPair(testCase)
	fmt.Println(result)
	
	testCase = []int {7, 3, 9, 4, 2, 1, 5}
	result = oddSumMaxPair(testCase)
	fmt.Println(result)
	
	testCase = []int {1, 1, 1, 1, 1}
	result = oddSumMaxPair(testCase)
	fmt.Println(result)
	
	testCase = []int {2, 4, 6, 8, 10}
	result = oddSumMaxPair(testCase)
	fmt.Println(result)
	
	testCase = []int {11, 8, 7, 5, 3, 1}
	result = oddSumMaxPair(testCase)
	fmt.Println(result)
	
	testCase = []int {-5, 3, -7, 2, -1, 4, -6}
	result = oddSumMaxPair(testCase)
	fmt.Println(result)
	
	testCase = []int {0, -1, 1, 1, -2}
	result = oddSumMaxPair(testCase)
	fmt.Println(result)
	
	testCase = []int {-3, -2, -1, -4, -5}
	result = oddSumMaxPair(testCase)
	fmt.Println(result)
	
}