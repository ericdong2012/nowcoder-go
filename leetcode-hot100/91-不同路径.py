"""62. 不同路径（DP / 组合数）
题目：m×n 网格从左上到右下（只能右/下）的路径数。
思路：dp[i][j] = dp[i-1][j] + dp[i][j-1]，第一行列全 1；
      或组合数 C(m+n-2, m-1)。
复杂度：DP 时间 O(mn)，空间 O(n)（滚动）。
"""


class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        row = [1] * n
        for _ in range(1, m):
            for j in range(1, n):
                row[j] += row[j - 1]
        return row[-1]


if __name__ == "__main__":
    s = Solution()
    assert s.uniquePaths(3, 7) == 28
    assert s.uniquePaths(3, 2) == 3
    assert s.uniquePaths(1, 1) == 1
    print("91 OK")
