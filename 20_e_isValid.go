package main

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order

func IsValid(s string) bool {
	// for open if encounter open before close return false
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
			if len(stack) == 0 {
				return false
			}
			end := len(stack) - 1

			pop := stack[end]

			if opens[r] != pop {
				return false
			}

			stack = stack[:end]
			continue

		}

		stack = append(stack, r)
	}

	return len(stack) == 0

}

// 100/89
