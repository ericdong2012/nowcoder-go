"""153. 寻找旋转排序数组中的最小值（二分）
题目：旋转过的升序数组中找最小值。
思路：与右端点比较——nums[mid] > nums[r] 说明最小在右半；
      否则最小在含 mid 的左半。收敛到 l。
复杂度：时间 O(log n)，空间 O(1)。
"""

from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        l, r = 0, len(nums) - 1
        while l < r:
            mid = (l + r) // 2
            if nums[mid] > nums[r]:       # 右半"断了"，最小在 mid 右侧
                l = mid + 1
            else:
                r = mid                   # 最小可能是 mid 本身
        return nums[l]


if __name__ == "__main__":
    s = Solution()
    assert s.findMin([3, 4, 5, 1, 2]) == 1
    assert s.findMin([4, 5, 6, 7, 0, 1, 2]) == 0
    assert s.findMin([11, 13, 15, 17]) == 11
    assert s.findMin([2, 1]) == 1
    print("67 OK")
