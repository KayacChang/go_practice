package main

import (
	"fmt"
	"math"
)

/**
   Find the smallest Subarray with given sum

   Example Input:
	[4, 2, 2, 7, 8, 1, 2, 8, 1, 0]
*/

func main() {
	input := []int{4, 2, 2, 7, 8, 1, 2, 8, 1, 0}

	sum := 8

	result := findSmallestSubarray(input, sum)

	fmt.Println(result)
}

func findSmallestSubarray(arr []int, targetSum int) int {

	min := func(x, y int) int {

		return int(math.Min(float64(x), float64(y)))
	}

	res := math.MaxInt64

	sum := 0
	start := 0

	for end := 0; end < len(arr); end++ {

		sum += arr[end]

		for sum >= targetSum {

			res = min(res, end-start+1)

			sum -= arr[start]

			start += 1
		}
	}

	return res
}
