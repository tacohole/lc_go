package main

import (
	"fmt"
)

func main() {
	//haystack := []int{5, 4, -1, 7, 8}

	str := "Hello World"

	find := lengthOfLastWord(str)

	fmt.Printf("%d", find)
}

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
