package main

// Given two binary strings a and b, return their sum as a binary string.

// 0 + 0 = 0
// 0 + 1 = 1
// 1 + 0 = 1
// 1 + 1 =10

func addBinary(a string, b string) string {

	// handle any differences in length
	shortest, longest := len(a), len(b)
	if shortest > longest {
		shortest, longest = longest, shortest
		a, b = b, a
	}

	// byte slice to put answer that's the size of the longest string
	ans := make([]byte, longest+1)

	for i := longest - 1; i >= 0; i-- {
		//
		ans[i+1] += b[i] - '0'
		if i >= longest-shortest {
			ans[i+1] += a[i+shortest-longest] - '0'
		}
		ans[i] = ans[i+1] >> 1
		ans[i+1] = (ans[i+1] & 1) + '0'

	}
	ans[0] = ans[0] + '0'
	if ans[0] == '0' {
		return string(ans[1:])
	} else {
		return string(ans)
	}

}

// 100/80.73
