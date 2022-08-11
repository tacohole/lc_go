package main

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
func LongestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	// the first prefix is the first item in the array
	prefix := strs[0]

	for i := 0; i < len(strs); i++ {
		// while the string at the index doesn't match the existing prefix
		for !hasPrefix(strs[i], prefix) {
			// trim the last char off of the prefix
			prefix = prefix[0 : len(prefix)-1]
			// if we get to no runes in the prefix we're done
			if len(prefix) == 0 {
				return ""
			}
		}

	}

	return prefix
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// 100, 100
