"""
正则表达式匹配
https://www.nowcoder.com/practice/28970c15befb4ff3a264189087b99ad4?tpId=196&tqId=37114

描述
请实现一个函数用来匹配包括'.'和'*'的正则表达式
1.模式中的字符'.'表示任意一个字符
2.模式中的字符'*'表示它前面的字符可以出现任意次（包含0次）
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，
但是与"aa.a"和"ab*a"均不匹配

数据范围:
1.str 只包含从 a-z 的小写字母
2.pattern 只包含从 a-z 的小写字母以及字符 . 和 *，  无连续的 '*'

示例1
输入："aaa", "a*a"
返回值：true

示例2
输入："aad", "c*a*d"
返回值：true

示例3
输入："a", ".*"
返回值：true
"""

# 动态规划
# 设 s 的长度为 m ， p 的长度为 n ；将 s 的第 i 个字符记为 si，p 的第 j 个字符记为 pj，
# 将 s 的前 i 个字符组成的子字符串记为 s[:i] , 同理将 p 的前 j 个字符组成的子字符串记为 p[:j] 。
# 因此，本题可转化为求 s[:m] 是否能和 p[:n] 匹配。
#
#    0 c * a * d
#  0 1 0 1 0 1 0
#  a 0 0 0 1 1 0
#  a 0 0 0 0 1 0
#  d 0 0 0 0 0 1
#
# 如果 pattern[j] == str[i] || pattern[j] == '.', 此时dp[i][j] = dp[i-1][j-1];
# 如果 pattern[j] == '*'
#     分两种情况:
#     1: 如果pattern[j-1] != str[i] && pattern[j-1] != '.', 此时dp[i][j] = dp[i][j-2] //a*匹配0次
#     2: 如果pattern[j-1] == str[i] || pattern[j-1] == '.'
#         此时 dp[i][j] = dp[i][j-2] // a*匹配0次
#         或者 dp[i][j] = dp[i][j-1] // a*匹配1次
#         或者 dp[i][j] = dp[i-1][j] // a*匹配多次


# 解法1
# def match(s, p):
#     # write code here
#     m, n = len(s), len(p)
#     # dp[i][j] 代表字符串 s 的前 i 个字符和 p 的前 j 个字符能否匹配。
#     dp = [[False] * (n + 1) for _ in range(m + 1)]
#     dp[0][0] = True  # 代表两个空字符串能够匹配。
#     # 初始化首行
#     # 首行 s 为空字符串，因此当 p 的偶数位为 * 时才能够匹配
#     for j in range(2, n + 1, 2):
#         dp[0][j] = dp[0][j-2] and p[j-1] == '*'
#
#     for i in range(1, m + 1):
#         for j in range(1, n + 1):
#             if p[j-1] == '*':
#                 if dp[i][j-2]:
#                     dp[i][j] = True
#                 elif dp[i-1][j] and s[i-1] == p[j-2]:
#                     dp[i][j] = True
#                 elif dp[i-1][j] and p[j-2] == '.':
#                     dp[i][j] = True
#             else:
#                 if dp[i-1][j-1] and s[i-1] == p[j-1]:
#                     dp[i][j] = True
#                 elif dp[i-1][j-1] and p[j-1] == '.':
#                     dp[i][j] = True
#     return dp[m][n]


# 解法2
def match2(str_val, pattern):
    m = len(str_val)
    n = len(pattern)
    dp = [[False] * (n + 1) for _ in range(m + 1)]

    def same(i, j):
        if i < 0 or j < 0:
            return False
        if pattern[j] == '.':
            return True
        if pattern[j] == '*':
            return False
        return str_val[i] == pattern[j]

    dp[0][0] = True
    for i in range(m + 1):
        for j in range(1, n + 1):
            # 出现多次
            if pattern[j - 1] == '*':
                # 默认
                dp[i][j] = dp[i][j] or dp[i][j - 2]
                # 在前一列, 有 . 或者 字符串相等
                if same(i - 1, j - 2):
                    dp[i][j] = dp[i][j] or dp[i - 1][j] or dp[i][j - 1]
            elif same(i - 1, j - 1):  # 有. 或者 字符串相等
                dp[i][j] = dp[i][j] or dp[i - 1][j - 1]
    return dp[m][n]
