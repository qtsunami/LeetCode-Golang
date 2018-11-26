package main

import (
	"fmt"
)

func main() {
	s := "hello world."

	s = reverseString(s)

	fmt.Println(s)
}

func reverseString(s string) string {
	var r string
	for i := len(s) - 1; i >= 0; i -- {
		r = r + fmt.Sprintf("%c", s[i])
	}
	return r
}
