"""72. 编辑距离（经典二维 DP）
题目：word1 变 word2 的最少操作数（插入/删除/替换一个字符）。
思路：dp[i][j] = w1 前 i 个 → w2 前 j 个的最少操作。
      w1[i-1]==w2[j-1]: dp[i][j]=dp[i-1][j-1]；
      否则 1 + min(替换 dp[i-1][j-1], 删除 dp[i-1][j], 插入 dp[i][j-1])。
复杂度：时间 O(mn)，空间 O(mn)（可滚动到 O(n)）。
"""


class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        m, n = len(word1), len(word2)
        dp = [[0] * (n + 1) for _ in range(m + 1)]
        for i in range(m + 1):              # 边界：变成空串需删光
            dp[i][0] = i
        for j in range(n + 1):
            dp[0][j] = j

        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if word1[i - 1] == word2[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1]
                else:
                    dp[i][j] = 1 + min(dp[i - 1][j - 1],   # 替换
                                       dp[i - 1][j],       # 删除 w1 字符
                                       dp[i][j - 1])       # 插入
        return dp[m][n]


if __name__ == "__main__":
    s = Solution()
    assert s.minDistance("horse", "ros") == 3
    assert s.minDistance("intention", "execution") == 5
    assert s.minDistance("", "abc") == 3
    print("93 OK")
