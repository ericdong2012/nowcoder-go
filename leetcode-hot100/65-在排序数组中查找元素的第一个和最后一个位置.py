"""34. 在排序数组中查找元素的第一个和最后一个位置（二分边界）
题目：升序数组中 target 的首末下标，不存在返回 [-1,-1]。
思路：两次 lower_bound——first = 第一个 >= target；
      last = 第一个 >= target+1 再减一；验证 first 在范围内且等于 target。
复杂度：时间 O(log n)，空间 O(1)。
"""

from typing import List


def lower_bound(nums: List[int], target: int) -> int:
    """第一个 >= target 的下标。"""
    l, r = 0, len(nums)
    while l < r:
        mid = (l + r) // 2
        if nums[mid] < target:
            l = mid + 1
        else:
            r = mid
    return l


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        first = lower_bound(nums, target)
        if first == len(nums) or nums[first] != target:
            return [-1, -1]
        last = lower_bound(nums, target + 1) - 1
        return [first, last]


if __name__ == "__main__":
    s = Solution()
    assert s.searchRange([5, 7, 7, 8, 8, 10], 8) == [3, 4]
    assert s.searchRange([5, 7, 7, 8, 8, 10], 6) == [-1, -1]
    assert s.searchRange([], 0) == [-1, -1]
    assert s.searchRange([1], 1) == [0, 0]
    print("65 OK")
