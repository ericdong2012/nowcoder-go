package main

import (
	"fmt"
)

/*
BM82 买卖股票的最好时机(三)
https://www.nowcoder.com/practice/4892d3ff304a4880b7a89ba01f48daf9?tpId=295&tqId=1073487&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
假设你有一个数组prices，长度为n，其中prices[i]是某只股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1. 你最多可以对该股票有两笔交易操作，一笔交易代表着一次买入与一次卖出，但是再次购买前必须卖出之前的股票
2. 如果不能获取收益，请返回0
3. 假设买入卖出均无手续费

数据范围：1 \le n \le 10^51≤n≤10
5
 ，股票的价格满足 1 \le val\le 10^41≤val≤10
4

要求: 空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
[8,9,3,5,1,3]
复制
返回值：
4
复制
说明：
第三天(股票价格=3)买进，第四天(股票价格=5)卖出，收益为2
第五天(股票价格=1)买进，第六天(股票价格=3)卖出，收益为2
总收益为4。
示例2
输入：
[9,8,4,1]
复制
返回值：
0
复制
示例3
输入：
[1,2,8,3,8]
复制
返回值：
12
复制
说明：
第一笔股票交易在第一天买进，第三天卖出；第二笔股票交易在第四天买进，第五天卖出；总收益为12。
因最多只可以同时持有一只股票，所以不能在第一天进行第一笔股票交易的买进操作，又在第二天进行第二笔股票交易的买进操作（此时第一笔股票交易还没卖出），最后两笔股票交易同时在第三天卖出，也即以上操作不满足题目要求。
备注：
总天数不大于200000。保证股票每一天的价格在[1,100]范围内。

*/

/*
step 1：用dp[i][0]dp[i][0]dp[i][0]表示第i天不持股到该天为止的最大收益，dp[i][1]dp[i][1]dp[i][1]表示第i天持股，到该天为止的最大收益。
step 2：（初始状态） 第一天不持股，则总收益为0，dp[0][0]=0dp[0][0]=0dp[0][0]=0；第一天持股，则总收益为买股票的花费，此时为负数，dp[0][1]=−prices[0]dp[0][1] = -prices[0]dp[0][1]=−prices[0]。
step 3：（状态转移） 对于之后的每一天，如果当天不持股，有可能是前面的若干天中卖掉了或是还没买，因此到此为止的总收益和前一天相同，也有可能是当天卖掉股票，我们选择较大的状态dp[i][0]=max(dp[i−1][0],dp[i−1][1]+prices[i])dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])dp[i][0]=max(dp[i−1][0],dp[i−1][1]+prices[i])；
step4：如果当天持股，可能是前几天买入的还没卖，因此收益与前一天相同，也有可能是当天买入，减去买入的花费，同样是选取最大值：dp[i][1]=max(dp[i−1][1],dp[i−1][0]−prices[i])dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])dp[i][1]=max(dp[i−1][1],dp[i−1][0]−prices[i])。


class Solution:
    def maxProfit(self , prices: List[int]) -> int:
        n = len(prices)
        #dp[i][0]表示某一天不持股到该天为止的最大收益，dp[i][1]表示某天持股，到该天为止的最大收益
        dp = [[0] * 2 for i in range(n)]
        #第一天不持股，总收益为0
        dp[0][0] = 0
        #第一天持股，总收益为减去该天的股价
        dp[0][1] = -prices[0]
        #遍历后续每天，状态转移
        for i in range(1, n):
            dp[i][0] = max(dp[i - 1][0], dp[i - 1][1] + prices[i])
            dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])
        #最后一天不持股，到该天为止的最大收益
        return dp[n - 1][0]
*/

func maxProfit2(prices []int) int {
	// write code here
	// dp
	//n := len(prices)
	//dp := make([][]int, n)
	//for i := 0; i < len(dp); i++ {
	//	dp[i] = make([]int, 2)
	//}
	//// 不持股
	//dp[0][0] = 0
	//// 持股
	//dp[0][1] = -prices[0]
	//for i := 1; i < n; i++ {
	//	// 当天不持有： 之前不持有或者之前持有当天已经卖出
	//	dp[i][0] = max6(dp[i-1][0], dp[i-1][1]+prices[i])
	//	// 当天持有:   之前的没有卖或者之前不持有当天买了
	//	dp[i][1] = max6(dp[i-1][1], dp[i-1][0]-prices[i])
	//}
	//// 最后一天持股的收益
	//return dp[n-1][0]

	// 贪心
	/*
		class Solution:
		    def maxProfit(self , prices: List[int]) -> int:
		        res = 0
		        for i in range(1, len(prices)):
		            #只要某段在递增就有收益
		            if prices[i - 1] < prices[i]:
		                #收益累加
		                res += prices[i] - prices[i - 1]
		        return res

	*/
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i-1] < prices[i] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

func max6(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main21() {
	profit2 := maxProfit2([]int{8, 9, 2, 5, 4, 7, 1})
	fmt.Println(profit2)
}
