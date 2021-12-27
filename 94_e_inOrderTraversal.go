package main

// Given the root of a binary tree, return the inorder traversal of its nodes' values.

// Until all nodes are traversed −
// Step 1 − Recursively traverse left subtree.
// Step 2 − Visit root node.
// Step 3 − Recursively traverse right subtree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	vals := []int{}
	inorder(root, &vals)
	return vals
}

func inorder(root *TreeNode, vals *[]int) {
	if root != nil {
		inorder(root.Left, vals)        // traverse until nil
		*vals = append(*vals, root.Val) // append each val
		inorder(root.Right, vals)       // do the same thing for right nodes
	}
}

// 100, 100
