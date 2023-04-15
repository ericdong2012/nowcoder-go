package main

/*
BM88 判断是否为回文字符串
https://www.nowcoder.com/practice/e297fdd8e9f543059b0b5f05f3a7f3b2?tpId=295&tqId=1089616&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个长度为 n 的字符串，请编写一个函数判断该字符串是否回文。如果是回文请返回true，否则返回false。

字符串回文指该字符串正序与其逆序逐字符一致。

数据范围：0 < n \le 10000000<n≤1000000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
"absba"
复制
返回值：
true
复制
示例2
输入：
"ranko"
复制
返回值：
false
复制
示例3
输入：
"yamatomaya"
复制
返回值：
false
复制
示例4
输入：
"a"
复制
返回值：
true
复制
备注：
字符串长度不大于1000000，且仅由小写字母组成


*/

func judge( str string ) bool {
	// write code here
	left, right :=0, len(str) -1
	for left <=right {
		if str[left] != str[right] {
			return false
		}
		left ++
		right--
	}
	return true
}
