"""
题目大意

给你两个数组，arr1 和 arr2
arr2 中的元素各不相同
arr2 中的每个元素都出现在 arr1 中
对 arr1 中的元素进行排序，使 arr1 中项的相对顺序和 arr2 中的相对顺序相同。
未在 arr2 中出现过的元素需要按照升序放在 arr1 的末尾。

提示：
arr1.length, arr2.length <= 1000
0 <= arr1[i], arr2[i] <= 1000
arr2 中的元素 arr2[i] 各不相同
arr2 中的每个元素 arr2[i] 都出现在 arr1 中

解题思路
给出 2 个数组 A 和 B，A 中包含 B 中的所有元素。要求 A 按照 B 的元素顺序排序，
B 中没有的元素再接着排在后面，从小到大排序。
这一题有多种解法。一种暴力的解法就是依照题意，先把 A 中的元素都统计频次放在 map 中，
然后按照 B 的顺序输出，接着再把剩下的元素排序接在后面。还有一种桶排序的思想，
由于题目限定了 A 的大小是 1000，这个数量级很小，所以可以用 1001 个桶装所有的数，
把数都放在桶里，这样默认就已经排好序了。接下来的做法和前面暴力解法差不多，按照频次输出。
B 中以外的元素就按照桶的顺序依次输出即可

https://github.com/halfrost/LeetCode-Go/tree/master/leetcode/1122.Relative-Sort-Array
"""


def relative_sort_array(arr1, arr2):
    # 使用哈希表记录 arr2 中出现过的数字及其对应的下标
    # index = {}
    # for i, v in enumerate(arr2):
    #     index[v] = i
    #
    # # 定义排序函数
    # arr1.sort(key=cmp_to_key(...))
    # （Python 中可用 functools.cmp_to_key，这里略）

    # 对 arr2 做 hash，arr2 中数据各不相同
    index = {v: i for i, v in enumerate(arr2)}
    # 拿到不在 arr2 中的 arr1 中数
    not_in = []
    for v in arr1:
        if v not in index:
            not_in.append(v)
    # 对 arr1 中数，计数，关键步骤
    count = {}
    for v in arr1:
        if v in index:
            count[v] = count.get(v, 0) + 1

    # 处理结果
    # 写法1
    # res = []
    # for v in index:
    #     temp = []
    #     for _ in range(count[v]):
    #         temp.append(v)
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


# 解法一 桶排序，时间复杂度 O(n^2)
def relative_sort_array1(a, b):
    count = [0] * 1001
    for x in a:
        count[x] += 1
    res = []
    for x in b:
        while count[x] > 0:
            res.append(x)
            count[x] -= 1
    for i in range(1001):
        while count[i] > 0:
            res.append(i)
            count[i] -= 1
    return res


# 解法二 模拟，时间复杂度 O(n^2)
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
