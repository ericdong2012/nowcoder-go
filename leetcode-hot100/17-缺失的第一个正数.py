"""41. 缺失的第一个正数（原地哈希）
题目：无序数组中没出现的最小正整数，要求 O(n) 时间 O(1) 空间。
思路：答案必在 [1, n+1]（n 为长度）。把值 v 放到下标 v-1 上
     （交换归位），再扫一遍第一个 nums[i] != i+1 的位置即答案。
复杂度：时间 O(n)（每个数最多被交换一次到位），空间 O(1)。
"""

from typing import List


class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        n = len(nums)
        for i in range(n):
            # 把 nums[i] 换到它应在的位置，直到当前位置放对或值越界/重复
            while 1 <= nums[i] <= n and nums[nums[i] - 1] != nums[i]:
                nums[nums[i] - 1], nums[i] = nums[i], nums[nums[i] - 1]
        for i in range(n):
            if nums[i] != i + 1:
                return i + 1
        return n + 1


if __name__ == "__main__":
    s = Solution()
    assert s.firstMissingPositive([1, 2, 0]) == 3
    assert s.firstMissingPositive([3, 4, -1, 1]) == 2
    assert s.firstMissingPositive([7, 8, 9, 11, 12]) == 1
    print("17 OK")
