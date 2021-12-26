package main

// You are climbing a staircase. It takes n steps to reach the top.
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

func climbStairs(n int) int {

	// 0 or 1 case
	if n <= 1 {
		return n
	}

	previous, current := 1, 1

	for i := 1; i < n; i++ {
		// take a step, add that to the possible paths until there's only 1 step left
		previous, current = current, previous+current
	}

	return current
}

// 100, 64
