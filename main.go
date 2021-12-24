package main

import (
	"fmt"
	"strings"
)

func main() {
	//haystack := []int{9, 9, 9, 9}

	a := "10"
	b := "10"

	find := addBinary(a, b)

	fmt.Printf("%s", find)
}

func addBinary(a string, b string) string {
	ansSlice := []string{}

	for i := len(a) - 1; i > 0; i-- {
		for j := len(b) - 1; j > 0; j-- {
			if a[i] == b[j] && string(b[j]) == "0" {
				ansSlice = append([]string{"0"}, ansSlice...)
				fmt.Print("b1")
			}
			if a[i] == b[j] && string(b[j]) == "1" {
				ansSlice = append([]string{"10"}, ansSlice...)
				fmt.Print("b2")
			} else {
				ansSlice = append([]string{"1"}, ansSlice...)
				fmt.Print("b3")
			}
		}
	}

	// back to string
	return strings.Join(ansSlice, "")
}
