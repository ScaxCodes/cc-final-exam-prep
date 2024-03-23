package main

import "fmt"

func oddSumMaxPair(numbers []int) []int {
	maxPair := make([]int, 2)
	maxPairSum := 0

	for _, value := range numbers {
		for i := value; i < len(numbers); i++ {
			currentSum := value + numbers[i]
			if currentSum % 2 != 0 && currentSum > maxPairSum {
				fmt.Printf("%d + %d is odd and has a higher sum!\n", value, numbers[i])
				maxPair[0] = value
				maxPair[1] = numbers[i]
				maxPairSum = maxPair[0] + maxPair[1]
			}
		}
	}

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

/* PSEUDO CODE
loop first element
	1 + 2, if even return
		if odd && > largestSum -> store both element in slice
	1 + 3, if even return
		(...)
	(...)
	stop when slice.length reached
loop second element
	2 + 3
	2 + 4
	(...) like above
return largestSum
*/