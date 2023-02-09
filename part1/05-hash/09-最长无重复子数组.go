package main

import "fmt"

/*
NC41 最长无重复子数组
https://www.nowcoder.com/practice/b56799ebfd684fb394bd315e89324fb4?tpId=117&tqId=37816&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

描述
给定一个长度为n的数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组

数据范围：0\le arr.length \le 10^5

*/

func maxLength(arr []int) int {
	// 双指针
	n := len(arr)
	m := make([]int, 10e5)
	res := 0
	for l, r := 0, 0; r < n; r++ {
		// l 左指针记录是否有重复元素出现，有的化，往右移动
		l = max1(l, m[arr[r]])
		// 记录个数，索引之差
		res = max1(res, r-l+1)
		// ??? 上面的m[s[r]] 中有值时，l要往下移动一位，所以上次保存的 r + 1
		m[arr[r]] = r + 1
	}
	return res
}

func max1(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	length := maxLength([]int{2, 3, 4, 5})
	fmt.Println(length)
}
