package main

import "fmt"

/*
BM71 最长上升子序列(一)
https://www.nowcoder.com/practice/5164f38b67f846fb8699e9352695cd2f?tpId=295&tqId=2281434&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个长度为 n 的数组 arr，求它的最长严格上升子序列的长度。
所谓子序列，指一个数组删掉一些数（也可以不删）之后，形成的新数组。例如 [1,5,3,7,3] 数组，其子序列有：[1,3,3]、[7] 等。但 [1,6]、[1,3,5] 则不是它的子序列。

示例1
输入：
[6,3,1,5,2,3,7]
返回值：
4

说明：
该数组最长上升子序列为 [1,2,3,7] ，长度为4


*/

//func LIS(arr []int) int {
//	// write code her
//	maxV := 0
//	for i := 0; i < len(arr); i++ {
//		i2 := run2(arr, i)
//		maxV = max4(maxV, i2)
//	}
//	return maxV
//
//}
//
//func run2(arr []int, index int) int {
//	temp := []int{arr[index]}
//	for index < len(arr) {
//		if arr[index] > temp[0] {
//			if arr[index] < temp[len(temp)-1] {
//				temp = temp[:len(temp)-1]
//			}
//			temp = append(temp, arr[index])
//		}
//		index++
//	}
//	return len(temp)
//}

func LIS(arr []int) int {
	// write code her
	maxV := 0
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			// 后面的元素比前面的大， 并且dp 小于等于前一个
			if arr[i] > arr[j] && dp[i] < dp[j]+1 {
				// 后面元素长度 +1
				dp[i] = dp[j] + 1
				// 比较下长度
				maxV = max4(maxV, dp[i])
			}
		}
	}
	return maxV
}

func max4(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main12() {
	//lis := LIS([]int{6, 3, 1, 5, 2, 3, 7})
	lis := LIS([]int{3, 5, 7, 1, 2, 4, 6, 3, 8, 9, 5, 6})
	fmt.Println(lis)
}
