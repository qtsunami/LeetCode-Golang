package main

// 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

// 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

// 示例:

// 给定 nums = [2, 7, 11, 15], target = 9

// 因为 nums[0] + nums[1] = 2 + 7 = 9
// 所以返回 [0, 1]

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 1, 3, 4}

	target := 7

	ret := twoSum(nums, target)

	fmt.Println(ret)

}

func twoSum(nums []int, target int) []int {

	find := false
	ret := make([]int, 2)
	for k, v := range nums {
		tux := target - v
		ret[0] = k
		for i := k + 1; i < len(nums); i++ {
			if nums[i] == tux {
				ret[1] = i
				find = true
				break
			}
		}

		if find {
			break
		}
	}

	return ret
}
