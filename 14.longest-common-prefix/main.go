package main

import (
	"fmt"
	"strings"
)

//编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。

//示例 1：
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//
//示例 2：
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。

//提示：
//1 <= strs.length <= 200
//0 <= strs[i].length <= 200
//strs[i] 仅由小写英文字母组成

func main() {
	strs := []string{"flower", "fl", "s"}
	ret := longestCommonPrefix(strs)

	fmt.Println(ret)
}

func longestCommonPrefix(strs []string) string {

	var commonPrefix []string
	shortStr := strs[0]
	over := false

	for i := 0; i < len(shortStr) && !over; i++ {
		for _, str := range strs {
			if i == len(str) || shortStr[i:i+1] != str[i:i+1] {
				over = true
			}
		}
		if !over {
			commonPrefix = append(commonPrefix, shortStr[i:i+1])
		}
	}
	return strings.Join(commonPrefix, "")
}
