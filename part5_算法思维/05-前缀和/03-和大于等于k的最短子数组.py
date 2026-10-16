"""
NC343 和大于等于K的最短子数组
https://www.nowcoder.com/practice/3e1fd3d19fb0479d94652d49c7e1ead1

知识点：二分、前缀和、队列、堆、双指针

描述
给定一个长度为 n 的整数数组，和一个目标值 k，请你找出这个整数数组中和大于等于 k 的最短子数组的长度。
如果不存在和大于等于 k 的子数组则输出 -1。

示例1
输入：[2,1,2,3],5
返回值：2

示例2
输入：[2,3,4,5],16
返回值：-1
"""

# 和「最短无序连续子数组」、「和为 k 的连续子数组」对比


def _min2(x, y):
    if x < y:
        return x
    return y


def shortest_subarray(nums, target):
    # 记录长度 (math.MaxInt32)
    res = 2 ** 31 - 1
    # 左右指针，可以理解为 r 快指针，l 慢指针
    l, r = 0, 0
    # 中间结果，与 target 做比较
    total = 0
    # [2,1,2,3], 5
    """
    2，1，2
    1，2，3
    2，3
    """
    while r < len(nums):
        total += nums[r]
        while total >= target:
            res = _min2(res, r - l + 1)
            total -= nums[l]
            l += 1
        r += 1
    if l != 0:
        return res
    else:
        return -1
