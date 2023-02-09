package main

import "fmt"

/*
BM80 买卖股票的最好时机(一)
https://www.nowcoder.com/practice/64b4262d4e6d4f6181cd45446a5821ec?tpId=295&tqId=625&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
假设你有一个数组prices，长度为n，其中prices[i]是股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1.你可以买入一次股票和卖出一次股票，并非每天都可以买入或卖出一次，总共只能买入和卖出一次，且买入必须在卖出的前面的某一天
2.如果不能获取到任何利润，请返回0
3.假设买入卖出均无手续费

数据范围： 0 \le n \le 10^5 , 0 \le val \le 10^40≤n≤10
5
 ,0≤val≤10
4

要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
[8,9,2,5,4,7,1]
复制
返回值：
5
复制
说明：
在第3天(股票价格 = 2)的时候买入，在第6天(股票价格 = 7)的时候卖出，最大利润 = 7-2 = 5 ，不能选择在第2天买入，第3天卖出，这样就亏损7了；同时，你也不能在买入前卖出股票。
示例2
输入：
[2,4,1]
复制
返回值：
2
复制
示例3
输入：
[3,2,1]
复制
返回值：
0

*/


//func maxProfit(prices []int) int {
//	// write code here
//	dp := []int{}
//	res := 0
//	for i := 0; i < len(prices); i++ {
//		temp := 0
//		for j := i + 1; j < len(prices); j++ {
//			temp = max5(temp, prices[j]-prices[i])
//		}
//		dp = append(dp, temp)
//	}
//
//	for i := 0; i < len(dp); i++ {
//		res = max5(res, dp[i])
//	}
//
//	return res
//}
//
//func max5(a, b int) int {
//	if a < b {
//		return b
//	} else {
//		return a
//	}
//}

/*

class Solution:
    def maxProfit(self , prices: List[int]) -> int:
        #维护最大收益
        res = 0
        #排除特殊情况
        if len(prices) == 0:
            return res
        #维护最低股票价格
        Min = prices[0]
        #遍历后续股票价格
        for i in range(1, len(prices)):
            #如果当日价格更低则更新最低价格
            Min = min(Min, prices[i])
            #维护最大值
            res = max(res, prices[i] - Min)
        return res
*/

func maxProfit(prices []int) int {
	// write code here
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	maxV := 0
	minPrice := prices[0]
	for _, v := range prices {
		minPrice = min5(minPrice, v)
		maxV = max5(maxV, v-minPrice)
	}
	return maxV
}

func max5(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min5(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	profit := maxProfit([]int{8, 9, 2, 5, 4, 7, 1})
	fmt.Println(profit)
}
