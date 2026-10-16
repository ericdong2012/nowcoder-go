"""74. 搜索二维矩阵（二分）
题目：每行升序且行首大于上一行末尾——整矩阵可视为一维升序数组。
思路：对 m×n 做 0..m*n-1 的整体二分，下标换算 matrix[mid // n][mid % n]。
复杂度：时间 O(log(mn))，空间 O(1)。
"""

from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])
        l, r = 0, m * n - 1
        while l <= r:
            mid = (l + r) // 2
            v = matrix[mid // n][mid % n]
            if v == target:
                return True
            if v < target:
                l = mid + 1
            else:
                r = mid - 1
        return False


if __name__ == "__main__":
    s = Solution()
    m = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
    assert s.searchMatrix(m, 3) is True
    assert s.searchMatrix(m, 13) is False
    print("64 OK")
