package main

import "fmt"

/*
BM64 最小花费爬楼梯
https://www.nowcoder.com/practice/6fe0302a058a4e4a834ee44af88435c7?tpId=295&tqId=2366451&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个整数数组 cost， 其中 cost[i]  是从楼梯第i个台阶向上爬需要支付的费用，下标从0开始。一旦你支付此费用，即可选择向上爬一个或者两个台阶。
你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
请你计算并返回达到楼梯顶部的最低花费。

示例1
输入：
[2,5,20]
返回值：
5
说明：
你将从下标为1的台阶开始，支付5 ，向上爬两个台阶，到达楼梯顶部。总花费为5

示例2
输入：
[1,100,1,1,1,90,1,1,80,1]
返回值：
6

说明：
你将从下标为 0 的台阶开始。
1.支付 1 ，向上爬两个台阶，到达下标为 2 的台阶。
2.支付 1 ，向上爬两个台阶，到达下标为 4 的台阶。
3.支付 1 ，向上爬两个台阶，到达下标为 6 的台阶。
4.支付 1 ，向上爬一个台阶，到达下标为 7 的台阶。
5.支付 1 ，向上爬两个台阶，到达下标为 9 的台阶。
6.支付 1 ，向上爬一个台阶，到达楼梯顶部。
总花费为 6

*/

func minCostClimbingStairs(cost []int) int {
	// write code here
	/*
		建立二维数组
		1 100 1 1 1 90 1 1 80 1
		0 0   1 2 2 3  3 4 4  5

		再建立方程
	*/
	//n := len(cost)
	//dp := make([]int, n)
	//dp[0] = cost[0]
	//dp[1] = cost[1]
	//for i := 2; i < n; i++ {
	//	dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	//}
	//return min(dp[n-1], dp[n-2])

	n := len(cost)
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	//stairs := minCostClimbingStairs([]int{2, 5, 20})
	stairs := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 90, 1, 1, 80, 1})
	fmt.Println(stairs)
}
