package main

import "fmt"

/*
BM67 不同路径的数目(一)
https://www.nowcoder.com/practice/166eaff8439d4cd898e3ba933fbc6358?tpId=295&tqId=685&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
一个机器人在m×n大小的地图的左上角（起点）。
机器人每次可以向下或向右移动。机器人要到达地图的右下角（终点）。
可以有多少种不同的路径从起点走到终点？
备注：m和n小于等于100,并保证计算结果在int范围内

示例1
输入：
2,1
返回值：
1

示例2
输入：
2,2
返回值：
2

*/

/*
递归
	有两种行走方式：如果向右走，相当于后面在一个(n−1)∗m(n-1)*m(n−1)∗m的矩阵中查找从左上角到右下角的不同路径数；而如果向下走，
    相当于后面在一个n∗(m−1)n*(m-1)n∗(m−1)的矩阵中查找从左上角到右下角不同的路径数。
	而(n−1)∗m(n-1)*m(n−1)∗m的矩阵与n∗(m−1)n*(m-1)n∗(m−1)的矩阵都是n∗mn*mn∗m矩阵的子问题，因此可以使用递归。

	(Python版本超时，不能完全通过)

	import sys
	#设置递归深度
	sys.setrecursionlimit(100000)
	class Solution:
		def uniquePaths(self , m: int, n: int) -> int:
			#矩阵只要有一条边为1，路径数就只有一种了
			if m == 1 or n == 1:
				return 1
			else:
				#两个分支
				return self.uniquePaths(m - 1, n) + self.uniquePaths(m, n - 1)

动态规划
	如果我们此时就在右下角的格子，那么能够到达该格子的路径只能是它的上方和它的左方两个格子，因此从左上角到右下角的路径数应该是从左上角到它的左边格子和上边格子的路径数之和，因此可以动态规划。
	step 1：用dp[i][j]表示大小为i*j的矩阵的路径数量，下标从1开始。
	step 2：（初始条件） 当i或者j为1的时候，代表矩阵只有一行或者一列，因此只有一种路径。
	step 3：（转移方程） 每个格子的路径数只会来自它左边的格子数和上边的格子数，因此状态转移为dp[i][j]=dp[i−1][j]+dp[i][j−1]
*/

func uniquePaths(m int, n int) int {
	// write code here
	if m == 1 || n == 1 {
		return 1
	}
	// 二维数组
	dp := make([][]int, m+1)
	// 赋值第一列为1，构建每行长度为n+1的slice
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	// 第一行也赋值1
	for j := 0; j < n+1; j++ {
		dp[0][j] = 1
	}
	// 行列往下或者往右走，两种可能都行，所以相加
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	// 拿右下角的值
	return dp[m-1][n-1]
}

func main06() {
	paths := uniquePaths(2, 2)
	fmt.Println(paths)
}
