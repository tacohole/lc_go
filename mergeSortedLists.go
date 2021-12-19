package main

type ListNode struct {
	Val  int
	Next *ListNode
}

var list1 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  5,
			Next: nil,
		},
	},
}

var list2 = ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	},
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = MergeTwoLists(list1, list2.Next)
	return list2
}

// 9.42/17.01
