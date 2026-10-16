"""73. 矩阵置零
题目：元素为 0 的行列全置零，原地 O(1) 额外空间。
思路：用第一行/第一列做"标记位"，另用两个变量记录第一行列本身是否含 0。
      先扫全矩阵打标记，再据标记置零，最后处理第一行/列。
复杂度：时间 O(mn)，空间 O(1)。
"""

from typing import List


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        m, n = len(matrix), len(matrix[0])
        first_row = any(matrix[0][j] == 0 for j in range(n))
        first_col = any(matrix[i][0] == 0 for i in range(m))

        for i in range(1, m):                    # 打标记
            for j in range(1, n):
                if matrix[i][j] == 0:
                    matrix[i][0] = matrix[0][j] = 0

        for i in range(1, m):                    # 据标记置零
            for j in range(1, n):
                if matrix[i][0] == 0 or matrix[0][j] == 0:
                    matrix[i][j] = 0

        if first_row:
            for j in range(n):
                matrix[0][j] = 0
        if first_col:
            for i in range(m):
                matrix[i][0] = 0


if __name__ == "__main__":
    s = Solution()
    m1 = [[1, 1, 1], [1, 0, 1], [1, 1, 1]]
    s.setZeroes(m1)
    assert m1 == [[1, 0, 1], [0, 0, 0], [1, 0, 1]]
    m2 = [[0, 1]]
    s.setZeroes(m2)
    assert m2 == [[0, 0]]
    print("18 OK")
