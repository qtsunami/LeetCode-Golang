package main

import (
	"fmt"
	"strconv"
)

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 例如，121 是回文，而 123 不是。
// 示例1 ： 输入：x = 121  输出：true
// 示例2 ： 输入：x = -121  输出：false   解释：从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
// 示例3 ： 输入：x = 10   输出：false   解释：从右向左读, 为 01 。因此它不是一个回文数。

func main() {

	fmt.Println(isPalindrome(1))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	var sx string = strconv.Itoa(x)
	var sl int = len(sx)
	rFlag := true
	for i := 0; i < sl/2; i++ {
		if sx[i] != sx[sl-1-i] {
			rFlag = false
			break
		}
	}
	return rFlag
}
