"""45. 跳跃游戏 II（贪心，BFS 分层思想）
题目：最少跳数到达末尾（保证可达）。
思路：把每跳可达范围当一层；cur_end 是本层边界，
      扫描层内更新 farthest；到 cur_end 必须起跳（跳数+1）。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:
        steps = 0
        cur_end = 0          # 本层（当前一跳）能覆盖的最右
        farthest = 0
        for i in range(len(nums) - 1):     # 最后一个位置不用起跳
            farthest = max(farthest, i + nums[i])
            if i == cur_end:
                steps += 1
                cur_end = farthest
        return steps


if __name__ == "__main__":
    s = Solution()
    assert s.jump([2, 3, 1, 1, 4]) == 2
    assert s.jump([2, 3, 0, 1, 4]) == 2
    assert s.jump([1]) == 0
    assert s.jump([1, 1, 1, 1]) == 3
    print("79 OK")
