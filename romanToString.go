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

func romanToInt(s string) int {
	sum := 0

	for i := 0; i < len(s); i++ {
		char := string(s[i])
		charVal := Roman[char]
		if i < len(s)-1 {
			nextChar := string(s[i+1])
			nextCharVal := Roman[nextChar]
			if charVal < nextCharVal {
				sum += (nextCharVal - charVal)
				i++
			} else {
				sum += charVal
			}
		} else {
			sum += charVal
		}
	}
	return sum
}

// 34%
