package main

import (
	"fmt"
)

func main() {
	haystack := []int{8, 9, 9, 9}

	// str := "Hello World"

	find := plusOne(haystack)

	fmt.Printf("%d", find)
}

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
			}
		}
	}
	return digits
}
