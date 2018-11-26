package main

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

// 示例:

// 输入: [0,1,0,3,12]
// 输出: [1,3,12,0,0]
// 说明:

// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。

import (
	"fmt"
)

func main() {

	nums := []int{0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0, 1}
	moveZeroes(nums)

	fmt.Println(nums)
}

func moveZeroes(nums []int) {

	len := len(nums)

	for i := 0; i < len; i++ {
		for j := 0; j < len-1; j++ {
			if nums[j] == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

}
