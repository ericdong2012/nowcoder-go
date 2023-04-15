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

/*
思路1： dp
思路2： 比较奇偶位的和大小 其实不行， 不一定下一个非得偷奇数或者偶数，还是一个大小问题
        // 5，2，3，7
		// 5，2，3，2
*/
func max8(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		// 5，2，3，7
		// 5，2，3，2
		// 不偷  +  偷（前一个数字num[i-1]加上更前一个偷的结果dp[i-2]） 比较较大值
		dp[i] = max8(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[n]
}

func main17() {
	i := rob([]int{5, 2, 3, 7})
	fmt.Println(i)
}
