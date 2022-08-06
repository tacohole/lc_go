package main

import (
	"fmt"
)

func main() {
	haystack := "baar"
	needle := "r"

	a := StrStr(haystack, needle)

	fmt.Printf("%d", a)
}
