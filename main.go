package main

import (
	"fmt"
)

func main() {
	haystack := []int{5, 4, -1, 7, 8}

	find := maxSubArray(haystack)

	fmt.Printf("%d", find)
}
