package main

import "strings"

// Given two binary strings a and b, return their sum as a binary string.

// 0 + 0 = 0
// 0 + 1 = 1
// 1 + 0 = 1
// 1 + 1 =10

func addBinary(a string, b string) string {
	strMap := make(map[byte]byte)
	// add bytes to map

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			strMap[a[i]] = b[j]
		}
	}

	ansSlice := []string{}

	// add according to rules
	for k, v := range strMap {
		if strMap[k] == v && v == 0 {
			ansSlice = append(ansSlice, string(v))
		}
		if strMap[k] == v && v == 1 {
			ansSlice = append(ansSlice, "10")
		} else {
			// either k or v is 0, append 1
			ansSlice = append(ansSlice, "1")
		}
	}
	// back to string
	return strings.Join(ansSlice, "")
}
