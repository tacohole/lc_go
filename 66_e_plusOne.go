package main

// You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer.
// The digits are ordered from most significant to least significant in left-to-right order.
// The large integer does not contain any leading 0's.
// Increment the large integer by one and return the resulting array of digits.

func plusOne(digits []int) []int {
	// find last digit
	end := len(digits) - 1

	// if not nine, just add
	if digits[end] != 9 {
		digits[end] += 1
	} else {
		// as long as we hit nines, flip to 0 and +1 to previous digit
		for i := end; digits[i] == 9; i-- {
			digits[i] = 0
			// what if we hit the start?
			if i == 0 {
				digits = append([]int{1}, digits...)
			} else {
				digits[i-1] += 1
			}
		}
	}
	return digits
}
