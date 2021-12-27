package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		} else {
			return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		}
	} else {
		return false
	}

}

// 100/24.37
