package main

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

func isSymmetric(root *TreeNode) bool {
	if root == nil { // easy
		return true
	}
	return isSymmetricLR(root.Left, root.Right)
}

func isSymmetricLR(left, right *TreeNode) bool {
	if left == nil && right == nil { // also easy
		return true
	}
	if (left == nil) != (right == nil) {
		return false
	}
	return (left.Val == right.Val) && // first check the current values
		isSymmetricLR(left.Left, right.Right) && // then move to the outer branches
		isSymmetricLR(left.Right, right.Left) // and also check the inner branches
}

// 100/50.57
// describe symmetry - rootA.Left.Value == rootB.Right.Value && rootA.Right.Value == rootB.Left.Value
