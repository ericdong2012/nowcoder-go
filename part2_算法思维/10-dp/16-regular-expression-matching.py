# -*- coding: utf-8 -*-
from typing import List
import re

"""
https://leetcode-cn.com/problems/regular-expression-matching/

10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p， 请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素

所谓匹配，是要涵盖整个字符串，而不是部分字符串。

说明:
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

示例 2:
输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。

示例 3:
输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。

示例 4:
输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。

示例 5:
输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""


class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        # https://leetcode-cn.com/problems/regular-expression-matching/solution/ji-yu-guan-fang-ti-jie-gen-xiang-xi-de-jiang-jie-b/
        # method1: 递归
        # if not p: return not s
        #
        # first = bool(s) and p[0] in {s[0], '.'}
        # # 星号前面的那个字符到底要重复几次呢？这需要计算机暴力穷举来算，假设重复N次吧。前文多次强调过，写递归的技巧是管好当下，之后的事抛给递归。
        # # 具体到这里，不管N是多少，当前的选择只有两个：匹配0次、匹配1次。所以可以这样处理：
        #
        #
        # if len(p) >= 2 and p[1] == '*':
        # # 解释：如果发现有字符和 '*' 结合，
        # # 或者匹配该字符 0 次，然后跳过该字符和 '*'
        # # 或者当 pattern[0] 和 text[0] 匹配后，移动 text
        #     return self.isMatch(s, p[2:]) or \
        #            first and self.isMatch(s[1:], p)
        # else:
        #     return first and self.isMatch(s[1:], p[1:])

        # method2: 带备忘录的动态规划
        """
        我选择使用「备忘录」递归的方法来降低复杂度。有了暴力解法，优化的过程及其简单，就是使用两个变量 i, j 记录当前匹配到的位置，从而避免使用子字符串切片，并且将 i, j 存入备忘录，避免重复计算即可。
        我将暴力解法和优化解法放在一起，方便你对比，你可以发现优化解法无非就是把暴力解法「翻译」了一遍，加了个 memo 作为备忘录，仅此而已

            m i s s i s s i p p i
        m   1 0 0 0 0 0 0 0 0 0 0
        i   0 1 0 0 0 0 0 0 0 0 0
        s   0 0 1 0 0 0 0 0 0 0 0
        *   0 0 1 1 1 1 1 1 1 1 1
        i   0 0 0 0 1 0 0 0 0 0 0
        s
        *
        p
        *
        .

        """
        memo = dict()  # 备忘录

        def dp(i, j):
            # 直接从memo中取，优化了代码
            if (i, j) in memo: return memo[(i, j)]

            # p是带通配符的那个字符串
            # 如果j走到了最后，此时i也应该走了s的最后
            if j == len(p): return i == len(s)

            # first 是个bool
            first = i < len(s) and p[j] in {s[i], '.'}

            # 匹配 *
            # 如果j未到p中倒数第二个字符，并且p的下一个字符是*
            # 结果等于纵向往下走2个字符 或者  横向走一个字符并且当前字符在s[i]中 或者 是 .
            # 比如: dp(2,2)
            if j <= len(p) - 2 and p[j + 1] == '*':
                # 通配符'*'匹配了0次， return dp(i, j + 2)
                # 通配符'*'匹配了多次，return dp(i + 1, j)
                ans = dp(i, j + 2) or first and dp(i + 1, j)
            # 匹配. 或者 其他字符, 比如: dp(1, 1), dp(2, 2)
            else:
                ans = first and dp(i + 1, j + 1)

            # 将坐标和结果 true / false 存放进去
            memo[(i, j)] = ans
            return ans

        return dp(0, 0)

s = Solution()
print(s.isMatch(s="mississippi", p="mis*is*p*."))
print(s.isMatch(s="aab", p="c*a*b"))
