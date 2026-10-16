"""35. 搜索插入位置（二分）
题目：升序数组找 target，找不到返回应插入位置。
思路：二分找第一个 >= target 的下标（lower_bound）。
复杂度：时间 O(log n)，空间 O(1)。
"""

from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums)                   # 左闭右开
        while l < r:
            mid = (l + r) // 2
            if nums[mid] < target:
                l = mid + 1
            else:
                r = mid
        return l


if __name__ == "__main__":
    s = Solution()
    assert s.searchInsert([1, 3, 5, 6], 5) == 2
    assert s.searchInsert([1, 3, 5, 6], 2) == 1
    assert s.searchInsert([1, 3, 5, 6], 7) == 4
    assert s.searchInsert([1, 3, 5, 6], 0) == 0
    print("63 OK")
