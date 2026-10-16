"""560. 和为 K 的子数组（前缀和+哈希）
题目：统计和恰好为 k 的连续子数组个数（元素可负，不能用滑窗）。
思路：前缀和 pre[i]；区间 [j+1, i] 和为 k ⇔ pre[i] - pre[j] == k，
      即查找之前出现过多少个前缀和等于 pre[i]-k。一遍扫描。
复杂度：时间 O(n)，空间 O(n)。
"""

from typing import List
from collections import defaultdict


class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        count = defaultdict(int)
        count[0] = 1                    # 空前缀：从头开始的子数组
        pre = 0
        ans = 0
        for x in nums:
            pre += x
            ans += count[pre - k]       # 有多少个前缀和等于 pre-k
            count[pre] += 1
        return ans


if __name__ == "__main__":
    s = Solution()
    assert s.subarraySum([1, 1, 1], 2) == 2
    assert s.subarraySum([1, 2, 3], 3) == 2
    assert s.subarraySum([1, -1, 0], 0) == 3     # 含负数用例
    print("10 OK")
