"""70. 爬楼梯（DP 入门）
题目：每次爬 1 或 2 阶，n 阶有多少种爬法。
思路：f(n) = f(n-1) + f(n-2)（最后一步是 1 或 2），滚动变量。
复杂度：时间 O(n)，空间 O(1)。
"""


class Solution:
    def climbStairs(self, n: int) -> int:
        a, b = 1, 1                    # f(0), f(1)
        for _ in range(2, n + 1):
            a, b = b, a + b
        return b


if __name__ == "__main__":
    s = Solution()
    assert s.climbStairs(2) == 2
    assert s.climbStairs(3) == 3
    assert s.climbStairs(10) == 89
    print("81 OK")
