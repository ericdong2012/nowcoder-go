"""
BM64 最小花费爬楼梯
https://www.nowcoder.com/practice/6fe0302a058a4e4a834ee44af88435c7?tpId=295&tqId=2366451

描述
给定一个整数数组 cost， 其中 cost[i] 是从楼梯第i个台阶向上爬需要支付的费用，下标从0开始。
你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。
一旦你支付此费用，即可选择向上爬一个或者两个台阶。
请你计算并返回达到楼梯顶部的最低花费。

示例1
输入：[2,5,20]
返回值：5

示例2
输入：[1,100,1,1,1,90,1,1,80,1]
返回值：6
"""


def minCostClimbingStairs(cost):
    # write code here
    # 建立二维数组(上方是数字，下方是花费)
    # 1 100 1 1 1 90 1 1 80 1
    # 0 0   1 2 2 3  3 4 4  5
    # 再建立方程
    n = len(cost)
    dp = [0] * n
    dp[0] = cost[0]
    dp[1] = cost[1]
    for i in range(2, n):
        # 当前的消费 = 上一步或者上两步花费的较小值 + 当前的花费
        dp[i] = min(dp[i - 1], dp[i - 2]) + cost[i]
    # for i in range(n - 2):
    #     dp[i] = min(cost[i+1], cost[i+2]) + cost[i]
    return min(dp[n - 1], dp[n - 2])
    # n = len(cost)
    # dp = [0] * (n + 1)
    # for i in range(2, n + 1):
    #     dp[i] = min(dp[i-1] + cost[i-1], dp[i-2] + cost[i-2])
    # return dp[n]


def main03():
    # stairs = minCostClimbingStairs([2, 5, 20])
    stairs = minCostClimbingStairs([1, 100, 1, 1, 1, 90, 1, 1, 80, 1])
    print(stairs)


if __name__ == "__main__":
    main03()
