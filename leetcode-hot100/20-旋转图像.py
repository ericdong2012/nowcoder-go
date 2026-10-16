"""48. 旋转图像（矩阵旋转）
题目：n×n 矩阵顺时针旋转 90°，原地。
思路：两步法——先沿主对角线转置，再每行左右翻转。
     （逆时针则是转置后上下翻转。）
复杂度：时间 O(n²)，空间 O(1)。
"""

from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        n = len(matrix)
        for i in range(n):                  # 转置
            for j in range(i + 1, n):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        for row in matrix:                  # 每行翻转
            row.reverse()


if __name__ == "__main__":
    s = Solution()
    m = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    s.rotate(m)
    assert m == [[7, 4, 1], [8, 5, 2], [9, 6, 3]]
    m2 = [[1, 2], [3, 4]]
    s.rotate(m2)
    assert m2 == [[3, 1], [4, 2]]
    print("20 OK")
