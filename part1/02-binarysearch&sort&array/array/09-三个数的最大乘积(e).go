package main

import "sort"

/*
NC106 三个数的最大乘积
https://www.nowcoder.com/practice/8ae05c2913fe438b8b14f3968f64fc0b?tpId=117&tqId=37830&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=


描述
给定一个长度为 n 的无序数组 A ，包含正数、负数和 0 ，请从中找出 3 个数，使得乘积最大，返回这个乘积。
要求时间复杂度： O(n) ，空间复杂度： O(1) 。


示例1
输入：
[3,4,1,2]
返回值：
24

*/

func solve(A []int) int64 {
	// write code here
	// 排序
	// 两绝对值最大的两个负+ 最大的正 或者 3个最大的正数
	sort.Ints(A)
	max1:=A[len(A)-1]*A[len(A)-2]*A[len(A)-3]
	max2:=A[0]*A[1]*A[len(A)-1]
	if max2 > max1 {
		return int64(max2)
	}else {
		return int64(max1)
	}

}
