package main

import "fmt"

// 题解： https://leetcode.cn/problems/roman-to-integer/solutions/2104303/luo-ma-shu-zi-zhuan-zheng-shu-ti-jie-by-uyqca/

//罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

//例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//提示：
//1 <= s.length <= 15
//s 仅含字符 ('I', 'V', 'X', 'L', 'C', 'D', 'M')
//题目数据保证 s 是一个有效的罗马数字，且表示整数在范围 [1, 3999] 内
//题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
//IL 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX 。
//关于罗马数字的详尽书写规则，可以参考 罗马数字 - Mathematics 。

// 最长的是： MMMDCCCCXXXVIII  = 3888

const (
	I  = 1
	IV = 4
	V  = 5
	IX = 9
	X  = 10
	XL = 40
	L  = 50
	XC = 90
	C  = 100
	CD = 400
	D  = 500
	CM = 900
	M  = 1000
)

var romanMap = map[string]int{
	"I":  I,
	"IV": IV,
	"V":  V,
	"IX": IX,
	"X":  X,
	"XL": XL,
	"L":  L,
	"XC": XC,
	"C":  C,
	"CD": CD,
	"D":  D,
	"CM": CM,
	"M":  M,
}

func main() {
	romanStr := "MMMDCCCLXXXVIII"
	ret := romanToInt(romanStr)
	fmt.Printf("roman number %s to int is %d", romanStr, ret)
}

func romanToInt(s string) int {
	sb := []byte(s)
	var ret int = 0
	sbLen := len(sb)
	for i := 0; i < sbLen; i++ {
		if i < sbLen-1 && romanMap[string(sb[i])] < romanMap[string(sb[i+1])] {
			ret = ret - romanMap[string(sb[i])]
		} else {
			ret += romanMap[string(sb[i])]
		}
	}
	return ret
}
