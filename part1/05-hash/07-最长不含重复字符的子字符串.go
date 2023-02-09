package main

import "fmt"

/*
NC170 最长不含重复字符的子字符串
https://www.nowcoder.com/practice/48d2ff79b8564c40a50fa79f9d5fa9c7?tpId=117&tqId=39315&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

描述
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
数据范围:
\ \text{s.length}\le 40000 s.length≤40000
示例1
输入：
"abcabcbb"
返回值：
3

说明：
因为无重复字符的最长子串是"abc"，所以其长度为 3。


示例2
输入：
"bbbbb"
返回值：
1

说明：
因为无重复字符的最长子串是"b"，所以其长度为 1。


示例3
输入：
"pwwkew"
返回值：
3

说明：
因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是子串的长度，"pwke" 是一个子序列，不是子串。

*/

func lengthOfLongestSubstring(s string) int {
	// write code here
	//if len(s) == 1 {
	//	return 1
	//}
	//var count int
	//temp := make(map[rune]bool)
	//for _, v := range s {
	//	if temp[rune(v)] {
	//		count = max(count, len(temp))
	//		temp[rune(v)] = false
	//	} else {
	//		temp[rune(v)] = true
	//	}
	//}
	//
	//if len(temp) != 0 {
	//	count = max(count, len(temp))
	//}
	//return count

	// 双指针
	n := len(s)
	// ascii 常用字符128个
	m := make([]int, 128)
	res := 0
	for l, r := 0, 0; r < n; r++ {
		// l 左指针记录是否有重复元素出现，有的化，往右移动
		l = max(l, m[s[r]])
		// 记录个数，索引之差
		res = max(res, r-l+1)
		// ??? 上面的m[s[r]] 中有值时，l要往下移动一位，所以上次保存的 r + 1
		m[s[r]] = r + 1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	//substring := lengthOfLongestSubstring("abcabcbb")
	substring := lengthOfLongestSubstring("dvdfv")
	//substring := lengthOfLongestSubstring("Kzz8")
	fmt.Println(substring)
}
