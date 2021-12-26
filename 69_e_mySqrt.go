package main

// Given a non-negative integer x, compute and return the square root of x.

// Since the return type is an integer, the decimal digits are truncated,
// and only the integer part of the result is returned.

func mySqrt(x int) int {
	// two pointers
	low, high := 0, x

	// gradually converge on sqrt
	for low <= high {
		mid := (low + high) / 2
		mul := mid * mid

		if mul == x {
			// perfect square, return
			return mid
		} else if mul < x {
			// only for x < 2
			low = mid + 1
		} else {
			// we got close but didn't nail it
			if (mid-1)*(mid-1) <= x {
				return mid - 1
			}
			// cut x in half and decrement by 1
			high = mid - 1
		}
	}
	return 0
}

// 100/31.23
