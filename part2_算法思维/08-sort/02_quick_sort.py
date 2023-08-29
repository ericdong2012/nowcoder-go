# -*- coding: utf-8 -*-

"""
思想：
1. 找基准值
2. 分治：将大于基准的放到右边，小于基准值的放到左边
3. 递归对两组数据分别排序

"""


# method1:
# def quick_sort(b):
#     """快速排序"""
#     if len(b) < 2:
#         return b
#
#     # 选取基准，随便选哪个都可以，选中间的便于理解
#     pivot = b[len(b) // 2]
#
#     # 定义基准值左右两个数列
#     left, right = [], []
#
#     # 从原始数组中移除基准值
#     b.remove(pivot)
#
#     for item in b:
#         # 大于基准值放右边
#         if item >= pivot:
#             right.append(item)
#         else:
#             # 小于基准值放左边
#             left.append(item)
#
#     # 使用迭代进行比较
#     return quick_sort(left) + [pivot] + quick_sort(right)
#
#
# b = [11, 99, 33, 69, 77, 88, 55, 11, 33, 36, 39, 66, 44, 22]
# print(quick_sort(b))


# method2: recommand
# 实际运行可能会报错，但是体现了思想
def quick_sort(array):
    if len(array) < 2:
        return array

    pivot = array[len(array) // 2]
    small_than_pivot = [x for x in array if x <= pivot]
    big_than_pivot = [x for x in array if x > pivot]
    return quick_sort(small_than_pivot) + [pivot] + quick_sort(big_than_pivot)


# func quickSort(arr []int) []int {
# 	if len(arr) < 2  {
# 		return arr
# 	}
#
# 	pivot := arr[len(arr) / 2]
# 	small_than_pivot := func () []int{
# 		result := []int{}
# 		for _, v := range arr {
# 			if v < pivot {
# 				result = append(result, v)
# 			}
# 		}
# 		return result
# 	}
#
# 	big_than_pivot := func () []int{
# 		result := []int{}
# 		for _, v := range arr {
# 			if v > pivot {
# 				result = append(result, v)
# 			}
# 		}
# 		return result
# 	}
# 	res := []int{}
# 	res = append(res, small_than_pivot()...)
# 	res = append(res, pivot)
# 	res = append(res, big_than_pivot()...)
# 	return res
# }


b = [11, 99, 33, 69, 77, 88, 55, 11, 33, 36, 39, 66, 44, 22]
print(quick_sort(b))
"""
时间复杂度：快排的时间复杂度为O(nlogn), 最坏 n2
空间复杂度：排序时需要另外申请空间，并且随着数列规模增大而增大，其复杂度为：O(nlogn)
"""
