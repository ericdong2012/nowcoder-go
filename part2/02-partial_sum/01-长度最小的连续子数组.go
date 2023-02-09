package main

import (
	"fmt"
)

/*
NC202 长度最小的连续子数组
https://www.nowcoder.com/practice/10dd5f8c5d984aa3bd69788d86aaef23?tpId=117&tqId=39391&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5059&title=

描述
给定一个数组 nums 和一个正整数 target , 找出满足和大于等于 target 的长度最短的连续子数组并返回其长度，如果不存在这种子数组则返回 0。

示例1
输入：
[1,2,4,4,1,1,1], 9
返回值：
3

示例2
输入：
[1,4,4,4,1,1,1], 3
返回值：
1

*/

func minSubarray(nums []int, target int) int {
	// write code here
	/*

		右指针每次向外扩展我们的这个窗口，
		判断如果左指针可以收缩，一直让左指针收缩，这样我们最后得到的就是满足条件的了，取出最小值即可

	*/

	left, right := 0, 0
	res := len(nums)
	curSum := nums[0]
	for right != len(nums) {
		// 当前和 大于目标值，窗口减小，左指针收缩，cursum减小
		if curSum >= target && left <= right {
			res = min(res, right-left+1)
			curSum -= nums[left]
			left++
		} else {
			// 右指针往外扩展，将当前值放入到cursum 中求和
			right++
			if right < len(nums) {
				curSum += nums[right]
			}
		}
	}
	return res
}

func minSubarray1(nums []int, target int) int {
	// write code here
	left, right := 0, 0
	res := len(nums)
	curSum := nums[0]
	for right != len(nums) {
		if curSum >= target && left <= right {
			res = min(res, right-left+1)
			curSum -= nums[left]
			left++
		} else {
			right++
			if right < len(nums) {

				curSum += nums[right]
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	subarray := minSubarray1([]int{1, 1, 1, 2, 3, 4, 5, 4}, 5)
	fmt.Println(subarray)
}
