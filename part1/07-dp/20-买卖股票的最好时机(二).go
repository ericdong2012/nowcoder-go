package main

import "fmt"

/*
BM81 买卖股票的最好时机(二)
https://www.nowcoder.com/practice/9e5e3c2603064829b0a0bbfca10594e9?tpId=295&tqId=1073471&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
假设你有一个数组prices，长度为n，其中prices[i]是某只股票在第i天的价格，请根据这个价格数组，返回买卖股票能获得的最大收益
1. 你可以多次买卖该只股票，但是再次购买前必须卖出之前的股票
2. 如果不能获取收益，请返回0
3. 假设买入卖出均无手续费

数据范围： 1 \le n \le 1 \times 10^51≤n≤1×10
5
  ， 1 \le prices[i] \le 10^41≤prices[i]≤10
4

要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
进阶：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
[8,9,2,5,4,7,1]
复制
返回值：
7
复制
说明：
在第1天(股票价格=8)买入，第2天(股票价格=9)卖出，获利9-8=1
在第3天(股票价格=2)买入，第4天(股票价格=5)卖出，获利5-2=3
在第5天(股票价格=4)买入，第6天(股票价格=7)卖出，获利7-4=3
总获利1+3+3=7，返回7
示例2
输入：
[5,4,3,2,1]
复制
返回值：
0
复制
说明：
由于每天股票都在跌，因此不进行任何交易最优。最大收益为0。
示例3
输入：
[1,2,3,4,5]
复制
返回值：
4
复制
说明：
第一天买进，最后一天卖出最优。中间的当天买进当天卖出不影响最终结果。最大收益为4。
备注：
总天数不大于200000。保证股票每一天的价格在[1,100]范围内。

*/

/*
区别： 两次买进，也可不买进

二维数组
dp方程

8	9	3	5	1	3


持股情况有了5种变化，我们用：
dp[i][0]表示到第i天为止没有买过股票的最大收益
dp[i][1]表示到第i天为止买过一次股票还没有卖出的最大收益
dp[i][2]表示到第i天为止买过一次也卖出过一次股票的最大收益
dp[i][3]表示到第i天为止买过两次只卖出过一次股票的最大收益
dp[i][4]表示到第i天为止买过两次同时也买出过两次股票的最大收益


step 1：（初始状态） 与上述提到的题类似，第0天有买入了和没有买两种状态：dp[0][0]=0、dp[0][1]=−prices[0]。
step 2：状态转移： 对于后续的每一天，如果当天还是状态0，则与前一天相同，没有区别；
step 3：如果当天状态为1，可能是之前买过了或者当天才第一次买入，选取较大值：dp[i][1]=max(dp[i−1][1],dp[i−1][0]−prices[i])dp[i][1] = max(dp[i - 1][1], dp[i - 1][0] - prices[i])dp[i][1]=max(dp[i−1][1],dp[i−1][0]−prices[i])；
step 4：如果当天状态是2，那必须是在1的状态下（已经买入了一次）当天卖出第一次，或者早在之前就卖出只是还没买入第二次，选取较大值：dp[i][2]=max(dp[i−1][2],dp[i−1][1]+prices[i])dp[i][2] = max(dp[i - 1][2], dp[i - 1][1] + prices[i])dp[i][2]=max(dp[i−1][2],dp[i−1][1]+prices[i])；
step 5：如果当天状态是3，那必须是在2的状态下（已经卖出了第一次）当天买入了第二次，或者早在之前就买入了第二次，只是还没卖出，选取较大值：dp[i][3]=max(dp[i−1][3],dp[i−1][2]−prices[i])dp[i][3] = max(dp[i - 1][3], dp[i - 1][2] - prices[i])dp[i][3]=max(dp[i−1][3],dp[i−1][2]−prices[i]);
step 6：如果当天是状态4，那必须是在3的状态下（已经买入了第二次）当天再卖出第二次，或者早在之前就卖出了第二次，选取较大值：dp[i][4]=max(dp[i−1][4],dp[i−1][3]+prices[i])dp[i][4] = max(dp[i - 1][4], dp[i - 1][3] + prices[i])dp[i][4]=max(dp[i−1][4],dp[i−1][3]+prices[i])。
step 7：最后我们还要从0、第一次卖出、第二次卖出中选取最大值，因为有可能没有收益，也有可能只交易一次收益最大。

因为状态转移的时候，辅助数组只使用到了第i列和第i-1列，因此可以不使用数组，直接用变量代替，优化空间复杂度。
*/

func maxProfit3(prices []int) int {
	// write code here
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 5)
	}
	// 不持股
	dp[0][0] = 0
	// 持股
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0
	for i := 1; i < n; i++ {
		//当天还是状态0，没有买， 则与前一天相同
		dp[i][0] = dp[i-1][0]
		// 当天状态为1(第一次买入)，可能是之前买过了或者当天才第一次买入
		dp[i][1] = max7(dp[i-1][1], dp[i-1][0]-prices[i])
		// 当天状态是2（第一次卖出），那必须是在1的状态下（已经买入了一次）当天卖出第一次，或者早在之前就卖出只是还没买入第二次
		dp[i][2] = max7(dp[i-1][2], dp[i-1][1]+prices[i])
		// 当天状态是3(第二次买入)，那必须是在2的状态下（已经卖出了第一次）当天买入了第二次，或者早在之前就买入了第二次，只是还没卖出
		dp[i][3] = max7(dp[i-1][3], dp[i-1][2]-prices[i])
		// 当天是状态4，那必须是在3的状态下（已经买入了第二次）当天再卖出第二次，或者早在之前就卖出了第二次
		dp[i][4] = max7(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	// 最后从0、第一次卖出、第二次卖出中选取最大值
	return max7(dp[n-1][2], max7(0, dp[n-1][4]))
}

func max7(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

/*
func maxProfit( prices []int ) int {
    // write code here
    if len(prices) < 2 {
        return 0
    }
    var buy1, buy2, profit1, profit2 int
    buy1 = prices[0]
    buy2 = prices[0]
    profit1 = 0
    profit2 = 0
    for _, price := range(prices) {
        buy1 = min(buy1, price)
        profit1 = max(profit1, price - buy1)
        buy2 = min(buy2, price - profit1)
        profit2 = max(profit2, price - buy2)
    }
    return profit2
}
func max(num1, num2 int) int {
        if num1> num2 {
            return num1
        }
        return num2
}

func min(num1, num2 int) int {
        if num1> num2 {
            return num2
        }
        return num1
}

*/
func main() {
	profit := maxProfit3([]int{8, 9, 3, 5, 1, 3})
	fmt.Println(profit)
}



