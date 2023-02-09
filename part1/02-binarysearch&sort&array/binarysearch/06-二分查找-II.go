package main

import "fmt"

/*
NC105 二分查找-II
https://www.nowcoder.com/practice/4f470d1d3b734f8aaf2afb014185b395?tpId=117&tqId=37829&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
请实现有重复数字的升序数组的二分查找
给定一个 元素有序的（升序）长度为n的整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的第一个出现的target，如果目标值存在返回下标，否则返回 -1

示例1
输入：
[1,2,4,4,5],4
返回值：
2
说明：
从左到右，查找到第1个为4的，下标为2，返回2

示例2
输入：
[1,2,4,4,5],3
返回值：
-1

示例3
输入：
[1,1,1,1,1],1
返回值：
0
*/

func search(nums []int, target int) int {
	// write code here
	left, right := 0, len(nums)-1
	stack := []int{}
	for left <= right {
		//if nums[left] == target {
		//	stack = append(stack, left)
		//	left++
		//} else if nums[right] == target {
		//	stack = append(stack, right)
		//	right -= 1
		//} else {
		//	left += 1
		//	right -= 1
		//}
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			stack = append(stack, mid)
			// 只要找最小的索引，不是找完整的
			right = mid  -1
		}
	}

	if len(stack) == 0 {
		return -1
	}

	minV := stack[0]
	for i := 0; i < len(stack); i++ {
		minV = min2(minV, stack[i])
	}
	return minV

}

func min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	i := search([]int{1, 2, 4, 4, 5}, 4)
	fmt.Println(i)
}
