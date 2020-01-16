package main

import (
	"fmt"
	"math"
)

/**
   Find the max sum subarray of a fixed size k

   Example Input:
	[4, 2, 1, 7, 8, 1, 2, 8, 1, 0]

   Reference:
   	Sliding Window Technique
*/

func main() {
	input := []int{4, 2, 1, 7, 8, 1, 2, 8, 1, 0}

	k := 3

	result := findMaxSum(input, k)

	fmt.Println(result)
}

func findMaxSum(arr []int, k int) int {

	max := func(x, y int) int {

		return int(math.Max(float64(x), float64(y)))
	}

	res := math.MinInt64
	sum := 0

	for i := 0; i < len(arr); i++ {

		sum += arr[i]

		if i >= k-1 {
			res = max(res, sum)

			sum -= arr[i-(k-1)]
		}
	}

	return res
}
