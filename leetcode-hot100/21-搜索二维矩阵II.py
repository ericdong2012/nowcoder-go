"""240. 搜索二维矩阵 II（BST 式查找）
题目：每行从左到右升序、每列从上到下升序，判断 target 是否存在。
思路：从右上角出发——当前值大于 target 左移，小于 target 下移。
      每步排除一行或一列，像在 BST 上走。
复杂度：时间 O(m+n)，空间 O(1)。
"""

from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        i, j = 0, len(matrix[0]) - 1        # 右上角
        while i < len(matrix) and j >= 0:
            v = matrix[i][j]
            if v == target:
                return True
            if v > target:
                j -= 1                      # 这一列剩余的都更大，排除
            else:
                i += 1                      # 这一行剩余的都更小，排除
        return False


if __name__ == "__main__":
    s = Solution()
    m = [[1, 4, 7, 11, 15], [2, 5, 8, 12, 19], [3, 6, 9, 16, 22],
         [10, 13, 14, 17, 24], [18, 21, 23, 26, 30]]
    assert s.searchMatrix(m, 5) is True
    assert s.searchMatrix(m, 20) is False
    print("21 OK")
