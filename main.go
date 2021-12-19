package main

import (
	"fmt"
)

func main() {
	haystack := []int{1, 2, 4, 5, 7, 8}

	needle := 6

	find := searchInsert(haystack, needle)

	fmt.Printf("%d", find)
}

func searchInsert(nums []int, target int) int {
	if target == 0 {
		return 0
	}
	mid := len(nums) / 2
	// if middle is greater than target, count down from middle to target and return index
	if nums[mid] > target {
		for i := mid; i > len(nums[:mid]); i-- {
			if nums[i] <= target {
				return i
			}
		}
	}
	// if middle is less than target, count up from middle to target and return index
	if nums[mid] < target {
		for i := mid; i < len(nums[mid:]); i++ {
			if nums[i] >= target {
				return i
			}
		}
	}

	return 0
}
