package main

import "math"

// Given a binary tree, determine if it is height-balanced.
// For this problem, a height-balanced binary tree is defined as:
// a binary tree in which the left and right subtrees of every node differ in height by no more than 1.

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	a := getDepth(root.Left)
	b := getDepth(root.Right)

	if math.Abs(float64(a-b)) <= 1 && isBalanced(root.Right) && isBalanced(root.Left) {
		return true
	}
	return false
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := getDepth(node.Left)
	right := getDepth(node.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

// 11/27
