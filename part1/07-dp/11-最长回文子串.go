package main

import "fmt"

/*
BM73 最长回文子串
https://www.nowcoder.com/practice/b4525d1d84934cf280439aeecc36f4af?tpId=295&tqId=25269&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
对于长度为n的一个字符串A（仅包含数字，大小写英文字母），请设计一个高效算法，计算其中最长回文子串的长度。


数据范围： 1 \le n \le 10001≤n≤1000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n^2)O(n
2
 )
进阶:  空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
示例1
输入：
"ababc"
复制
返回值：
3
复制
说明：
最长的回文子串为"aba"与"bab"，长度都为3
示例2
输入：
"abbba"
复制
返回值：
5
复制
示例3
输入：
"b"
复制
返回值：
1

*/

/*
中心扩散法
判断回文的过程是从首尾到中间，那我们找最长回文串可以逆着来，从中间延伸到首尾，这就是中心扩展法。

step 1：遍历字符串每个字符。
step 2：以每次遍历到的字符为中心（分奇数长度和偶数长度两种情况），不断向两边扩展。
step 3：如果两边都是相同的就是回文，不断扩大到最大长度即是以这个字符（或偶数两个）为中心的最长回文子串。
step 4：我们比较完每个字符为中心的最长回文子串，取最大值即可。
*/

func getLongestPalindrome(A string) int {
	// write code here
	maxLen := 1
	for i := 0; i < len(A); i++ {
		// 偶数
		left := run(A, i, i)
		// 奇数
		right := run(A, i, i+1)
		maxLen = max1(maxLen, max1(left, right))
	}
	return maxLen
}

func run(s string, begin, end int) int {
	// 中心扩散
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

func main11() {
	palindrome := getLongestPalindrome("ababc")
	fmt.Println(palindrome)
}
