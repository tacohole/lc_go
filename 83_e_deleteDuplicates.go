package main

func deleteDuplicates(head *ListNode) *ListNode {
	cursor := head

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}
	// if the current value matches the next value
	for cursor.Next != nil {
		// rewind to the current node
		if cursor.Next.Val == cursor.Val {
			cursor.Next = cursor.Next.Next
		} else {
			cursor = cursor.Next
		}

	}

	return head
}

// 83/100
