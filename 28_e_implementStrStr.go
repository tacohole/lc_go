package main

import "strings"

// Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if strings.Contains(haystack, needle) {
		sub := len(needle)
		bale := strings.Split(haystack, "")

		for i := range bale {
			if strings.Join(bale[i:i+sub], "") == needle {
				return i
			}
		}
	}

	return -1
}

// 5.09/5.40
