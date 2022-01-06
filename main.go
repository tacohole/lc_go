package main

import (
	"fmt"
)

func main() {
	haystack := "baar"
	needle := "r"

	a := strStr(haystack, needle)

	fmt.Printf("%d", a)
}

func strStr(haystack string, needle string) int {
	sub := len(needle)
	if sub == 0 {
		return 0
	} else if sub == 1 {
		for i := 0; i < len(haystack); i++ {
			if string(haystack[i]) == needle {
				return i
			}
		}
	} else if haystack == needle {
		return 0
	} else {
		for i := 0; i < len(haystack); i++ {
			if haystack[i:i+sub] == needle {
				return i
			}
		}
	}

	return -1
}
