package _5_单调栈

/*
NC209 最短无序连续子数组
https://www.nowcoder.com/practice/d17f4abd1d114617b51e951027be312e?tpId=117&tqId=39398&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=1579&title=


描述
给定一个整数数组，你需要找出一个连续子数组，将这个子数组升序排列后 整个数组都将是升序数组。
请你找出满足题设的最短的子数组。

示例1
输入：
[2,6,4,8,10,9,15]
返回值：
5
说明：
只需对 6,4,8,10,9 排序即可得到升序数组

示例2
输入：
[1,2,3,5,4]
返回值：
2
说明：
对 5,4 排序即可得到升序数组

*/

import "math"

// 找的是两个波峰，用单调栈解决
func findUnsortedSubarray( nums []int ) int {
	// write code here
	l, r := -1, -1
	n := len(nums)
	mx, md := math.MinInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		// 从前往后走, 找到开始降序的那个，记录到r
		if nums[i] >= mx {
			mx = nums[i]
		} else {
			r = i
		}
		// 从后往前走, 找到升序的那个，记录到l中
		if nums[n-i-1] <= md {
			md = nums[n-i-1]
		} else {
			l = n-i-1
		}
	}
	// 只要不是完全升序
	if r != -1 {
		return r-l+1
	}
	return 0
}