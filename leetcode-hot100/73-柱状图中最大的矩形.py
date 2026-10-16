"""84. 柱状图中最大的矩形（单调栈 + 哨兵）
题目：直方图中能组成的最大矩形面积。
思路：对每根柱子求"两侧第一个更矮的位置"——单调递增栈。
      哨兵技巧：末尾加 0 高度，强制栈清空，统一结算。
      面积 = 高度 × 宽度（弹栈时确定）。
复杂度：时间 O(n)，空间 O(n)。
"""

from typing import List


class Solution:
    def largestRectangleArea(self, heights: List[int]) -> int:
        stack = [-1]                    # 哨兵：栈底边界
        best = 0
        for i, h in enumerate(heights + [0]):   # 末尾哨兵柱
            while stack[-1] != -1 and heights[stack[-1]] >= h:
                height = heights[stack.pop()]
                width = i - stack[-1] - 1       # 左边界是新的栈顶
                best = max(best, height * width)
            stack.append(i)
        return best


if __name__ == "__main__":
    s = Solution()
    assert s.largestRectangleArea([2, 1, 5, 6, 2, 3]) == 10
    assert s.largestRectangleArea([2, 4]) == 4
    assert s.largestRectangleArea([1]) == 1
    print("73 OK")
