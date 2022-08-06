package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
func TwoSum(nums []int, target int) ([]int, error) {
	if nums == nil || target < 1 {
		return nil, fmt.Errorf("input values cannot be empty")
	}
	seen := map[int]int{}

	for i, num := range nums {
		match := target - num
		if j, found := seen[match]; found {
			return []int{i, j}, nil
		}
		seen[num] = i
	}

	return nil, fmt.Errorf("no values in array sum to target")
}

// 96/29
