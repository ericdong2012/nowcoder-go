"""11. 盛最多水的容器（双指针）
题目：每条竖线高度 heights[i]，与 x 轴构成容器，求最大装水量。
思路：两端指针向中间收。水量 = min(h[l], h[r]) * 宽。
      移动较短的一侧（移动长侧只会更小），每步比较记录最大。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        l, r, best = 0, len(height) - 1, 0
        while l < r:
            best = max(best, min(height[l], height[r]) * (r - l))
            if height[l] < height[r]:
                l += 1
            else:
                r -= 1
        return best


if __name__ == "__main__":
    s = Solution()
    assert s.maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]) == 49
    assert s.maxArea([1, 1]) == 1
    print("05 OK")
