"""
题目大意

最长公共子序列，最长公共子串

给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。如果不存在公共子序列，返回 0。
一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符
（也可以不删除任何字符）后组成的新字符串。例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
两个字符串的公共子序列是这两个字符串所共同拥有的子序列。

解题思路
这一题是经典的最长公共子序列的问题。解题思路是二维动态规划。假设字符串 text1 和 text2 的长度分别为 m 和 n，
创建 m+1 行 n+1 列的二维数组 dp，定义 dp[i][j] 表示长度为 i 的 text1[0:i-1] 和长度为 j 的 text2[0:j-1] 的
最长公共子序列的长度。

当 text1[i-1] = text2[j-1] 时，dp[i][j] = dp[i-1][j-1] + 1。
当 text1[i-1] != text2[j-1] 时，dp[i][j] = max(dp[i-1][j], dp[i][j-1])。

最终结果存储在 dp[len(text1)][len(text2)] 中。时间复杂度 O(mn)，空间复杂度 O(mn)。

https://www.nowcoder.com/practice/6d29638c85bb4ffd80c020fe244baf11
"""


def _max11(a, b):
    if a > b:
        return a
    return b


# dp
def longest_common_subsequence(text1, text2):
    """
      m
        0 a b c d e
    n 0 0 0 0 0 0 0
      a 0 1 1 1 1 1
      c 0 1 1 2 2 2
      e 0 1 1 2 2 3
    """
    if len(text1) == 0 or len(text2) == 0:
        return 0
    # 构造了一个二维数组
    dp = [[0] * (len(text2) + 1) for _ in range(len(text1) + 1)]
    for i in range(1, len(text1) + 1):
        for j in range(1, len(text2) + 1):
            # 如果前一个相等，当前位置的值等于对角线的值 + 1；否则在左边和上边找较大值
            # 看第一个 1，2，3，行列各加了一个长度，所以 text1[i-1] == text2[j-1] 代表的是当前字符相等 (a, c, e)，否则会越界
            if text1[i - 1] == text2[j - 1]:
                dp[i][j] = dp[i - 1][j - 1] + 1
            else:
                dp[i][j] = _max11(dp[i][j - 1], dp[i - 1][j])
    # len(text1) 代表的是最后一个元素
    return dp[len(text1)][len(text2)]
