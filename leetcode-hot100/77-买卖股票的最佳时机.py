"""121. 买卖股票的最佳时机（一次遍历）
题目：一次买卖的最大利润。
思路：记录历史最低价，每天算"今天卖"的利润，取最大。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        min_price = float("inf")
        best = 0
        for p in prices:
            min_price = min(min_price, p)
            best = max(best, p - min_price)
        return best


if __name__ == "__main__":
    s = Solution()
    assert s.maxProfit([7, 1, 5, 3, 6, 4]) == 5
    assert s.maxProfit([7, 6, 4, 3, 1]) == 0
    print("77 OK")
