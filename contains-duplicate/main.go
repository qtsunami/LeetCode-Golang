package main

// 给定一个整数数组，判断是否存在重复元素。

// 如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

// 示例 1:

// 输入: [1,2,3,1]
// 输出: true
// 示例 2:

// 输入: [1,2,3,4]
// 输出: false
// 示例 3:

// 输入: [1,1,1,3,3,4,3,2,4,2]
// 输出: true

import (
	"fmt"
)

func main() {
	arr := []int{1, 4, 3, 7, 3, 9, 16, 17}

	ret := containsDuplicate(arr)

	fmt.Println(ret)
}

// func containsDuplicate(nums []int) bool {

// 	numLen := len(nums)

// 	for i := 0; i < numLen; i++ {
// 		for j := i + 1; j < numLen; j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// func containsDuplicate(nums []int) bool {
// 	m := make(map[int]int)

// 将键值进行调换
// 	for index, value := range nums {
// 		m[value] = index
// 	}

// 判断比较原数组nums的长度与键值调换后的数组m的长度
// 	if len(m) < len(nums) {
// 		return true
// 	}

// 	return false
// }

// 根据value出现的次数进行判断
func containsDuplicate(nums []int) bool {
	countMap := make(map[int]int, len(nums))

	for _, num := range nums {
		fmt.Println(num, countMap[num])

		if countMap[num] == 1 {
			// return true
		}
		countMap[num]++
	}
	fmt.Println(countMap)
	return false
}
