"""118. 杨辉三角
题目：生成 numRows 行杨辉三角。
思路：每行首尾 1，中间 = 上一行相邻两数之和。
复杂度：时间 O(n²)，空间 O(n²)。
"""

from typing import List


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        res = []
        for i in range(numRows):
            row = [1] * (i + 1)
            for j in range(1, i):
                row[j] = res[i - 1][j - 1] + res[i - 1][j]
            res.append(row)
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.generate(5) == [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1], [1, 4, 6, 4, 1]]
    assert s.generate(1) == [[1]]
    print("82 OK")
