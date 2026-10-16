"""
NC213 除自身以外数组的乘积
https://www.nowcoder.com/practice/0786aa81c1c64c2a990e393fac811b45

知识点：数组、前缀和

描述
给定一个长度为 n 的数组 nums，返回一个数组 res，res[i] 是 nums 数组中除了 nums[i] 本身以外其余所有元素的乘积，即：
res[i] = nums[1] × nums[2] × ...... × nums[i-1] × nums[i+1] ...... × nums[n]
1.请不要使用除法，并且在 O(n) 时间复杂度内完成此题。
2.题目数据保证 res 数组的元素都在 32 位整数范围内。
3.有 O(1) 空间复杂度的做法，返回的 res 数组不计入空间复杂度计算。

输入：[1,2,3,4]
返回值：[24,12,8,6]

说明：
res[0] = 1 * 2*3*4 = 24
res[1] = 1 * 1*3*4 = 12
res[2] = 1*2 * 1*4 = 8
res[3] = 1*2*3 * 1 = 6
"""


def times_except_self(nums):
    """write code here"""
    n = len(nums)
    l, r = [0] * n, [0] * n
    l[0] = 1
    r[n - 1] = 1
    """
    l: 1, 1, 1*2, 1*2*3
    r: (从右往左) 1*4*3*2, 1*4*3, 1*4, 1

    res[0] = 1 * 2*3*4 = 24
    res[1] = 1 * 1*3*4 = 12
    res[2] = 1*2 * 1*4 = 8
    res[3] = 1*2*3 * 1 = 6
    """
    for i in range(1, n):
        l[i] = l[i - 1] * nums[i - 1]

    for j in range(n - 2, -1, -1):
        r[j] = r[j + 1] * nums[j + 1]

    res = [0] * n
    for k in range(n):
        res[k] = l[k] * r[k]

    return res


# 对于大量数据，可能时间复杂度会高一些
# def times_except_self(nums):
#     n = len(nums)
#     res = [0] * n
#     for i in range(n):
#         temp = 1
#         for j in range(i):
#             temp *= nums[j]
#         for k in range(i + 1, n):
#             temp *= nums[k]
#         res[i] = temp
#     return res
