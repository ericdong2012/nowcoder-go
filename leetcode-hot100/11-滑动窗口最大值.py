"""239. 滑动窗口最大值（单调队列）
题目：返回每个大小为 k 的滑动窗口的最大值。
思路：双端队列存下标，保持对应值单调递减；队头即当前窗口最大。
      新元素入队前弹出所有 <= 它的队尾（它们永远不可能再成为最大）。
复杂度：时间 O(n)（每元素进出队各一次），空间 O(k)。
"""

from typing import List
from collections import deque


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        q = deque()          # 存下标，值单调递减
        res = []
        for i, x in enumerate(nums):
            while q and nums[q[-1]] <= x:
                q.pop()
            q.append(i)
            if q[0] <= i - k:             # 队头滑出窗口
                q.popleft()
            if i >= k - 1:
                res.append(nums[q[0]])
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.maxSlidingWindow([1, 3, -1, -3, 5, 3, 6, 7], 3) == [3, 3, 5, 5, 6, 7]
    assert s.maxSlidingWindow([1], 1) == [1]
    print("11 OK")
