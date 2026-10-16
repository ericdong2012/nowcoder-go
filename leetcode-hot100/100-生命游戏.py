"""289. 生命游戏（原地标记）
题目：按四条规则更新棋盘一代，要求原地（不用 O(mn) 额外空间）。
思路：用复合状态编码——
      -1: 曾活今死   2: 曾死今活
      遍历邻居时 val in (1, -1) 视为"本来活着"；第二遍按新值写回 0/1。
复杂度：时间 O(mn)，空间 O(1)。
"""

from typing import List


class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:
        m, n = len(board), len(board[0])
        dirs = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]

        def alive(i, j) -> bool:
            if not (0 <= i < m and 0 <= j < n):
                return False
            return board[i][j] in (1, -1)      # -1 表示曾活今死

        for i in range(m):
            for j in range(n):
                neighbors = sum(alive(i + di, j + dj) for di, dj in dirs)
                if board[i][j] == 1 and neighbors in (2, 3):
                    continue                    # 活着保持活（暂时仍是 1）
                if board[i][j] == 1:
                    board[i][j] = -1            # 活 → 死
                elif board[i][j] == 0 and neighbors == 3:
                    board[i][j] = 2             # 死 → 活

        for i in range(m):
            for j in range(n):
                board[i][j] = 1 if board[i][j] in (1, 2) else 0


if __name__ == "__main__":
    s = Solution()
    b = [[0, 1, 0], [0, 0, 1], [1, 1, 1], [0, 0, 0]]
    s.gameOfLife(b)
    assert b == [[0, 0, 0], [1, 0, 1], [0, 1, 1], [0, 1, 0]]
    print("100 OK")
