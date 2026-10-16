package main

/*
NC145 01背包
https://www.nowcoder.com/practice/2820ea076d144b30806e72de5e5d4bbf?tpId=117&tqId=37856&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=593&title=

已知一个背包最多能容纳体积之和为v的物品, 现有n个物品，  第i个物品的体积为 vi , 重量为 wi
求当前背包最多能装多大重量的物品?

示例1
输入：
10,2,[[1,3],[10,4]]
返回值：
4
说明：
第一个物品的体积为1，重量为3，第二个物品的体积为10，重量为4。只取第二个物品可以达到最优方案，取物重量为4

示例2
输入：
10,2,[[1,3],[9,8]]
返回值：
11
说明：
两个物品体积之和等于背包能装的体积，所以两个物品都取是最优方案

*/

// V是体积
// n是个数
// 体积，重量的数组
func knapsack(V int, n int, vw [][]int) int {
	// 索引是体积, 存的是重量(好好体会)    长度是最大体积+1, 因为要求最大能到V
	var dp = make([]int, V+1)
	for i := 0; i < n; i++ {
		var v, w = vw[i][0], vw[i][1]
		for j := V; j >= v; j-- {
			dp[j] = Max(dp[j], dp[j-v]+w)
		}
		//for j := V; j >= vw[i][0]; j-- {
		//	dp[j] = Max(dp[j], dp[j-vw[i][0]]+vw[i][1])
		//}
	}
	return dp[V]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
