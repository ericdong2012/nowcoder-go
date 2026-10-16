"""
https://www.nowcoder.com/practice/b3b59248e61f499482eaba636305474b

给定两个长度为 n 和 m 的升序数组（后一个数一定大于等于前一个数），请你找到这两个数组中全部元素的中位数。

[1,2,3,4,5],[6,7,8,9] => 5

[1,2,3,8,9],[4,5,6,7] => 5
"""


# 合并两个有序数组(double pointer)类似， 先将两个有序数组合并，  最后加上中间数判断即可
def median(nums1, nums2):
    m = len(nums1)
    n = len(nums2)
    index = m + n
    result = [0] * (m + n)
    while m > 0 and n > 0:
        # 先比较最后一位
        if nums1[m - 1] >= nums2[n - 1]:
            result[index - 1] = nums1[m - 1]
            m -= 1
        else:
            result[index - 1] = nums2[n - 1]
            n -= 1
        index -= 1
    # B数组依然有剩余
    while n > 0 and index > 0:
        result[index - 1] = nums2[n - 1]
        n -= 1
        index -= 1
    # A数组依然有剩余
    while m > 0 and index > 0:
        result[index - 1] = nums1[m - 1]
        m -= 1
        index -= 1

    length = len(nums1) + len(nums2)
    if length % 2 == 0:
        return (result[length // 2 - 1] + result[length // 2]) / 2.0
    else:
        return float(result[length // 2])


# def median(nums1, nums2):
#     l1, l2 = len(nums1), len(nums2)
#     r1, r2 = 0, 0
#     cnt, p1, p2 = 0, 0, 0
#     # [1,2,3,4] [6,7,8,9]
#     # [1,2,3,4,5] [6,7,8,9]
#     while cnt <= (l1 + l2) // 2:
#         # 比右侧小前， 左侧一直往前走(p1) r1, r2一直在变化
#         if p1 < l1 and p2 < l2 and nums1[p1] < nums2[p2] or p2 == l2:
#             # r1: 0   r2: 1   p1: 1
#             # r1: 1   r2: 2   p1: 2
#             # r1: 2   r2: 3   p1: 3
#             # r1: 3   r2: 4   p1: 4    cnt: 4
#
#             # 左侧有多的 比如 [1,2,3,4,5]
#             # r1: 4   r2: 5   p1: 5    cnt: 5
#             r1, r2 = r2, nums1[p1]
#             p1 += 1
#         else:
#             # r1: 4  r2: 6   p2: 1  p1:4  cnt: 5
#             r1, r2 = r2, nums2[p2]
#             p2 += 1
#         cnt += 1
#
#     # r1: 4  r2: 6   p2: 1  p1:4  cnt: 5
#     if (l1 + l2) % 2 == 0:
#         return float(r1 + r2) / 2
#
#     return float(r2)


# def find_median_sorted_arrays(nums1, nums2):
#     n, m = len(nums1), len(nums2)
#     left, right = (n + m + 1) // 2, (n + m + 2) // 2
#     return (get_kth(nums1, nums2, left) + get_kth(nums1, nums2, right)) / 2
#
#
# def get_kth(nums1, nums2, k):
#     n, m = len(nums1), len(nums2)
#     if n > m:
#         return get_kth(nums2, nums1, k)
#     if n == 0:
#         return nums2[k - 1]
#     if k == 1:
#         return min(nums1[0], nums2[0])
#     i, j = min(n, k // 2), min(m, k // 2)
#     if nums1[i - 1] > nums2[j - 1]:
#         return get_kth(nums1, nums2[j:], k - j)
#     else:
#         return get_kth(nums1[i:], nums2, k - i)


def min(x, y):
    if x < y:
        return x
    return y


if __name__ == "__main__":
    print(median([1, 2, 3, 4, 5], [6, 7, 8, 9]))
    print(median([1, 2, 3, 8, 9], [4, 5, 6, 7]))
