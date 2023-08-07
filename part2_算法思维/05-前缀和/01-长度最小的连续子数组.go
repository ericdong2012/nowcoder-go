package _5_前缀和

/*
NC202 长度最小的连续子数组
https://www.nowcoder.com/practice/10dd5f8c5d984aa3bd69788d86aaef23?tpId=196&tqId=39383&ru=/exam/oj

知识点
二分
前缀和
双指针

描述
给定一个数组 nums 和一个正整数 target , 找出满足和大于等于 target 的长度最短的连续子数组并返回其长度，如果不存在这种子数组则返回 0。

示例1
输入：
[1,2,4,4,1,1,1],9
返回值：
3

示例2
输入：
[1,4,4,4,1,1,1],3
返回值：
1

*/

func minSubarray( nums []int ,  target int ) int {
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
