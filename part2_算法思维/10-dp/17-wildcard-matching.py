# -*- coding: utf-8 -*-
from typing import List

"""
https://leetcode-cn.com/problems/wildcard-matching/

44. 通配符匹配
给定一个字符串 (s) 和一个字符模式 (p) ，实现一个支持 '?' 和 '*' 的通配符匹配。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。 多个字符串，不限格式

两个字符串完全匹配才算匹配成功。

说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 ? 和 *。

示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
输入:
s = "aa"
p = "*"
输出: true
解释: '*' 可以匹配任意字符串。

示例 3:
输入:
s = "cb"
p = "?a"
输出: false
解释: '?' 可以匹配 'c', 但第二个 'a' 无法匹配 'b'。

示例 4:
输入:
s = "adceb"
p = "*a*b"
输出: true
解释: 第一个 '*' 可以匹配空字符串, 第二个 '*' 可以匹配字符串 "dce".

示例 5:
输入:
s = "acdcb"
p = "a*c?b"
输入: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/wildcard-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""


class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        """

            "" a c d c b
        ""   1 0 0 0 0 0
        a    0 1 0 0 0 0
        *    0 1 1 1 1 1
        c    0 0 1 0 0 0
        ?    0 0 0 1 0 0
        b    0 0 0 0 0 0

        """
        s = " " + s
        p = " " + p

        s_len = len(s)
        p_len = len(p)

        # base case
        dp = [[False] * p_len for _ in range(s_len)]
        dp[0][0] = True

        for i in range(1, p_len):
            # dp[i][0] = dp[i - 1][0] & (p[i-1] == "*")
            dp[0][i] = dp[0][i - 1] and p[i] == "*"

        # print(dp)
        for i in range(1, s_len):
            for j in range(1, p_len):
                # 匹配单个
                if s[i] == p[j] or p[j] == "?":
                    dp[i][j] = dp[i - 1][j - 1]
                # 匹配多个
                # 当前值等于上面或者前面的值
                elif p[j] == "*":
                    dp[i][j] = dp[i - 1][j] or dp[i][j - 1]

        return dp[-1][-1]


s = Solution()
print(s.isMatch(s="acdcb", p="a*c?b"))
print(s.isMatch(s="adceb", p="*a*b"))
print(s.isMatch(s="aa", p="a"))

# 和15题的不同， “*” 的定义不一样
"""
func match2(str string, pattern string) bool {
	str = " " + str
	pattern = " " + pattern
	m := len(str)
	n := len(pattern)

	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}

	dp[0][0] = true
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] && pattern[i] == '*'
	}


	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if pattern[j] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if (str[i] == pattern[j] || pattern[j] == '?' ){
				dp[i][j] = dp[i-1][j-1]
			}
		}

	}

	return dp[m-1][n-1]
}



"""


