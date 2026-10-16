"""53. 最大子数组和（DP / 前缀思想）
题目：连续子数组的最大和。
思路：dp[i] = 以 i 结尾的最大子数组和 = max(dp[i-1]+x, x)；
      "前面的和若为负不如重新开始"。空间压到一维变量。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List
import math


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        cur = best = nums[0]
        for x in nums[1:]:
            cur = max(cur + x, x)         # 接上前缀 vs 从这里重新开始
            best = max(best, cur)
        return best


if __name__ == "__main__":
    s = Solution()
    assert s.maxSubArray([-2, 1, -3, 4, -1, 2, 1, -5, 4]) == 6
    assert s.maxSubArray([1]) == 1
    assert s.maxSubArray([-1]) == -1
    assert s.maxSubArray([5, 4, -1, 7, 8]) == 23
    print("13 OK")
