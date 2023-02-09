package main

/*
NC228 判断子序列
https://www.nowcoder.com/practice/39be6c2d0a9b4c30a7b04053d5960a84?tpId=117&tqId=39453&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=


描述
给定两个字符串 S 和 T ，判断 S 是否是 T 的子序列。
即是否可以从 T 删除一些字符转换成 S。

示例1
输入：
"nowcoder","nowcoder"
返回值：
true

示例2
输入：
"nower","nowcoder"
返回值：
true

示例3
输入：
"nowef","nowcoder"
返回值：
false

*/

func isSubsequence(S string, T string) bool {
	// write code here
	m, n := len(S), len(T)
	i, j := 0, 0
	for i < m && j < n {
		if S[i] == T[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == m
}
