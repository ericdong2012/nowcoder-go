"""152. 乘积最大子数组（DP 双状态）
题目：连续子数组的最大乘积（含负数：负×负=正，所以还要维护最小值）。
思路：同时维护以 i 结尾的最大乘积 mx 和最小乘积 mn；
      新数 x：mx = max(x, mx*x, mn*x)，mn = min(x, mx_old*x, mn_old*x)。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        mx = mn = best = nums[0]
        for x in nums[1:]:
            candidates = (x, mx * x, mn * x)
            mx, mn = max(candidates), min(candidates)
            best = max(best, mx)
        return best


if __name__ == "__main__":
    s = Solution()
    assert s.maxProduct([2, 3, -2, 4]) == 6
    assert s.maxProduct([-2, 0, -1]) == 0
    assert s.maxProduct([-2, 3, -4]) == 24     # 负负得正
    print("88 OK")
