package main

func maxDepth(root *TreeNode) int {
	max := 0
	var traversal func(l int, node *TreeNode)
	traversal = func(l int, node *TreeNode) {
		if node == nil {
			if l > max {
				max = l
			}
			return
		}
		l++
		traversal(l, node.Left)
		traversal(l, node.Right)
	}
	traversal(0, root)
	return max
}

// 86/51
