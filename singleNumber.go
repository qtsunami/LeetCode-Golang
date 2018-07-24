package main

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

// 说明：你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

// 示例 1:
// 输入: [2,2,1]
// 输出: 1
// 示例 2:

// 输入: [4,1,2,1,2]
// 输出: 4

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 1, 3, 4}

	ret := singleNumber(nums)

	fmt.Println(ret)
}

// 根据 按位异或 ^ (一个为 1 另一个为 0 的位设为 1) 来判断
func singleNumber(nums []int) int {

	returnValue := 0
	for _, n := range nums {
		returnValue = returnValue ^ n
	}
	return returnValue
}

// 最普遍的方法，将当前元素与除自己之外的所有元素全部比对一次
// func singleNumber(nums []int) int {

// 	len := len(nums)
// 	var returnValue int

// 	for i := 0; i < len; i++ {
// 		isSingle := true
// 		for j := 0; j < len; j++ {
// 			if i == j {
// 				continue
// 			}

// 			if nums[i] == nums[j] {
// 				isSingle = false
// 			}
// 		}

// 		if isSingle {
// 			returnValue = nums[i]
// 			break
// 		}

// 	}

// 	return returnValue
// }
