package main

import "strings"

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 0; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[0 : len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}

	}

	return prefix
}

// 100, 100
