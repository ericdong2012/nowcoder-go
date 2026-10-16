"""54. 螺旋矩阵（边界收缩）
题目：顺时针螺旋遍历矩阵所有元素。
思路：维护上下左右四个边界，按"右下左上"走一圈，每走完一条边收缩对应边界。
复杂度：时间 O(mn)，空间 O(1)（不算输出）。
"""

from typing import List


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        top, bottom = 0, len(matrix) - 1
        left, right = 0, len(matrix[0]) - 1
        res = []
        while top <= bottom and left <= right:
            for j in range(left, right + 1):       # → 上边
                res.append(matrix[top][j])
            top += 1
            for i in range(top, bottom + 1):       # ↓ 右边
                res.append(matrix[i][right])
            right -= 1
            if top <= bottom:                      # 剩单行时不回头
                for j in range(right, left - 1, -1):
                    res.append(matrix[bottom][j])
                bottom -= 1
            if left <= right:                      # 剩单列时不回头
                for i in range(bottom, top - 1, -1):
                    res.append(matrix[i][left])
                left += 1
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.spiralOrder([[1, 2, 3], [4, 5, 6], [7, 8, 9]]) == [1, 2, 3, 6, 9, 8, 7, 4, 5]
    assert s.spiralOrder([[1, 2, 3, 4]]) == [1, 2, 3, 4]
    assert s.spiralOrder([[1], [2], [3]]) == [1, 2, 3]
    print("19 OK")
