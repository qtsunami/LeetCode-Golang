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

	fmt.Println(isValid("{[]}"))
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

type stack struct {
	top   int
	value []string
}

func isValid(s string) bool {
	var st *stack = &stack{
		top:   -1,
		value: []string{},
	}
	isValid := true
	for _, value := range s {
		if strings.Contains("([{", string(value)) {
			push(st, string(value))
		} else {
			v := pop(st)
			if v != SymbolMap[string(value)] {
				isValid = false
				break
			}
		}
	}

	return isValid && st.top == -1
}

// push 入栈
func push(st *stack, s string) {
	st.top++
	st.value = append(st.value, s)
}

// pop 出栈
func pop(st *stack) string {
	if st.top == -1 {
		return ""
	}
	v := st.value[st.top]
	st.top--
	if st.top == -1 {
		st.value = []string{}
	} else {
		st.value = st.value[0 : len(st.value)-1]
	}
	return v
}
