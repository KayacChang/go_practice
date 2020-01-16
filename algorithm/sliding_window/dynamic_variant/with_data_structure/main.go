package main

import (
	"fmt"
	"math"
)

/**
  Find the longest substring length with k distinct characters

  Example Input:
  	AAAHHIBC
*/

func main() {
	input := "AAAHHIBC"

	k := 3

	result := findLongestSubstring(input, k)

	fmt.Println(result)
}

func findLongestSubstring(input string, k int) int {

	max := func(x, y int) int {
		return int(math.Max(float64(x), float64(y)))
	}

	start := 0

	charMap := map[string]int{}

	res := 0

	for end := 0; end < len(input); end++ {

		char := string(input[end])

		charMap[char] += 1

		for len(charMap) >= k {

			char := string(input[start])

			charMap[char] -= 1

			if charMap[char] == 0 {
				delete(charMap, char)
			}

			start += 1
		}

		res = max(res, end-start+1)
	}

	return res
}
