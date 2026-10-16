"""
NC125 和为K的连续子数组
https://www.nowcoder.com/practice/704c8388a82e42e58b7f5751ec943a11

描述
给定一个无序数组 arr，其中元素可正、可负、可 0。给定一个整数 k，求 arr 所有连续子数组中
累加和为 k 的最长连续子数组长度。保证至少存在一个合法的连续子数组。
[1,2,3] 的连续子数组有 [1,2]，[2,3]，[1,2,3]，但是 [1,3] 不是

示例1
输入：[1,-2,1,1,1],0
返回值：3

示例2
输入：[0,1,2,3],3
返回值：3
"""

# 1. 连续
# 2. 和为 k
# 3. 最长
#
# 前缀和 + hash
#
# 和前缀和中的「和大于等于 k 的最短子数组」比较，一个最短，一个最长
# 和最长不含重复字符的子字符串比较，需要用到前缀和，hash


def maxlen_equal_k(arr, k):
    # [1,-2,1,1,1], 0 => 3
    # 前缀和 + hash
    if len(arr) == 0:
        return 0

    max_len = 0
    total = 0
    # value 存索引，不是在或者不在
    temp = {}
    # 此处很重要，两个索引相减，如果是首位，长度需要加 1
    temp[0] = -1
    for i, v in enumerate(arr):
        total += v
        # 符合条件的数据在 i, j 之间
        # j 是之前元素的索引
        # i 是 sum 加上 v 的索引，当前索引
        if (total - k) in temp:
            j = temp[total - k]
            max_len = max(max_len, i - j)
        # 不存在，构建 sum 为 key，索引为 value 的 temp hash 表
        if total not in temp:
            temp[total] = i
    return max_len


if __name__ == "__main__":
    # k = maxlen_equal_k([1, -2, 1, 1, 1], 0)
    # k = maxlen_equal_k([1, -2, 1, 1, 1, -1, 1], 2)
    k = maxlen_equal_k([1, -2, 1, 1], 2)
    print(k)
