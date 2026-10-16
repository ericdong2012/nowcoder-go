"""300. 最长递增子序列（DP / 贪心+二分）
题目：最长严格递增子序列长度。
思路 A：dp[i] = 以 i 结尾的 LIS 长度，O(n²)。
思路 B（推荐）：patience 数组 tails[k] = 长度 k+1 的 LIS 的最小结尾；
      每个数二分找插入位（lower_bound），替换或追加。O(n log n)。
复杂度：B 时间 O(n log n)，空间 O(n)。
"""

from typing import List
import bisect


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        tails = []                       # 有序数组
        for x in nums:
            i = bisect.bisect_left(tails, x)
            if i == len(tails):
                tails.append(x)          # 延长最长子序列
            else:
                tails[i] = x             # 让结尾更小，利于后续延伸
        return len(tails)


if __name__ == "__main__":
    s = Solution()
    assert s.lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18]) == 4
    assert s.lengthOfLIS([0, 1, 0, 3, 2, 3]) == 4
    assert s.lengthOfLIS([7, 7, 7, 7]) == 1
    print("87 OK")
