# -*- coding: utf-8 -*-
from typing import List


# 很重要的前提是: 有序
def binarySearch(arr, x):
    # 很重要的前提是： 有序
    # 时间复杂度: O(n)
    # sorted(arr)
    low, high = 0, len(arr) - 1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] > x:
            high = mid - 1
        elif arr[mid] < x:
            low = mid + 1
        else:
            return True
    return False

# func  binarySearch(arr []int , target int) int {
# 	left, right := 0 , len(arr) - 1
# 	for left <= right {
# 		mid := (left + right ) /2
# 		if arr[mid] < target {
# 			left = mid +1
# 		}else if arr[mid] > target {
# 			right = mid -1
# 		} else {
# 			// return true
# 			return mid
# 		}
# 	}
# 	// return false
# 	return -1
# }


arr = [2, 3, 40, 10, 4, 11, 12, 15]
x = 13

result = binarySearch(arr, x)
print(result)
