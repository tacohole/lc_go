package main

import "math"

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// make a slice to compare sums of subarrays
	sums := make([]int, len(nums))

	sums[0] = nums[0]

	// for nums[i], keep a running sum and if it decrements, keep the larger sum
	for i := 1; i < len(nums); i++ {
		sums[i] = max(nums[i]+sums[i-1], nums[i])
	}

	maxVal := math.MinInt32

	// for the sums of subarrays, find the largest one
	for i := 0; i < len(sums); i++ {
		maxVal = max(sums[i], maxVal)
	}
	return maxVal
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

// 91/17
