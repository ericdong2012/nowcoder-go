package main

import "fmt"

/*
BM79 打家劫舍(二)
https://www.nowcoder.com/practice/a5c127769dd74a63ada7bff37d9c5815?tpId=295&tqId=2285837&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
你是一个经验丰富的小偷，准备偷沿湖的一排房间，每个房间都存有一定的现金，为了防止被发现，你不能偷相邻的两家，即，如果偷了第一家，就不能再偷第二家，如果偷了第二家，那么就不能偷第一家和第三家。
沿湖的房间组成一个闭合的圆形，即第一个房间和最后一个房间视为相邻。
给定一个长度为n的整数数组nums，数组中的元素表示每个房间存有的现金数额，请你计算在不被发现的前提下最多的偷窃金额。

示例1
输入：
[1,2,3,4]
返回值：
6
说明：
最优方案是偷第 2 4 个房间

示例2
输入：
[1,3,6]
返回值：
6
说明：
由于 1 和 3 是相邻的，因此最优方案是偷第 3 个房间

*/

// 第一家和最后一家不能同时取到
/*
	    情况1：偷第一家的钱，不偷最后一家的钱。初始状态与状态转移不变，只是遍历的时候数组最后一位不去遍历。
	    情况2：偷最后一家的请，不偷第一家的钱。初始状态就设定了dp[1]=0，第一家就不要了，然后遍历的时候也会遍历到数组最后一位。

		class Solution:
		    def rob(self , nums: List[int]) -> int:
		        #dp[i]表示长度为i的数组，最多能偷取多少钱
		        dp1 = [0 for i in range(len(nums) + 1)]
		        #选择偷了第一家
		        dp1[1] = nums[0]
		        #最后一家不能偷
		        for i in range(2, len(nums)):
		            #对于每家可以选择偷或者不偷
		            dp1[i] = max(dp1[i - 1], nums[i - 1] + dp1[i - 2])
				// 最后一家不偷
		        res = dp1[len(nums) - 1];

		        #第二次循环
		        dp2 = [0 for i in range(len(nums) + 1)]
		        #不偷第一家
		        dp2[1] = 0
		        #可以偷最后一家 dp2[len(nums)]
		        for i in range(2, len(nums) + 1):
		            #对于每家可以选择偷或者不偷
		            dp2[i] = max(dp2[i - 1], nums[i - 1] + dp2[i - 2])

		        return max(res, dp2[len(nums)])

*/

// 加了一个限制条件，首位不能同时偷
// 偷第一家，不偷最后一家，控制点在于i < n,   dp[1] = nums[0]
// 不偷第一家，偷最后一家, 控制点在于i <= n， dp1[1] = 0
func rob1(nums []int) int {
	n := len(nums)
	// 因为下方 dp[i-1] 要保证有值，所以dp[1] 得有值，dp 长度设置为n + 1
	dp := make([]int, n+1)
	// 偷第一家，不偷最后一家，控制点在于i < n
	dp[1] = nums[0]
	for i := 2; i < n; i++ {
		// 偷(前一个不偷) 和 不偷(前一个被偷， 前一个数字num[i-1]加上前一个偷的结果dp[i-2]) 比较下最大值
		dp[i] = max9(dp[i-1], nums[i-1]+dp[i-2])
	}
	// dp[n-1] 不偷最后一家
	res := dp[n-1]

	dp1 := make([]int, n+1)
	// 不偷第一家，偷最后一家, 控制点在于i <= n
	dp1[1] = 0
	for i := 2; i <= n; i++ {
		// 不偷 和 偷（前一个数字num[i-1]加上前一个偷的结果dp[i-2]） 比较下最大值
		dp1[i] = max9(dp1[i-1], nums[i-1]+dp1[i-2])
	}
	return max9(res, dp1[n])
}

func max9(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main18() {
	i := rob1([]int{1, 2, 3, 4})
	fmt.Println(i)
}
