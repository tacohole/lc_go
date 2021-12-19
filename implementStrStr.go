package main

import "strings"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	if strings.Contains(haystack, needle) {
		sub := len(needle)
		bale := strings.Split(haystack, "")

		for i, _ := range bale {
			if strings.Join(bale[i:i+sub], "") == needle {
				return i
			}
		}
	}

	return -1
}

// 5.09/5.40
