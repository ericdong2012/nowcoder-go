"""198. 打家劫舍（DP）
题目：不能偷相邻两家，求最大金额。
思路：rob[i] = max(rob[i-1], rob[i-2] + nums[i])——偷这家(上一家必不偷)
      或不偷这家。滚动变量。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        prev, cur = 0, 0               # rob[i-2], rob[i-1]
        for x in nums:
            prev, cur = cur, max(cur, prev + x)
        return cur


if __name__ == "__main__":
    s = Solution()
    assert s.rob([1, 2, 3, 1]) == 4
    assert s.rob([2, 7, 9, 3, 1]) == 12
    assert s.rob([2, 1, 1, 2]) == 4    # 偷第1和第4家
    print("83 OK")
