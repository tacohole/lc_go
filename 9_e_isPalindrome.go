package main

import (
	"strconv"
	"strings"
)

// Given an integer x, return true if x is palindrome integer.
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	intToStr := strconv.Itoa(x)

	strArray := strings.Split(intToStr, "")

	reversed := ""
	for _, char := range strArray {
		reversed = char + reversed
	}

	return reversed == intToStr
}

// 63/10
