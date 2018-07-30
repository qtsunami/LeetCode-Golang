package main

// 给定两个数组，写一个方法来计算它们的交集。

// 例如:
// 给定 nums1 = [1, 2, 2, 1], nums2 = [2, 2], 返回 [2, 2].

// 注意：

//    输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
//    我们可以不考虑输出结果的顺序。
// 跟进:

// 如果给定的数组已经排好序呢？你将如何优化你的算法？
// 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
// 如果nums2的元素存储在磁盘上，内存是有限的，你不能一次加载所有的元素到内存中，你该怎么办？

import (
	"fmt"
	// "sort"
)

func main() {

	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 1, 2, 4, 1, 5}

	ret := intersect(nums1, nums2)

	fmt.Println(ret)
}

func intersect(nums1 []int, nums2 []int) []int {
	// 思路： 将其中一个数组，以值为键，以出现的次数为值。 将另一个数组进行循环对比，查看该值是否在新数组中且次数大于0，如符合代表互为交集，并将次数减1
	dict := make(map[int]int)
	ret := []int{}

	for _, v := range nums1 {
		dict[v]++
	}

	for _, v := range nums2 {
		if _, ok := dict[v]; ok && dict[v] > 0 {
			ret = append(ret, v)
			dict[v]--
		}
	}

	return ret
}
