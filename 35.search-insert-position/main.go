package main

import "fmt"

// https://leetcode.cn/problems/search-insert-position/
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法

// 示例 1: 输入: nums = [1,3,5,6], target = 5  输出: 2
// 示例 2: 输入: nums = [1,3,5,6], target = 2  输出: 1
// 示例 3: 输入: nums = [1,3,5,6], target = 7  输出: 4

// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 为 无重复元素 的 升序 排列数组
// -104 <= target <= 104

func main() {

	var nums = []int{1, 3, 5, 8, 11, 15}
	//var nums = []int{1, 3}
	var target = 2

	index := searchInsert(nums, target)
	fmt.Println(index)
}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	targetIndex := n

	for left <= right {
		mid := (right-left)>>1 + left
		fmt.Println("right - left >> 1 + left", mid, "  ", nums[mid])
		if target <= nums[mid] {
			targetIndex = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return targetIndex
}
