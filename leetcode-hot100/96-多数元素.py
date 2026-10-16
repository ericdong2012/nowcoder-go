"""169. 多数元素（摩尔投票）
题目：出现次数超过 n/2 的元素，O(1) 空间。
思路：摩尔投票——不同候选互相抵消；超过半数的元素必存活。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        candidate, count = 0, 0
        for x in nums:
            if count == 0:
                candidate = x
            count += 1 if x == candidate else -1
        return candidate


if __name__ == "__main__":
    s = Solution()
    assert s.majorityElement([3, 2, 3]) == 3
    assert s.majorityElement([2, 2, 1, 1, 1, 2, 2]) == 2
    print("96 OK")
