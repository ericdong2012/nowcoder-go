package main

/*
BM92 最长无重复子数组
https://www.nowcoder.com/practice/b56799ebfd684fb394bd315e89324fb4?tpId=295&tqId=1008889&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个长度为n的数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组

数据范围：0\le arr.length \le 10^50≤arr.length≤10
5
 ，0 < arr[i] \le 10^50<arr[i]≤10
5

示例1
输入：
[2,3,4,5]
复制
返回值：
4
复制
说明：
[2,3,4,5]是最长子数组
示例2
输入：
[2,2,3,4,3]
复制
返回值：
3
复制
说明：
[2,3,4]是最长子数组
示例3
输入：
[9]
复制
返回值：
1
复制
示例4
输入：
[1,2,3,1,2,3,2,2]
复制
返回值：
3
复制
说明：
最长子数组为[1,2,3]
示例5
输入：
[2,2,3,4,8,99,3]
复制
返回值：
5
复制
说明：
最长子数组为[2,3,4,8,99]

*/

func maxLength(arr []int) int {
	// write code here
	// 双指针
	n := len(arr)
	m := make([]int, 10e5)
	res := 0
	for l, r := 0, 0; r < n; r++ {
		// l 左指针记录是否有重复元素出现，有的化，往右移动到重复元素的索引位置
		l = max2(l, m[arr[r]])
		// 记录个数，索引之差
		res = max2(res, r-l+1)
		// 上面的m[arr[r]] 中有值时，l要往下移动一位，所以保存的是下一次的索引 r + 1
		m[arr[r]] = r + 1
	}
	return res
}

func max2(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
