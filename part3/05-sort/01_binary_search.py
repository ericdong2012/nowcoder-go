# -*- coding: utf-8 -*-
from typing import List


# 很重要的前提是： 有序


def binarySearch(arr, x):
    # 很重要的前提是： 有序
    # 时间复杂度: O(n)
    low, high = 0, len(arr) - 1
    sorted(arr)
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] > x:
            high = mid - 1
        elif arr[mid] < x:
            low = mid + 1
        else:
            return True
    return False


arr = [2, 3, 40, 10, 4, 11, 12, 15]
x = 13

result = binarySearch(arr, x)
print(result)
