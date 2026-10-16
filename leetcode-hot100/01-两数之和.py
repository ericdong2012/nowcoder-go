"""1. 两数之和（哈希）
题目：nums 中找两数之和 == target，返回下标。
思路：一遍遍历，哈希表存 {值: 下标}；对每个 x 查 target-x 是否已出现。
复杂度：时间 O(n)，空间 O(n)。
"""

from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}                                   # 值 -> 下标
        for i, x in enumerate(nums):
            if target - x in seen:
                return [seen[target - x], i]
            seen[x] = i
        return []


if __name__ == "__main__":
    s = Solution()
    assert s.twoSum([2, 7, 11, 15], 9) == [0, 1]
    assert s.twoSum([3, 2, 4], 6) == [1, 2]
    assert s.twoSum([3, 3], 6) == [0, 1]
    print("01 OK")
