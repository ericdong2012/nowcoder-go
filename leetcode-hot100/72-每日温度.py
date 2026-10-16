"""739. 每日温度（单调栈）
题目：等几天后温度更高；没有更高的为 0。
思路：单调递减栈存下标；新温度比栈顶高 → 弹出并结算天数。
复杂度：时间 O(n)（每元素进出栈各一次），空间 O(n)。
"""

from typing import List


class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        answer = [0] * n
        stack = []                     # 下标，温度单调递减
        for i, t in enumerate(temperatures):
            while stack and temperatures[stack[-1]] < t:
                j = stack.pop()
                answer[j] = i - j
            stack.append(i)
        return answer


if __name__ == "__main__":
    s = Solution()
    assert s.dailyTemperatures([73, 74, 75, 71, 69, 72, 76, 73]) == [1, 1, 4, 2, 1, 1, 0, 0]
    assert s.dailyTemperatures([30, 40, 50, 60]) == [1, 1, 1, 0]
    assert s.dailyTemperatures([30, 60, 90]) == [1, 1, 0]
    print("72 OK")
