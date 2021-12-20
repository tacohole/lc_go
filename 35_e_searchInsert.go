package main

//Given a sorted array of distinct integers and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
func searchInsert(nums []int, target int) int {

	// find midpoint
	mid := len(nums) / 2

	// if the number at the midpoint is greater than target, count down from middle to target and return index
	if nums[mid] >= target {
		for i := mid; i > -1; i-- {
			// if we hit or pass up the target return
			if nums[i] == target {
				return i
			}
			if nums[i] < target {
				return i + 1
			}
		}
		return 0
	}
	// if middle is less than target, count up from middle to target and return index
	if nums[mid] < target {
		for i := mid; i < len(nums); i++ {
			if nums[i] >= target {
				return i
			}
		}
	}
	// if we get here our target is the end of the array
	return len(nums)
}

// 80.75/35.72
