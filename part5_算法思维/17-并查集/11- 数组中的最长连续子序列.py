"""
NC95 数组中的最长连续子序列
https://www.nowcoder.com/practice/eac1c953170243338f941959146ac4bf

描述
给定无序数组arr，返回其中最长的连续序列的长度(要求值连续，位置可以不连续,例如 3,4,5,6为连续的自然数）

示例1
输入：
[100,4,200,1,3,2]
返回值：
4

示例2
输入：
[1,1,1]
返回值：
1
"""


"""
题目理解： 连续子序列，必须是连续，没有要求是升序，可以是乱序， 比如： 1，2，3，4 或者 4，3，2，1 都是满足要求，
思路：    先排序，后遍历，找到后一个比前一个大1， 计数加1， 否则置1; 边界情况： 前一个和当前相等， 跳过
"""


def mls(arr):
    # write code here
    arr.sort()
    count = 1
    res = 1
    for i in range(1, len(arr)):
        if arr[i] == arr[i - 1]:
            continue
        if arr[i] - arr[i - 1] == 1:
            count += 1
        else:
            count = 1
        res = max(res, count)

    return res


def max(a, b):
    if a < b:
        return b
    return a


# def mls(arr):
#     # write code here
#     num_set = set(arr)
#     longest_seq = 0
#     for num in num_set:
#         if num - 1 not in num_set:
#             curr_seq = 1
#             curr = num
#             while curr + 1 in num_set:
#                 curr_seq += 1
#                 curr += 1
#             if curr_seq > longest_seq:
#                 longest_seq = curr_seq
#     return longest_seq


if __name__ == "__main__":
    print(mls([100, 4, 200, 1, 3, 2]))
    print(mls([1, 1, 1]))
