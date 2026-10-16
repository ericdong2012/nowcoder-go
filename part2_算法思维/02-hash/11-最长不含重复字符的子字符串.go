package _3_hash

import "fmt"

/*
NC170 最长不含重复字符的子字符串
https://www.nowcoder.com/practice/48d2ff79b8564c40a50fa79f9d5fa9c7?tpId=117&tqId=39315&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

描述
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度.

示例1
输入：
"abcabcbb"
返回值：
3
说明：
因为无重复字符的最长子串是"abc"，所以其长度为3.


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



/*
连续
不重复
最长

前缀和 + hash + 双指针
*/
func lengthOfLongestSubstring(s string) int {
	// 用左右分别来统计存在和不存在
	left, right := 0, 0
	maxlen := 0
	temp := make(map[byte]int, 1000000)
	for left < len(s) {
		// 已经走到了最后，防止索引越界
		if right == len(s) {
			return maxlen
		}
		// arr[right] 在 temp 中不存在
		if temp[s[right]] == 0 {
			temp[s[right]]++
			right++
		} else {
			// arr[right] 在temp中存在
			// 有重复的将上一个置零(从头部开始置0)
			temp[s[left]] = 0
			left++
		}
		maxlen = max(maxlen, right-left)
	}

	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	//substring := lengthOfLongestSubstring("abcabcbb")
	substring := lengthOfLongestSubstring("dvdfv")
	//substring := lengthOfLongestSubstring("Kzz8")
	fmt.Println(substring)
}
