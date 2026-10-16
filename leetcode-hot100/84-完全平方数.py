"""279. 完全平方数（DP）
题目：和为 n 的最少完全平方数个数。
思路：完全背包——dp[i] = min(dp[i - s*s] + 1)，s*s ≤ i。
复杂度：时间 O(n·√n)，空间 O(n)。
"""


class Solution:
    def numSquares(self, n: int) -> int:
        dp = [0] + [float("inf")] * n
        for i in range(1, n + 1):
            s = 1
            while s * s <= i:
                dp[i] = min(dp[i], dp[i - s * s] + 1)
                s += 1
        return dp[n]


if __name__ == "__main__":
    s = Solution()
    assert s.numSquares(12) == 3      # 4+4+4
    assert s.numSquares(13) == 2      # 4+9
    assert s.numSquares(1) == 1
    print("84 OK")
