package main

/*
NC17 最长回文子串
https://www.nowcoder.com/practice/b4525d1d84934cf280439aeecc36f4af?tpId=117&tqId=37789&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=


描述
对于长度为n的一个字符串A（仅包含数字，大小写英文字母），请设计一个高效算法，计算其中最长回文子串的长度。
要求: 空间复杂度 O(1)，时间复杂度 O(n^2)
进阶: 空间复杂度 O(n)，时间复杂度 O(n)

示例1
输入:
"ababc"
返回值：
3
说明：
最长的回文子串为"aba"与"bab"，长度都为3

示例2
输入:
"abbba"
返回值：
5

示例3
输入:
"b"
返回值：
1

*/

/*
中心扩散法

回文串，有着左右对称的特征，从首尾一起访问，遇到的元素都是相同的。
但是我们这里正是要找最长的回文串，并不事先知道长度，怎么办？判断回文的过程是从首尾到中间，
那我们找最长回文串可以逆着来，从中间延伸到首尾，这就是中心扩展法

step 1：遍历字符串每个字符。
step 2：以每次遍历到的字符为中心（分奇数长度和偶数长度两种情况），不断向两边扩展。
step 3：如果两边都是相同的就是回文，不断扩大到最大长度即是以这个字符（或偶数两个）为中心的最长回文子串。
step 4：我们比较完每个字符为中心的最长回文子串，取最大值即可。

*/

// 中心扩散法
func getLongestPalindrome1(A string) int {
	// "baabccc"
	// 看运行过程 最后一个元素，倒数第二个元素...
	maxLen := 1
	for i := len(A) - 1; i >= 0; i-- {
		// 两种情况：可能是类似 aba 的字符串，也可能是类似 abba 的情况
		// 偶数
		left := run(A, i, i)
		// 奇数
		right := run(A, i, i+1)
		maxLen = max1(maxLen, max1(left, right))
	}
	return maxLen
}

func run(s string, begin, end int) int {
	for begin >= 0 && end < len(s) && s[begin] == s[end] {
		begin--
		end++
	}
	return end - begin - 1
}

func max1(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
