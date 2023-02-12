package main

import (
	"fmt"
	"strings"
)

// https://leetcode.cn/problems/valid-parentheses/
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//有效字符串需满足：
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//每个右括号都有一个对应的相同类型的左括号。

//示例 1：   输入：s = "()"       输出：true
//示例 2：   输入：s = "()[]{}"   输出：true
//示例 3：   输入：s = "(]"       输出：false

func main() {

	s := "(){}[]"

	startStr := "({["

	stackArray := []string{}

	for _, v := range s {
		if strings.Contains(startStr, string(v)) {
			stackArray = append(stackArray, string(v))
		} else {

		}
	}

	//fmt.Println(isValid("(){}[]"))
}

const (
	LeftBracket        = "("
	RightBracket       = ")"
	LeftMiddleBracket  = "["
	RightMiddleBracket = "]"
	LeftCurlyBracket   = "{"
	RightCurlyBracket  = "}"
)

var SymbolMap = map[string]string{
	RightBracket:       LeftBracket,
	RightMiddleBracket: LeftMiddleBracket,
	RightCurlyBracket:  LeftCurlyBracket,
}

func isValid(s string) bool {
	for _, value := range s {
		fmt.Println(string(value))
		fmt.Println(strings.Contains(s, string(value)))
	}

	return false
}

func PushStack() {

}

func OutStack() {

}
