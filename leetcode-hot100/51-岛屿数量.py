"""200. 岛屿数量（DFS 染色）
题目：'1'是陆地，统计岛屿数。
思路：扫到 '1' 计数 +1，DFS 把整座岛改成 '0'（沉岛）防止重复统计。
复杂度：时间 O(mn)，空间 O(mn) 最坏递归栈。
"""

from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        m, n = len(grid), len(grid[0])
        count = 0

        def sink(i: int, j: int) -> None:
            if not (0 <= i < m and 0 <= j < n) or grid[i][j] != "1":
                return
            grid[i][j] = "0"
            sink(i + 1, j)
            sink(i - 1, j)
            sink(i, j + 1)
            sink(i, j - 1)

        for i in range(m):
            for j in range(n):
                if grid[i][j] == "1":
                    count += 1
                    sink(i, j)
        return count


if __name__ == "__main__":
    s = Solution()
    g1 = [["1", "1", "0", "0", "0"], ["1", "1", "0", "0", "0"], ["0", "0", "1", "0", "0"],
          ["0", "0", "0", "1", "1"]]
    assert s.numIslands(g1) == 3
    assert s.numIslands([["1"]]) == 1
    print("51 OK")
