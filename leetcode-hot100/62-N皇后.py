"""51. N 皇后（回溯 + 集合剪枝）
题目：n×n 棋盘放 n 个皇后互不攻击，返回所有摆法。
思路：逐行放（行天然不冲突）；用三个集合记录已占的列、主对角线(r-c)、
      副对角线(r+c)；冲突剪枝，成功到第 n 行收集方案。
复杂度：时间 O(n!)（带剪枝），空间 O(n)。
"""

from typing import List


class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        res = []
        queens = [-1] * n                     # queens[r] = 该行皇后所在列
        cols, diag1, diag2 = set(), set(), set()

        def backtrack(r: int) -> None:
            if r == n:
                res.append(["." * c + "Q" + "." * (n - c - 1) for c in queens])
                return
            for c in range(n):
                if c in cols or (r - c) in diag1 or (r + c) in diag2:
                    continue
                queens[r] = c
                cols.add(c); diag1.add(r - c); diag2.add(r + c)
                backtrack(r + 1)
                cols.remove(c); diag1.remove(r - c); diag2.remove(r + c)

        backtrack(0)
        return res


if __name__ == "__main__":
    s = Solution()
    out4 = s.solveNQueens(4)
    assert len(out4) == 2
    assert out4[0] == [".Q..", "...Q", "Q...", "..Q."] or out4[0] == ["..Q.", "Q...", "...Q", ".Q.."]
    assert len(s.solveNQueens(1)) == 1
    print("62 OK")
