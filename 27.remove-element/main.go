package main

import "fmt"

// https://leetcode.cn/problems/remove-element/
//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
//说明:
//为什么返回数值是整数，但输出的答案是数组呢?
//请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//你可以想象内部操作如下:
//// nums 是以“引用”方式传递的。也就是说，不对实参作任何拷贝
//int len = removeElement(nums, val);
//// 在函数里修改输入数组对于调用者是可见的。
//// 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
//for (int i = 0; i < len; i++) {
//print(nums[i]);
//}
//示例 1：
//输入：nums = [3,2,2,3], val = 3
//输出：2, nums = [2,2]
//解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。

func main() {

	var a = []int{3, 1, 1, 3}
	fmt.Println(removeElement1(a, 1))
	fmt.Println(a)
}

func removeElement(nums []int, val int) int {
	newLen := 0
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if nums[i] == val {
			for j := numsLen - 1; j >= 0; j-- {
				if nums[j] != val && i < j {
					nums[i], nums[j] = nums[j], nums[i]
					newLen++
					break
				}
			}
		} else {
			newLen++
		}
	}
	return newLen
}

// 更加简洁的代码
func removeElement1(nums []int, val int) int {
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	return j
}
