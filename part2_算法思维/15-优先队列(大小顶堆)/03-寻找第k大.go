package _5_优先队列_大小顶堆_

/*
https://www.nowcoder.com/practice/e016ad9b7f0b45048c58a9f27ba618bf

有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。


示例1
输入：
[1,3,5,2,2],5,3
返回值：
2

示例2
输入：
[10,10,9,9,8,7,5,6,4,3,4,2],12,3
返回值：
9
说明：
去重后的第3大是8，但本题要求包含重复的元素，不用去重，所以输出9

*/

// 方式1： 直接排序， 取第k个
// 方式2： 快排(分治)
func findKth(a []int, n int, K int) int {
	return qsort(a, 0, n-1, K)
}

func qsort(nums []int, left, right, k int) int {
	// 会有重复值
	if right >= left {
		// 分区，找索引
		pos := partition(nums, left, right)
		// 递归
		// 最终结果 找到第k大， 根据索引找到对应的值
		if pos == len(nums)-k {
			return nums[pos]
			// 如果它右边的元素比k−1大，说明第k大在其左边，直接二分法进入左边
		} else if pos > len(nums)-k {
			return qsort(nums, left, pos-1, k)
			// 如果它左边的元素比k−1少，说明第k大在其右边，直接二分法进入右边
		} else {
			return qsort(nums, pos+1, right, k)
		}
	} else {
		return -1
	}
}

func partition(nums []int, left, right int) int {
	// 定基准
	pivot := nums[left]
	for left < right {
		for left < right && pivot > nums[left] {
			left++
		}
		for left < right && pivot < nums[right] {
			right--
		}
		// 将num调整顺序 小的在前，大的在后
		nums[left], nums[right] = nums[right], nums[left]
	}
	// 返回 k 附近的索引
	return left
}
