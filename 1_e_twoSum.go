package main

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
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

// 96/29
