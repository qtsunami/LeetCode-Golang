package main

import (
	"fmt"
)

// 58. 最后一个单词的长度
// https://leetcode.cn/problems/length-of-last-word/
// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。
// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

// 示例 1： 输入：s = "Hello World"   输出：5     解释：最后一个单词是“World”，长度为5。
// 示例 2： 输入：s = "   fly me   to   the moon  "    输出：4    解释：最后一个单词是“moon”，长度为4。
// 示例 3： 输入：s = "luffy is still joyboy"       输出：6   解释：最后一个单词是长度为6的“joyboy”。

// 提示：
// 1 <= s.length <= 104
// s 仅有英文字母和空格 ' ' 组成
// s 中至少存在一个单词

func main() {
	s := "   fly me   to   the moon"
	fmt.Println(lengthOfLastWord(s))
}

// 解法一： 循环字符串，从第一个不为空的字符开始，到再次遇到为空的字符结束
func lengthOfLastWord(s string) int {
	wl := 0
	start, end := false, false
	i := len(s) - 1
	for i >= 0 {
		if string(s[i]) != " " {
			start = true
		}
		if start && string(s[i]) == " " {
			end = true
		}
		if start && !end {
			wl++
		}
		i--
	}
	return wl
}

// 解法二： 将字符串根据" " 分割，获取第一个不为空的元素，计算其长度
//func lengthOfLastWord(s string) int {
//	wl := 0
//	sa := strings.Split(s, " ")
//	for i := len(sa) - 1; i >= 0; i-- {
//		fmt.Println(len(sa[i]), "  ", sa[i], "  ", sa[i] == "")
//		if sa[i] != "" {
//			wl = len(sa[i])
//			break
//		}
//	}
//	return wl
//}
