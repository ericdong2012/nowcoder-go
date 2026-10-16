"""
NC41 最长无重复子数组
https://www.nowcoder.com/practice/b56799ebfd684fb394bd315e89324fb4

描述
给定一个长度为 n 的数组 arr，返回 arr 的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如 [1,3,5,7,9] 的子数组有 [1,3]，[3,5,7] 等等，但是 [1,3,7] 不是子数组
"""

# 11-最长不含重复字符的子字符串.go 几乎一模一样，只是 map 中的 key 的类型不一样


def _max2(a, b):
    if a > b:
        return a
    else:
        return b


def max_length(arr):
    # # 双指针
    # n = len(arr)
    # m = {}
    # res = 0
    # l = 0
    # for r in range(n):
    #     # l 左指针记录是否有重复元素出现，有的话，往右移动
    #     l = max1(l, m.get(arr[r], 0))
    #     # 记录个数，索引之差
    #     res = max1(res, r - l + 1)
    #     # ??? 上面的 m[arr[r]] 中有值时，l 要往下移动一位，所以上次保存的 r + 1
    #     m[arr[r]] = r + 1
    # return res

    left, right = 0, 0
    maxlen = 0
    temp = {}
    while left < len(arr):
        if right == len(arr):
            return maxlen
        # arr[right] 在 temp 中不存在
        if temp.get(arr[right], 0) == 0:
            temp[arr[right]] = temp.get(arr[right], 0) + 1
            right += 1
        else:
            # 有重复的将上一个清零
            temp[arr[left]] = 0
            left += 1
        maxlen = _max2(maxlen, right - left)

    return maxlen


if __name__ == "__main__":
    length = max_length([2, 3, 4, 5])
    print(length)
