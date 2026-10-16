"""79. 单词搜索（回溯）
题目：网格中按上下左右相邻寻找单词路径，格子不可重复使用。
思路：从每个起点 DFS；匹配当前字符后临时置空（防重复用），
      四方向递归，失败恢复原字符。
复杂度：时间 O(mn·3^L)，空间 O(L)。
"""

from typing import List


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        m, n = len(board), len(board[0])

        def dfs(i: int, j: int, k: int) -> bool:
            if k == len(word):
                return True
            if not (0 <= i < m and 0 <= j < n) or board[i][j] != word[k]:
                return False
            board[i][j] = "#"                 # 占位防重复
            found = (dfs(i + 1, j, k + 1) or dfs(i - 1, j, k + 1)
                     or dfs(i, j + 1, k + 1) or dfs(i, j - 1, k + 1))
            board[i][j] = word[k]             # 恢复
            return found

        return any(dfs(i, j, 0) for i in range(m) for j in range(n))


if __name__ == "__main__":
    s = Solution()
    b = [["A", "B", "C", "E"], ["S", "F", "C", "S"], ["A", "D", "E", "E"]]
    assert s.exist(b, "ABCCED") is True
    assert s.exist(b, "SEE") is True
    assert s.exist(b, "ABCB") is False
    print("60 OK")
