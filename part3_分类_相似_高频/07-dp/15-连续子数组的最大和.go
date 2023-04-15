package main

/*
BM72 连续子数组的最大和
https://www.nowcoder.com/practice/459bd355da1549fa8a49e350bf3df484?tpId=196&tqId=37130&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%25E8%25BF%259E%25E7%25BB%25AD%25E5%25AD%2590%25E6%2595%25B0%25E7%25BB%2584%25E7%259A%2584%25E6%259C%2580%25E5%25A4%25A7%25E5%2592%258C%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=&title=%E8%BF%9E%E7%BB%AD%E5%AD%90%E6%95%B0%E7%BB%84%E7%9A%84%E6%9C%80%E5%A4%A7%E5%92%8C

描述
输入一个长度为n的整型数组array，数组中的一个或连续多个整数组成一个子数组，子数组最小长度为1。求所有子数组的和的最大值。
要求:时间复杂度为 O(n)，空间复杂度为 O(n)
进阶:时间复杂度为 O(n)，空间复杂度为 O(1)

示例1
输入：
[1, -2, 3, 10, -4, 7, 2, -5]
返回值：
18
说明：
经分析可知，输入数组的子数组[3, 10, -4, 7, 2]可以求得最大和为18

示例2
输入：
[2]
返回值：
2

示例3
输入：
[-10]
返回值：
-10

*/

/*

注意理解题目意思： 示例1 中元素有正有副，连续最大即可。

step1：可以用dp数组表示以下标iii为终点的最大连续子数组和。
step2：遍历数组，每次遇到一个新的数组元素，连续的子数组要么加上变得更大，要么这个元素本身就更大，要么会更小，更小我们就舍弃，因此状态转移为dp[i]=max(dp[i−1]+array[i], array[i])。
step3：因为连续数组可能会断掉，每一段只能得到该段最大值，因此我们需要维护一个最大值。

*/

func FindGreatestSumOfSubArray(array []int) int {
	// write code here
	// [1, -2, 3, 10, -4, 7, 2, -5]    [3, 10, -4, 7, 2]  => 18
	// 中间值
	dp := make([]int, len(array))
	dp[0] = array[0]
	// 结果值
	maxV := dp[0]
	for i := 1; i < len(array); i++ {
		dp[i] = max(dp[i-1]+array[i], array[i])
		maxV = max(maxV, dp[i])
	}
	return maxV
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
