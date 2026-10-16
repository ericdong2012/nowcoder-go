"""
BM62 斐波那契数列
https://www.nowcoder.com/practice/c6c7742f5ba7442aada113136ddea0c3?tpId=295&tqId=23255&ru=/exam/oj&qru=/ta/format-top101/question-ranking

描述
大家都知道斐波那契数列，现在要求输入一个正整数 n ，请你输出斐波那契数列的第 n 项。
斐波那契数列是一个满足 fib(1)=1, fib(2)=1, fib(x)=fib(x-1)+fib(x-2) (x>2) 的数列
数据范围：1<=n<=40
要求：空间复杂度 O(1)，时间复杂度 O(n)，本题也有时间复杂度 O(logn) 的解法

示例1
输入：4
返回值：3

示例2
输入：1
返回值：1
"""

# 递归
# 动态规划
# 双指针， 记忆化搜索


def Fibonacci(n):
    dp = [0] * 50
    dp[1] = 1
    dp[2] = 1
    for i in range(3, n + 1):
        dp[i] = dp[i - 1] + dp[i - 2]
    return dp[n]


if __name__ == "__main__":
    print(Fibonacci(4))
