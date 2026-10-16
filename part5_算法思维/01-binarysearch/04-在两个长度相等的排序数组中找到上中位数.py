"""
NC36 在两个长度相等的排序数组中找到上中位数
https://www.nowcoder.com/practice/6fbe70f3a51d44fa9395cfc49694404f

描述
给定两个递增数组 arr1 和 arr2，已知两个数组的长度都为 N，求两个数组中所有数的上中位数。
上中位数：假设递增序列长度为 n，为第 n/2 个数

要求：时间复杂度 O(n)，空间复杂度 O(1)
进阶：时间复杂度为 O(logN)，空间复杂度为 O(1)

示例1
输入：[1,2,3,4], [3,4,5,6]
返回：3
说明：总共有 8 个数，上中位数是第 4 小的数，所以返回 3。
"""


def _min(a, b):
    if a < b:
        return a
    return b


def find_median_in_two_sorted_array(arr1, arr2):
    """write code here"""
    # 方式一：两个合并，排序，算长度，算中位数是第几个，直接取值
    """
    temp = arr1 + arr2
    temp.sort()
    return temp[len(temp) // 2 - 1]
    """

    # 方式二：
    if len(arr1) == 1:
        return _min(arr1[0], arr2[0])

    l1 = 0
    r1 = len(arr1) - 1
    l2 = 0
    r2 = len(arr2) - 1
    # 长度一样，以任意一个为主
    while l1 < r1:
        pos = (r1 - l1) % 2
        mid1 = (l1 + r1) // 2
        mid2 = (l2 + r2) // 2
        if arr1[mid1] > arr2[mid2]:
            l2 = mid2 + pos
            r1 = mid1
        elif arr1[mid1] < arr2[mid2]:
            l1 = mid1 + pos
            r2 = mid2
        else:
            return arr1[mid1]
    return _min(arr1[l1], arr2[l2])

    # 方式三：
    # 从后往前排，大的先放进去，然后依次比较，最终完成排序，最后取出第 n/2-1 个
