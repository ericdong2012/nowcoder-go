package main

/*题目大意

最长公共子序列， 最长公共子串

给定两个字符串 text1 和 text2，返回这两个字符串的最长 公共子序列 的长度。如果不存在 公共子序列 ，返回 0 。一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的 公共子序列 是这两个字符串所共同拥有的子序列。

解题思路
这一题是经典的最长公共子序列的问题。解题思路是二维动态规划。假设字符串 text1 和 text2 的长度分别为 m 和 n，创建 m+1 行 n+1 列的二维数组 dp，定义 dp[i][j] 表示长度为 i 的 text1[0:i-1] 和长度为 j 的 text2[0:j-1] 的最长公共子序列的长度。先考虑边界条件。当 i = 0 时，text1[] 为空字符串，它与任何字符串的最长公共子序列的长度都是 0，所以 dp[0][j] = 0。同理当 j = 0 时，text2[] 为空字符串，它与任何字符串的最长公共子序列的长度都是 0，所以 dp[i][0] = 0。由于二维数组的大小特意增加了 1，即 m+1 和 n+1，并且默认值是 0，所以不需要再初始化赋值了。

当 text1[i−1] = text2[j−1] 时，将这两个相同的字符称为公共字符，考虑 text1[0:i−1] 和 text2[0:j−1] 的最长公共子序列，再增加一个字符（即公共字符）即可得到 text1[0:i] 和 text2[0:j] 的最长公共子序列，所以 dp[i][j]=dp[i−1][j−1]+1。当 text1[i−1] != text2[j−1] 时，最长公共子序列一定在 text[0:i-1], text2[0:j] 和 text[0:i], text2[0:j-1] 中取得。即 dp[i][j] = max(dp[i-1][j], dp[i][j-1])。所以状态转移方程如下：

$$dp[i][j] = \left{\begin{matrix}dp[i-1][j-1]+1 &,text1[i-1]=text2[j-1]\max(dp[i-1][j],dp[i][j-1])&,text1[i-1]\neq text2[j-1]\end{matrix}\right.$$

最终结果存储在 dp[len(text1)][len(text2)] 中。时间复杂度 O(mn)，空间复杂度 O(mn)，其中 m 和 n 分别是 text1 和 text2 的长度。

https://www.nowcoder.com/practice/8cb175b803374e348a614e34b80ae191?tpId=196&tqId=39284&rp=1&ru=/exam/company&qru=/exam/company&sourceUrl=%2Fexam%2Fcompany&difficulty=undefined&judgeStatus=undefined&tags=&title=%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E5%BA%8F%E5%88%97

https://www.nowcoder.com/practice/6d29638c85bb4ffd80c020fe244baf11?tpId=196&tqId=37131&rp=1&ru=/exam/company&qru=/exam/company&sourceUrl=%2Fexam%2Fcompany&difficulty=undefined&judgeStatus=undefined&tags=&title=%E6%9C%80%E9%95%BF%E5%85%AC%E5%85%B1%E5%AD%90%E5%BA%8F%E5%88%97

*/

// dp
func longestCommonSubsequence(text1 string, text2 string) int {
	/*
		  m
			0 a b c d e
		n 0 0 0 0 0 0 0
		  a 0 1 1 1 1 1
		  c 0 1 1 2 2 2
		  e 0 1 1 2 2 3
	*/
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	// 构造了一个二维数组
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			// 如果前一个相等，当前位置的值 等于 对角线的值 + 1; 否则在左边和上边找较大值
			// 看第一个1， 2， 3, 行列各加了一个长度，所以text1[i-1] == text2[j-1] 代表的是 当前字符相等  (a, c, e)  否则会越界
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max11(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	// len(text1) 代表的是最后一个元素
	return dp[len(text1)][len(text2)]
}

func max11(a, b int) int {
	if a > b {
		return a
	}
	return b
}
