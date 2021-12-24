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
		for i := end; i >= 0; i-- {
			// what if we hit a 9 at the start?
			if i == 0 && digits[i] == 9 {
				digits = append([]int{1}, digits...)
				digits[i+1] = 0
			} else if digits[i-1] == 9 {
				// if the next digit is also a 9, flip to 0 and continue the loop
				digits[i] = 0
				continue
			} else {
				// we hit the last 9
				digits[i] = 0
				digits[i-1] += 1
				break
			}
		}
	}
	return digits
}

// 100/100
