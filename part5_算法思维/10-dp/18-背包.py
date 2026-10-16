"""
NC145 01背包
https://www.nowcoder.com/practice/2820ea076d144b30806e72de5e5d4bbf?tpId=117&tqId=37856

已知一个背包最多能容纳体积之和为v的物品, 现有n个物品， 第i个物品的体积为 vi , 重量为 wi
求当前背包最多能装多大重量的物品?

示例1
输入：10,2,[[1,3],[10,4]]
返回值：4

示例2
输入：10,2,[[1,3],[9,8]]
返回值：11
"""

# V是体积
# n是个数
# 体积，重量的数组


def knapsack(V, n, vw):
    # 索引是体积, 存的是重量(好好体会)    长度是最大体积+1, 因为要求最大能到V
    dp = [0] * (V + 1)
    for i in range(n):
        v, w = vw[i][0], vw[i][1]
        for j in range(V, v - 1, -1):
            dp[j] = Max(dp[j], dp[j - v] + w)
        # for j in range(V, vw[i][0] - 1, -1):
        #     dp[j] = Max(dp[j], dp[j - vw[i][0]] + vw[i][1])
    return dp[V]


def Max(a, b):
    if a > b:
        return a
    return b


if __name__ == "__main__":
    print(knapsack(10, 2, [[1, 3], [10, 4]]))
