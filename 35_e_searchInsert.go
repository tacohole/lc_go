package main

//Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.

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
