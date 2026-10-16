"""994. 腐烂的橘子（多源 BFS）
题目：每分钟烂橘子感染四邻新鲜橘子，求全部腐烂的最少分钟数；不可能返回 -1。
思路：多源 BFS——所有烂橘子同时入队作为第 0 层，逐层扩散；
      最后检查是否还有新鲜的。
复杂度：时间 O(mn)，空间 O(mn)。
"""

from typing import List
from collections import deque


class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        q = deque()
        fresh = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 2:
                    q.append((i, j))
                elif grid[i][j] == 1:
                    fresh += 1
        minutes = 0
        while q and fresh:
            minutes += 1
            for _ in range(len(q)):
                i, j = q.popleft()
                for di, dj in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                    ni, nj = i + di, j + dj
                    if 0 <= ni < m and 0 <= nj < n and grid[ni][nj] == 1:
                        grid[ni][nj] = 2
                        fresh -= 1
                        q.append((ni, nj))
        return minutes if fresh == 0 else -1


if __name__ == "__main__":
    s = Solution()
    assert s.orangesRotting([[2, 1, 1], [1, 1, 0], [0, 1, 1]]) == 4
    assert s.orangesRotting([[2, 1, 1], [0, 1, 1], [1, 0, 1]]) == -1
    assert s.orangesRotting([[0, 2]]) == 0
    print("52 OK")
