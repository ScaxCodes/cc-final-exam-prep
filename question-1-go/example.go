package main

import "fmt"

func largestInteger(numbers []int) int {
	if len(numbers) == 0 {
return 0 }
	largest := numbers[0]
	for _, num := range numbers {
				 if num > largest {
							 largest = num
} }
	return largest
}

func main() {
	testCase1 := []int{-7, -3, -9, -4, -2, -1, -2, -3}
	answer1 := -1
	result := largestInteger(testCase1)
	testCase1Correct := result == answer1
	fmt.Println(testCase1Correct)
}
