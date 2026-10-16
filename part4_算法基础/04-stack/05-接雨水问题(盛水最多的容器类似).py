"""
BM94 接雨水问题
https://www.nowcoder.com/practice/31c1aed01b394f0b8b7734de0324e00f

描述
给定一个整形数组 arr，已知其中所有的值都是非负的，将这个数组看作一个柱子高度图，
计算按此排列的柱子，下雨之后能接多少雨水。(数组以外的区域高度视为 0)

要求：时间复杂度 O(n)

示例1
输入：[3,1,2,5,2,4]
返回值：5

示例2
输入：[4,5,1,3,2]
返回值：2
"""

# 总结：第二高减去剩下的，并累加
#
# 方法一：双指针
#   那这种情况下它是不是将一个水桶分割成了两个水桶，而中间的那条边就是两个水桶的边。
#   step 1：检查数组是否为空的特殊情况
#   step 2：准备双指针，分别指向数组首尾元素，代表最初的两个边界
#   step 3：指针往中间遍历，遇到更低柱子就是底，用较短的边界减去底(计算公式)就是这一列的接水量，
#           遇到更高的柱子就是新的边界，更新(累加)边界大小。
#
# 方法二：单调栈


def max_water(arr):
    """write code here"""
    if len(arr) == 0:
        return 0
    res = 0
    left, right = 0, len(arr) - 1
    max_l, max_r = 0, 0
    while left <= right:
        # [3,1,2,5,2,4] 看运行过程，总结规律，再验证特例/边界情况
        # 找左边右边较大值，一边结果一直在增加，当找到最大值，条件反转(走到另一个 if 条件中)，
        # 另一边结果一直增加；直到 left、right 走到一起，即最大值的地方，跳出循环
        max_l = max(max_l, arr[left])
        max_r = max(max_r, arr[right])
        if arr[left] <= arr[right]:
            res += max_l - arr[left]
            left += 1
        else:
            res += max_r - arr[right]
            right -= 1

    return res


if __name__ == "__main__":
    pass
