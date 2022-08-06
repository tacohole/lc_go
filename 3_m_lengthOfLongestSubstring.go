package main

func lengthOfLongestSubstring(s string) int {
	// storing the substring we have
	var arr = make([]byte, 0)
	// two pointers and the max length
	var i, j, res = 0, 0, 0
	// while neither pointer is out of range
	for i < len(s) && j < len(s) {
		// if we find the 2nd pointer char in the array
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
			if res < len(arr) {
				// reset the max length
				res = len(arr)
			}
		}
	}
	return res
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
