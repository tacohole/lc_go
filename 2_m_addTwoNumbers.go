package main

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// we are going to sum the values in place in l1
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// we need to keep track of the remainder in case the sum
	// of two node values is >= 10
	remainder, head := 0, l1

	for {
		// add the two values and whatever remainder we happen to have
		l1.Val += l2.Val + remainder

		// pull the new remainder
		remainder = int(l1.Val / 10)
		l1.Val = l1.Val % 10

		// if we're done with l2 then break
		if l2.Next == nil {
			break
			// if we're done with l1 then pull over Next from l2
		} else if l1.Next == nil {
			l1.Next = l2.Next
			break
		}

		// otherwise we're continuing to the next node
		l1 = l1.Next
		l2 = l2.Next

	}

	// handling values in remainder if there are any
	for remainder != 0 {
		// if we have remainder we need somewhere to put it
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		l1.Next.Val += remainder

		// checking once again for overflow
		remainder = l1.Next.Val / 10
		l1.Next.Val = l1.Next.Val % 10

		// continuing as long as there's still a value in remainder
		l1 = l1.Next
	}

	// return the pointer to the head of l1
	return head
}

// 84/95
