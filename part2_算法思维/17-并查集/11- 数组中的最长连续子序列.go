package main

import (
	"sort"
)

/*
NC95 数组中的最长连续子序列
https://www.nowcoder.com/practice/eac1c953170243338f941959146ac4bf?tpId=117&tqId=37810&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196&difficulty=undefined&judgeStatus=undefined&tags=&title=

描述
给定无序数组arr，返回其中最长的连续序列的长度(要求值连续，位置可以不连续,例如 3,4,5,6为连续的自然数）

示例1
输入：
[100,4,200,1,3,2]
返回值：
4

示例2
输入：
[1,1,1]
返回值：
1
*/


/*
题目理解： 连续子序列，必须是连续，没有要求是升序，可以是乱序， 比如： 1，2，3，4 或者 4，3，2，1 都是满足要求，
思路：    先排序，后遍历，找到后一个比前一个大1， 计数加1， 否则置1; 边界情况： 前一个和当前相等， 跳过
*/

func MLS( arr []int ) int {
	// write code here
	sort.Ints(arr)
	count := 1
	res :=1
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[i-1] {
			continue
		}
		if arr[i] -  arr[i-1] == 1 {
			count ++
		}else {
			count = 1
		}
		res = max(res, count)
	}

	return res

}

func max(a ,b int) int {
	if a < b {
		return b
	}

	return a
}



//func MLS( arr []int ) int {
//	// write code here
//	numSet := make(map[int]bool,len(arr))
//	for i := range arr {
//		numSet[arr[i]] = true
//	}
//	longestSeq := 0
//	for num,_ := range numSet {
//		if !numSet[num-1] {
//			currSeq := 1
//			for ;numSet[num+1];currSeq++{num++}
//			if currSeq > longestSeq {
//				longestSeq = currSeq
//			}
//		}
//	}
//	return longestSeq
//}