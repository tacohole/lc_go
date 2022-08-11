package main

var Roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// Given a roman numeral, convert it to an integer.

func RomanToInt(s string) int {
	var sum int

	for i := 0; i < len(s); i++ {
		// get the rune at the index
		char := string(s[i])

		// convert the rune to an int
		charVal := Roman[char]
		// as long as we're not too close to the end of the slice range
		if i < len(s)-1 {
			// get the rune at the next index
			nextChar := string(s[i+1])
			// convert thar rune to an int
			nextCharVal := Roman[nextChar]
			// check nextchar to see  if it's an IV or IX type situation
			if charVal < nextCharVal {
				// if so, subtract the small from the large and tack it to the sum
				sum += (nextCharVal - charVal)
				i++
			} else {
				// just tack it on
				sum += charVal
			}
			// handle when we've hit the last char
		} else {
			sum += charVal
		}
	}
	return sum
}

// 100/42.41%
