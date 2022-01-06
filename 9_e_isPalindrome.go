package main

import (
	"strconv"
	"strings"
)

// Given an integer x, return true if x is palindrome integer.
func IsPalindrome(x int) bool {

	intToStr := strconv.Itoa(x)

	strArray := strings.Split(intToStr, "")
	end := len(strArray) - 1

	for i := 0; i < end; i++ {
		if strArray[i] != strArray[end] {
			return false
		}
		end--
	}

	return true
}

// 63/10
