"""139. 单词拆分（DP）
题目：字符串能否由字典中的单词拼接而成（可重复使用）。
思路：dp[i] = 前 i 个字符可拆分；转移：存在 j<i 使 dp[j] 且 s[j:i] 在字典。
复杂度：时间 O(n²·L)（L 为串内比较开销），空间 O(n)。
"""

from typing import List


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        words = set(wordDict)
        n = len(s)
        dp = [True] + [False] * n        # dp[0]=空串 True
        for i in range(1, n + 1):
            for j in range(i):
                if dp[j] and s[j:i] in words:
                    dp[i] = True
                    break
        return dp[n]


if __name__ == "__main__":
    s = Solution()
    assert s.wordBreak("leetcode", ["leet", "code"]) is True
    assert s.wordBreak("applepenapple", ["apple", "pen"]) is True
    assert s.wordBreak("catsandog", ["cats", "dog", "sand", "and", "cat"]) is False
    print("86 OK")
