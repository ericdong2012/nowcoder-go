"""42. 接雨水（双指针/前缀最大值）
题目：柱状图凹槽能接多少水。
思路：每个位置的水量 = min(左边最大高度, 右边最大高度) - 自身高度。
      双指针版：lmax/rmax 分别从两侧维护，谁的最大值小就结算谁的位置
     （小的一侧的瓶颈已确定）。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        l, r = 0, len(height) - 1
        lmax = rmax = 0
        water = 0
        while l < r:
            if height[l] < height[r]:
                lmax = max(lmax, height[l])
                water += lmax - height[l]
                l += 1
            else:
                rmax = max(rmax, height[r])
                water += rmax - height[r]
                r -= 1
        return water


if __name__ == "__main__":
    s = Solution()
    assert s.trap([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]) == 6
    assert s.trap([4, 2, 0, 3, 2, 5]) == 9
    assert s.trap([1, 2, 3]) == 0
    print("07 OK")
