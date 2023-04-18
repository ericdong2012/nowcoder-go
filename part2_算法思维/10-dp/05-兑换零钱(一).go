package main

import "fmt"

/*
BM70 兑换零钱(一)
https://www.nowcoder.com/practice/3911a20b3f8743058214ceaa099eeb45?tpId=295&tqId=988994&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定数组arr，arr中所有的值都为正整数且不重复。 每个值代表一种面值的货币，每种面值的货币可以使用任意张，再给定一个aim，代表要找的钱数，求组成aim的最少货币数。
如果无解，请返回-1.


示例1
输入：
[5,2,3],20
返回值：
4

示例2
输入：
[5,2,3],0
返回值：
0

示例3
输入：
[3,5],2
返回值：
-1

*/



/*
三数之和的升级版
背包问题简化版
*/
func minMoney(arr []int, aim int) int {
	// write code here
	if aim < 1 {
		return 0
	}
	// 建立dp, 给首位赋值0， 其他位aim + 1
	dp := make([]int, aim+1)
	dp[0] = 0
	for i := 1; i <= aim; i++ {
		dp[i] = aim + 1
	}
	// dp 要凑的钱
	for i := 1; i < aim+1; i++ {
		// arr 面值(每种面值的货币都要枚举)
		for j := 0; j < len(arr); j++ {
			// 如果面值 小于 要凑的钱, 在相等的时候，dp[i-arr[j]] = 0 ， 而面值最小为1元，所以在最后加上1
			if arr[j] <= i {
				dp[i] = min2(dp[i], dp[i-arr[j]] + 1)
			}
		}
	}
	// 最终，如果dp倒数第二位大于aim, 超过了aim（可能是每个元素都大于aim）
	if dp[aim] > aim {
		return -1
	} else {
		return dp[aim]
	}

}

func min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main05() {
	money := minMoney([]int{5, 2, 3}, 20)
	fmt.Println(money)
}
