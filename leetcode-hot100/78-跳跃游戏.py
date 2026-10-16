"""55. 跳跃游戏（贪心）
题目：能否从下标 0 跳到最后（nums[i] 是最大步长）。
思路：维护"当前能到达的最远下标"；扫到某位置 > 最远 → 卡死；
     最远 ≥ n-1 即成功。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        farthest = 0
        for i, step in enumerate(nums):
            if i > farthest:
                return False
            farthest = max(farthest, i + step)
            if farthest >= len(nums) - 1:
                return True
        return True


if __name__ == "__main__":
    s = Solution()
    assert s.canJump([2, 3, 1, 1, 4]) is True
    assert s.canJump([3, 2, 1, 0, 4]) is False
    assert s.canJump([0]) is True
    print("78 OK")
