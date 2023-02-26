package main

import (
	"fmt"
)

/*
NC88 寻找第K大
https://www.nowcoder.com/practice/e016ad9b7f0b45048c58a9f27ba618bf?tpId=117&tqId=37791&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。

给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。
要求：时间复杂度 O(nlogn)，空间复杂度 O(1)

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

// 快速排序
//func findKth(a []int, n int, K int) int {
//	rand.Seed(time.Now().UnixNano())
//	var quickFind func(int, int)
//	var res int
//	quickFind = func(left, right int) {
//		// 右侧索引
//		r := rand.Int()%(right-left+1) + left
//		a[r], a[left] = a[left], a[r]
//		eIndex := left
//		for i := left + 1; i <= right; i++ {
//			if a[i] >= a[left] {
//				eIndex++
//				a[i], a[eIndex] = a[eIndex], a[i]
//			}
//		}
//		a[left], a[eIndex] = a[eIndex], a[left]
//		if eIndex == K-1 {
//			res = a[eIndex]
//			return
//		} else if eIndex > K-1 {
//			quickFind(left, eIndex-1)
//		} else {
//			quickFind(eIndex+1, right)
//		}
//	}
//	quickFind(0, n-1)
//	return res
//}

func findKth(a []int, n int, K int) int {
	return qsort(a, 0, n-1, K)
}

func qsort(nums []int, left, right, k int) int {
	if right > left {
		// 分区
		pos := partition(nums, left, right)
		// 递归
		if pos == len(nums)-k {
			// 最终推出调节
			return nums[pos]
		} else if pos > len(nums)-k {
			return qsort(nums, left, pos-1, k)
		} else {
			return qsort(nums, pos+1, right, k)
		}
	} else {
		// 左边大于等于右边，长度为1， 直接返回
		return nums[left]
	}
}

func partition(nums []int, left, right int) int {
	// 定基准
	pivot := nums[left]
	for left < right {
		for left < right && pivot <= nums[right] {
			right--
		}
		for left < right && pivot >= nums[left] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
		// left++
		// right--
	}
	// nums[left] = pivot
	return left
}

func main() {
	kth := findKth([]int{1, 3, 5, 2, 2}, 5, 3)
	fmt.Println(kth)
}
