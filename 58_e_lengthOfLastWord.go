package main

// Given a string s consisting of some words separated by some number of spaces,
// return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.
func lengthOfLastWord(s string) int {
	var length int

	for i := range s {
		if string(s[i]) != " " {
			if i > 0 && string(s[i-1]) == string(" ") {
				length = 0
			}
			length++
		}
	}
	return length
}

// 100/81
