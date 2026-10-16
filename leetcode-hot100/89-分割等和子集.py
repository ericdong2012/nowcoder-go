"""416. 分割等和子集（0-1 背包可行性）
题目：能否把数组分成两个元素和相等的子集。
思路：总和为奇数必假；否则问"能否凑出 sum/2"——
      dp[j] 表示能否凑出和 j，倒序更新（每数只能用一次）。
复杂度：时间 O(n·sum)，空间 O(sum)。
"""

from typing import List


class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        total = sum(nums)
        if total % 2:
            return False
        target = total // 2
        dp = [True] + [False] * target
        for x in nums:
            for j in range(target, x - 1, -1):     # 倒序：每个数只用一次
                dp[j] = dp[j] or dp[j - x]
            if dp[target]:
                return True
        return dp[target]


if __name__ == "__main__":
    s = Solution()
    assert s.canPartition([1, 5, 11, 5]) is True
    assert s.canPartition([1, 2, 3, 5]) is False
    assert s.canPartition([1, 1]) is True
    print("89 OK")
