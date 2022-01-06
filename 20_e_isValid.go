package main

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order

func IsValid(s string) bool {
	// if an odd number, return early
	if len(s)%2 != 0 {
		return false
	}

	stack := []rune{}

	opens := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, r := range s {
		if _, ok := opens[r]; ok {
			if len(stack) == 0 { // if we hit a close with an empty stack
				return false
			}
			end := len(stack) - 1 // top of the stack

			pop := stack[end]

			if opens[r] != pop { // compare the top of the stack to the current rune
				return false
			}

			stack = stack[:end] // trim the item from the stack
			continue

		}

		stack = append(stack, r) // this means we hit an open and so add it to the stack
	}

	return len(stack) == 0 // if there are any dangling items on the stack

}

// 100/89
