package main

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same.

func removeDuplicates(nums []int) (int, []int) {

	for i := 0; i < (len(nums) - 1); i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums), nums
}

// 14.87/40.53
