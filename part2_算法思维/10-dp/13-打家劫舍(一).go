package main

import "fmt"

/*
BM78 打家劫舍(一)
https://www.nowcoder.com/practice/c5fbf7325fbd4c0ea3d0c3ea6bc6cc79?tpId=295&tqId=2285793&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
你是一个经验丰富的小偷，准备偷沿街的一排房间，每个房间都存有一定的现金，为了防止被发现，你不能偷相邻的两家，即，如果偷了第一家，就不能再偷第二家；如果偷了第二家，那么就不能偷第一家和第三家。
给定一个整数数组nums，数组中的元素表示每个房间存有的现金数额，请你计算在不被发现的前提下最多的偷窃金额。

示例1
输入：
[1,2,3,4]
返回值：
6
说明：
最优方案是偷第 2，4 个房间

示例2
输入：
[1, 3, 6]
返回值：
7
说明：
最优方案是偷第 1，3个房间

示例3
输入：
[2, 10, 5]
返回值：
10
说明：
最优方案是偷第2个房间

*/

func max8(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return max8(nums[0], nums[1])
	}

	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]    // dp[1] = nums[0]  代表可以偷头
	// i <=n  是为了能走到最后，偷尾， 对应的当前值是 nums[i-1] (细品)
	for i := 2; i <= n; i++ {
		// 5，2，3，7
		// 0, 5，5，8，8/9/12   （首尾偷）
		// 不偷(当前不偷，前一个偷了，保持前一个数据不变)  +  偷（当前偷，当前num[i-1]加上更前一个偷的结果dp[i-2]）    比较较大值
		dp[i] = max8(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[n]
}

// 原来的想法(没有问题，但是头尾同时偷应该怎么表示没有考虑清楚)
//func rob(nums []int) int {
//	if len(nums) < 2 {
//		return nums[0]
//	}
//
//	dp := make([]int, len(nums))
//	dp[0] = nums[0]
//	dp[1] = nums[1]
//	for i := 2; i < len(nums); i++ {
//		// 问题在于首尾同时偷这种情况没有考虑进去  5, 2, 3, 7
//		dp[i] = max(dp[i-1], dp[i-2] + nums[i])
//	}
//	return dp[len(nums)-1]
//}

func main17() {
	//i := rob([]int{5, 2, 3, 7})
	i := rob([]int{1, 3, 6})
	fmt.Println(i)
}
