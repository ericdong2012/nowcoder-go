"""33. 搜索旋转排序数组（二分，O(log n)）
题目：升序数组在未知点旋转过，找 target 下标；无则 -1。
思路：每次二分后，左右两半必有一半有序——判断哪半有序，
      target 在有序半段的范围内就去那半，否则去另一半。
复杂度：时间 O(log n)，空间 O(1)。
"""

from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums) - 1
        while l <= r:
            mid = (l + r) // 2
            if nums[mid] == target:
                return mid
            if nums[l] <= nums[mid]:              # 左半有序
                if nums[l] <= target < nums[mid]:
                    r = mid - 1
                else:
                    l = mid + 1
            else:                                  # 右半有序
                if nums[mid] < target <= nums[r]:
                    l = mid + 1
                else:
                    r = mid - 1
        return -1


if __name__ == "__main__":
    s = Solution()
    assert s.search([4, 5, 6, 7, 0, 1, 2], 0) == 4
    assert s.search([4, 5, 6, 7, 0, 1, 2], 3) == -1
    assert s.search([1], 0) == -1
    assert s.search([3, 1], 1) == 1
    print("66 OK")
