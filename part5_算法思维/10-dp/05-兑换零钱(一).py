"""
BM70 兑换零钱(一)
https://www.nowcoder.com/practice/3911a20b3f8743058214ceaa099eeb45?tpId=295&tqId=988994

描述
给定数组arr，arr中所有的值都为正整数且不重复。 每个值代表一种面值的货币，每种面值的货币可以使用任意张，
再给定一个aim，代表要找的钱数，求组成aim的最少货币数。
如果无解，请返回-1.

示例1
输入：[5,2,3],20
返回值：4

示例2
输入：[5,2,3],0
返回值：0

示例3
输入：[3,5],2
返回值：-1
"""

# 三数之和的升级版
# 背包问题简化版
# 最长公共子序列(二) 也是背包问题
#
# 背包问题：
# https://blog.csdn.net/qq_22328011/article/details/124138071
#
# coin:
# https://blog.csdn.net/nettee/article/details/106935507


def minMoney(arr, aim):
    # write code here
    if aim < 1:
        return 0
    # 建立dp, 存的是最终的张数，索引是从小到aim， 给首位赋值0， 其他位aim + 1
    dp = [0] * (aim + 1)
    dp[0] = 0
    for i in range(1, aim + 1):
        dp[i] = aim + 1
    # dp 要凑的钱, 最小是1， 最大是aim
    for i in range(1, aim + 1):
        # arr 面值(每种面值的货币都要枚举)
        for j in range(len(arr)):
            # 在相等的时候, dp[i-arr[j]] = 0, 而张数最小为1, 所以在最后加上1
            if arr[j] <= i:
                dp[i] = min(dp[i], dp[i - arr[j]] + 1)
    # 最终，如果dp倒数第二位大于aim, 可能是每个元素都大于aim.
    if dp[aim] > aim:
        return -1
    return dp[aim]


def main05():
    money = minMoney([5, 2, 3], 20)
    # money = minMoney([5, 2, 3], 11)  # 3
    print(money)


if __name__ == "__main__":
    main05()
