"""64. 最小路径和（DP）
题目：网格从左上到右下，路径上数字和最小。
思路：dp[i][j] = grid[i][j] + min(上, 左)；首行首列只能单向来。
复杂度：时间 O(mn)，空间 O(n)（原地改 grid 可 O(1)）。
"""

from typing import List


class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        for j in range(1, n):
            grid[0][j] += grid[0][j - 1]
        for i in range(1, m):
            grid[i][0] += grid[i - 1][0]
            for j in range(1, n):
                grid[i][j] += min(grid[i - 1][j], grid[i][j - 1])
        return grid[-1][-1]


if __name__ == "__main__":
    s = Solution()
    assert s.minPathSum([[1, 3, 1], [1, 5, 1], [4, 2, 1]]) == 7
    assert s.minPathSum([[1, 2, 3], [4, 5, 6]]) == 12
    print("92 OK")
