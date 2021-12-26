package main

import (
	"fmt"
)

func main() {
	//haystack := []int{9, 9, 9, 9}

	a := 3

	find := climbStairs(a)

	fmt.Printf("%d", find)
}

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}

	previous, current := 1, 1

	for i := 1; i < n; i++ {
		previous, current = current, previous+current
	}

	return current
}
