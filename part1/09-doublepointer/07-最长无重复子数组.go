package main

/*
BM92 最长无重复子数组
https://www.nowcoder.com/practice/b56799ebfd684fb394bd315e89324fb4?tpId=295&tqId=1008889&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个长度为n的数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组

示例1
输入：
[2,3,4,5]
返回值：
4
说明：
[2,3,4,5]是最长子数组

示例2
输入：
[2,2,3,4,3]
返回值：
3
说明：
[2,3,4]是最长子数组

示例4
输入：
[1,2,3,1,2,3,2,2]
返回值：
3
说明：
最长子数组为[1,2,3]
*/

func maxLength(arr []int) int {
	// write code here
	//left, right := 0, 0
	//maxlen := 0
	//temp := make([]int, 1000000)
	//for right <  len(arr) {
	//	// 在temp 中不存在， arr[right] 不重复
	//	if temp[arr[right]] == 0 {
	//		temp[arr[right]]++
	//		//right++
	//		maxlen = max2(maxlen, right-left+1)
	//		right++
	//	} else {
	//		// 有重复的将上一个清零
	//		temp[arr[left]] = 0
	//		left++
	//	}
	//}
	//
	//return maxlen

	left, right := 0, 0
	maxLen := 0
	temp := make([]int, 1000000)
	for left < len(arr) {
		if right == len(arr) {
			return maxLen
		}
		// 在temp 中不存在， arr[right] 不重复
		if temp[arr[right]] == 0 {
			temp[arr[right]]++
			right++
		} else {
			// 有重复的将上一个清零
			temp[arr[left]]--
			left++
		}
		maxLen = max2(maxLen, right-left)
	}

	return maxLen
}

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
