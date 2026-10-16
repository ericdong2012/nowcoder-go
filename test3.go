package main

import "fmt"

func minMoney(arr []int, aim int) int {
	// write code here
	if aim < 1 {
		return 0
	}
	// 建立dp, 存的是最终的张数，索引是从小到aim， 给首位赋值0， 其他位aim + 1
	dp := make([]int, aim+1)
	dp[0] = 0
	for i := 1; i <= aim; i++ {
		dp[i] = aim + 1
	}
	// dp 要凑的钱, 最小是1， 最大是aim
	for i := 1; i < aim+1; i++ {
		// arr 面值(每种面值的货币都要枚举)
		for j := 0; j < len(arr); j++ {
			//在相等的时候, dp[i-arr[j]] = 0, 而张数最小为1, 所以在最后加上1
			if arr[j] <= i {
				fmt.Println(i, arr[j], dp[i], dp[i-arr[j]] + 1 )
				/*
				5 5 21 1
				5 2 1 2
				5 3 1 2
				*/
				dp[i] = min2(dp[i], dp[i-arr[j]]+1)
			}
		}
	}
	// 最终，如果dp倒数第二位大于aim, 可能是每个元素都大于aim.
	if dp[aim] > aim {
		return -1
	}

	return dp[aim]
}

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	money := minMoney([]int{5, 2, 3}, 20)
	//money := minMoney([]int{5, 2, 3}, 11)  // 3
	fmt.Println(money)
}
