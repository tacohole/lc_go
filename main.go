package main

import (
	"fmt"
)

func main() {
	haystack := []int{1, 2, 4, 5, 7, 8}

	needle := 6

	find := twoSum(haystack, needle)

	fmt.Printf("%d", find)
}

func twoSum(nums []int, target int) []int {

	seen := map[int]int{}

	for i, num := range nums {
		match := target - num
		if j, found := seen[match]; found {
			return []int{i, j}
		}
		seen[num] = i
	}

	return []int{}
}
