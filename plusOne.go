package main

// 给定一个非负整数组成的非空数组，在该数的基础上加一，返回一个新的数组。

// 最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

// 你可以假设除了整数 0 之外，这个整数不会以零开头。

// 示例 1:

// 输入: [1,2,3]
// 输出: [1,2,4]
// 解释: 输入数组表示数字 123。
// 示例 2:

// 输入: [4,3,2,1]
// 输出: [4,3,2,2]
// 解释: 输入数组表示数字 4321。
import (
	"fmt"
	"math"
)

func main() {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 4, 2, 1, 3, 4, 4, 3, 2, 4, 6, 7, 8, 9, 9, 4, 3, 2, 3, 4, 5, 6, 6, 7, 6, 2}
	ret := plusOne(digits)

	fmt.Println(ret)
}

// 极限值（数组元素过多时会报错）
// func plusOne(digits []int) []int {

// 	var sum int
// 	ret := []int{}

// 	len := len(digits)

// 	for i := 0; i < len; i++ {
// 		sum += digits[i] * int(math.Pow10(len-i-1))
// 	}

// 	sum += 1

// 	first := sum / int(math.Pow10(len))
// 	if first > 0 {
// 		ret = append(ret, first)
// 		sum = int(sum) % int(math.Pow10(len))
// 	}

// 	for i := 0; i < len; i++ {
// 		ret = append(ret, sum/int(math.Pow10(len-i-1)))
// 		sum = int(sum) % int(math.Pow10(len-i-1))
// 	}

// 	return ret
// }

// 临界值未考虑（错误版本）
// func plusOne(digits []int) []int {

// 	lens := len(digits)
// 	digits[lens-1] = digits[lens-1] + 1

// 	for i := lens - 1; i >= 0; i++ {
// 		if digits[i] > 9 {
// 			digits[i-1], digits[i] = digits[i-1]+1, digits[i]-10
// 		} else {
// 			break
// 		}
// 	}

// 	return digits
// }

// 思考点： 该题主要难点有两点：一在于临界值的处理,ex. [1, 2, 9]，二在于数组元素+1的处理, ex. [9]或者[9, 9]
func plusOne(digits []int) []int {

	lens := len(digits)
	digits[lens-1] = digits[lens-1] + 1
	ret := []int{}

	for i := lens - 1; i >= 0; i-- {
		if digits[i] > 9 && i > 0 {
			digits[i-1], digits[i] = digits[i-1]+1, digits[i]-10
		} else {
			break
		}
	}

	if digits[0] > 9 {
		ret = append(ret, 1)
		ret = append(ret, digits[0]-10)
	} else {
		ret = append(ret, digits[0])
	}

	for j := 1; j < lens; j++ {
		ret = append(ret, digits[j])
	}

	return ret
}
