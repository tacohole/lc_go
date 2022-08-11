package main

// Given an integer x, return true if x is palindrome integer.
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// using a divisor to figure out how many times x goes into 10
	divisor := 1

	// build up the divisor
	for x/divisor >= 10 {
		divisor *= 10
	}

	for x != 0 {
		// checking the first number vs the last number
		if x/divisor != x%10 {
			return false
		}

		x %= divisor
		x /= 10

		divisor /= 100
	}

	return true
}

// 98/61
