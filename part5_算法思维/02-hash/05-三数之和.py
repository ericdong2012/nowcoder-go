"""
NC54 三数之和
https://www.nowcoder.com/practice/345e2ed5f81d4017bbb8cc6055b0b711

描述
给出一个有 n 个元素的数组 S，S 中是否有元素 a,b,c 满足 a+b+c=0 找出数组 S 中所有满足条件的三元组。

注意：
三元组（a、b、c）中的元素可以按任意顺序排列。
解集中不能包含重复的三元组。

示例1
输入：[-10,0,10,20,-10,-40]
返回值：[[-10,-10,20],[-10,0,10]]

示例2
输入：[-2,0,1,1,2]
返回值：[[-2,0,2],[-2,1,1]]
"""

# 1、特判，如果数组为 null 或者数组长度小于 3，返回 []。
# 2、对数组进行排序。
# 3、遍历排序后数组:
#     1、若 num[i]>0：因为已经排序好，所以后面不可能有三个数加和等于 0，直接返回结果。
#     2、对于重复元素(num[i] == num[i-1])：跳过
#     3、令左指针 L=i+1，右指针 R=n-1，
#         1、当 nums[i]+nums[L]+nums[R]==0，执行循环，写入结果，同时判断左界和右界是否和下一位置重复，
#            去除重复解，并将 L,R 移到下一位置
#         2、若和大于 0，说明 nums[R] 太大，R 左移
#         3、若和小于 0，说明 nums[L] 太小，L 右移

import bisect


# def three_sum(num):
#     # 1.三元组非降序排列 -> 将 num 升序排列
#     # 2.数求和问题 -> 双指针
#     if len(num) < 3 or num is None:
#         return []
#     # 排序
#     num.sort()
#     n = len(num)
#     res = []
#     for i in range(n - 2):
#         if num[i] > 0:
#             break  # 排好序后的数组，如果当前数字大于 0，则三数之和一定大于 0，后面两个数字一定比当前大
#         if i > 0 and num[i] == num[i - 1]:
#             continue  # 去重
#         # 双指针求和
#         low, high = i + 1, n - 1  # 一开始就是最大值和最小值
#         while low < high:
#             total = num[i] + num[low] + num[high]
#             if total == 0:
#                 # 写入结果
#                 res.append([num[i], num[low], num[high]])
#                 # 头部去重（如果后面一个数跟当前的数字相等，则代表有重复的结果生成，跳过
#                 while low < high and num[low] == num[low + 1]:
#                     low += 1
#                 low += 1
#                 high -= 1
#             elif total > 0:
#                 high -= 1  # 太大前移
#             else:
#                 low += 1  # 太小后移
#     return res


def three_sum(nums):
    nums.sort()
    result_map = {}
    for i in range(len(nums) - 2):
        for j in range(i + 1, len(nums) - 1):
            sub = nums[j + 1:]
            k = bisect.bisect_left(sub, -nums[i] - nums[j])
            if k < len(sub) and nums[i] + nums[j] + sub[k] == 0:
                result_map[(nums[i], nums[j], sub[k])] = None

    result = []
    for k in result_map:
        result.append([k[0], k[1], k[2]])
    return result


if __name__ == "__main__":
    total = three_sum([-2, 0, 1, 2, 1, 2])
    print(total)
