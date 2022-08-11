package main

// Given a string s, find the length of the longest substring without repeating characters.
func LengthOfLongestSubstring(s string) int {
	// storing the substring we have
	var arr = make([]byte, 0)
	// two pointers and the max length
	var i, j, max = 0, 0, 0
	// while neither pointer is out of range
	for i < len(s) && j < len(s) {
		// if we find the char at the 2nd pointer in the array
		if isExist(arr, s[j]) {
			// truncate the array to remove the duplicate
			arr = arr[1:]
			// advance the first pointer
			i++
		} else {
			// add the unique char to the array
			arr = append(arr, s[j])
			// advance the 2nd pointer
			j++
			// if we've exceeded the max length we've seen
			if max < len(arr) {
				// reset the max length
				max = len(arr)
			}
		}
	}
	return max
}

func isExist(arr []byte, key byte) bool {
	for _, value := range arr {
		if value == key {
			return true
		}
	}
	return false
}

// 100, 86
