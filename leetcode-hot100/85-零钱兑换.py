"""322. 零钱兑换（完全背包）
题目：coins 凑出 amount 的最少硬币数（每种无限用）；凑不出返回 -1。
思路：dp[i] = min(dp[i - coin] + 1)；dp[0]=0。
复杂度：时间 O(amount·len(coins))，空间 O(amount)。
"""

from typing import List


class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = [0] + [float("inf")] * amount
        for coin in coins:
            for i in range(coin, amount + 1):
                dp[i] = min(dp[i], dp[i - coin] + 1)
        return dp[amount] if dp[amount] != float("inf") else -1


if __name__ == "__main__":
    s = Solution()
    assert s.coinChange([1, 2, 5], 11) == 3
    assert s.coinChange([2], 3) == -1
    assert s.coinChange([1], 0) == 0
    print("85 OK")
