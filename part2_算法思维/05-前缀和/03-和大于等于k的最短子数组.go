package _5_前缀和

import "math"

/*
NC343 和大于等于K的最短子数组
https://www.nowcoder.com/practice/3e1fd3d19fb0479d94652d49c7e1ead1?tpId=196&tqId=40409&ru=/exam/oj


知识点
二分
前缀和
队列
堆
双指针


描述
给定一个长度为 n 的整数数组，和一个目标值 k ，请你找出这个整数数组中和大于等于 k 的最短子数组的长度。如果不存在和大于等于 k 的子数组则输出 -1。

示例1
输入：
[2,1,2,3],5
返回值：
2

示例2
输入：
[2,3,4,5],16
返回值：
-1

*/

func shortestSubarray(nums []int, target int) int {
	l, h := 0, 0
	sum := 0
	length := math.MaxInt32
	for h < len(nums) {
		sum += nums[h]
		for sum >= target {
			length = min2(length, h-l+1)
			sum -= nums[l]
			l++
		}
		h++
	}
	if l != 0 {
		return length
	} else {
		return -1
	}
}
func min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}
