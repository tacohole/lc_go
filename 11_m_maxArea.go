package main

// You are given an integer array height of length n. There are n vertical lines
// drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container
// contains the most water.
// Return the maximum amount of water a container can store.
// Notice that you may not slant the container.

func MaxArea(height []int) int {
	// initialize return value, x[0], x[i]
	max, start, end := 0, 0, len(height)-1

	// while the start index is less than the end index
	for start < end {
		// get the width
		width := end - start

		// initialize the height
		var high int

		// take the lowest of the two y axes and then increment that index
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		// check the area given the current values of x axis and y axis
		temp := width * high
		// if the recent area exceeds the highest recorded, update the var
		if temp > max {
			max = temp
		}
	}

	return max
}

// 77.78, 92.79
