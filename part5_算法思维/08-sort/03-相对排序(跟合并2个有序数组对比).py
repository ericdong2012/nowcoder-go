"""
题⽬⼤意

给你两个数组，arr1 和 arr2
arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进⾏排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。未在 arr2 中出现过的
元素需要按照升序放在 arr1 的末尾。

提示：
arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
arr2 中的元素 arr2[i] 各不相同
arr2 中的每个元素 arr2[i] 都出现在 arr1 中


解题思路
给出 2 个数组 A 和 B，A 中包含 B 中的所有元素。要求 A 按照 B 的元素顺序排序，B 中没有的元
素再接着排在后⾯，从⼩到⼤排序。
这⼀题有多种解法。⼀种暴⼒的解法就是依照题意，先把 A 中的元素都统计频次放在 map 中，然
后 按照 B 的顺序输出，接着再把剩下的元素排序接在后⾯。还有⼀种桶排序的思想，由于题⽬限
定了 A 的⼤⼩是 1000，这个数量级很⼩，所以可以⽤ 1001 个桶装所有的数，把数都放在桶⾥，
这样默认就已经排好序了。接下来的做法和前⾯暴⼒解法差不多，按照频次输出。B 中以外的元素
就按照桶的顺序依次输出即可

https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1122.Relative-Sort-Array
"""


def relative_sort_array(arr1, arr2):
    # 使用哈希表记录 arr2 中出现过的数字及其对应的下标
    # index = {}
    # for i, v in enumerate(arr2):
    #     index[v] = i
    #
    # # 定义排序函数
    # arr1.sort(key=lambda x: (
    #     index[x] if x in index else float('inf'),
    #     x if x not in index else 0
    # ))
    # return arr1

    # 对arr2 做hash, arr2 中数据各不相同
    index = {}
    for i, v in enumerate(arr2):
        index[v] = i
    # 对arr1中在arr2的数，计数, 关键步骤
    count = {}
    # 拿到不在arr2中的arr1中数
    not_in = []
    for v in arr1:
        if v not in index:
            not_in.append(v)
        else:
            count[v] = count.get(v, 0) + 1

    # 处理结果
    # 写法1
    # res = []
    # for v in index:
    #     temp = [v] * count.get(v, 0)
    #     res.extend(temp)

    # 写法2
    res = []
    for v in arr2:
        while count.get(v, 0) > 0:
            res.append(v)
            count[v] -= 1

    # 将剩余的排序，并添加进结果中
    not_in.sort()
    res.extend(not_in)

    return res


# 解法⼀ 桶排序，时间复杂度 O(n^2)
def relative_sort_array1(A, B):
    count = [0] * 1001
    for a in A:
        count[a] += 1
    res = []
    for b in B:
        while count[b] > 0:
            res.append(b)
            count[b] -= 1
    for i in range(1001):
        while count[i] > 0:
            res.append(i)
            count[i] -= 1
    return res


# 解法⼆ 模拟，时间复杂度 O(n^2)
def relative_sort_array2(arr1, arr2):
    leftover, m, res = [], {}, []
    for v in arr1:
        m[v] = m.get(v, 0) + 1

    for s in arr2:
        cnt = m[s]
        for _ in range(cnt):
            res.append(s)
        m[s] = 0
    for v, cnt in m.items():
        for _ in range(cnt):
            leftover.append(v)
    leftover.sort()
    res.extend(leftover)
    return res


if __name__ == "__main__":
    print(relative_sort_array([2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19], [2, 1, 4, 3, 9, 6]))
