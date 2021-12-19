package main

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
