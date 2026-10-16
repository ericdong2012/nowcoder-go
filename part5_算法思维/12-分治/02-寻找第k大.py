"""
https://www.nowcoder.com/practice/e016ad9b7f0b45048c58a9f27ba618bf

有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。

输入：
[1,3,5,2,2],5,3
返回值：
2
"""

"""

参考： topK 问题
https://www.cnblogs.com/lesroad/p/10621074.html#_label3
https://blog.csdn.net/wjinjie/article/details/117480013

1. 暴力(快排，取第k个)
2. 大小堆
"""

import heapq


def find_kth(a, n, K):
    """
    NOTE: 原始 Go 实现有逻辑错误：它在堆 pop 时顺序不对，且构建的是大顶堆却
    没有正确利用它来求第 K 大。这里用 heapq 小顶堆重写为正确做法：
    维护一个大小为 K 的大顶堆（存前 K 大的元素），堆顶即为第 K 大。
    Python heapq 是小顶堆，所以取负值模拟大顶堆。
    """
    if len(a) < K or K == 0:
        return -1

    # 构建大顶堆（用负值模拟），存储最大的 K 个数据
    # 先全部取负入堆，再 pop 掉 n-K 个最小的（即最大的元素之外的那些）
    max_heap = [-x for x in a]
    heapq.heapify(max_heap)

    # pop 掉前 n-K 个，剩下的堆顶就是第 K 大
    for _ in range(n - K):
        heapq.heappop(max_heap)

    return -heapq.heappop(max_heap)

    # 1、因为是第K大，所以倒序快排，
    # 2、和中间点比较，可以每次只需处理一半的数据
    # mid, i = a[0], 1
    # first, tail = 0, n - 1
    # while first < tail:
    #     if a[i] < mid:
    #         a[i], a[tail] = a[tail], a[i]
    #         tail -= 1
    #     else:
    #         a[i], a[first] = a[first], a[i]
    #         first += 1
    #         i += 1
    # a[first] = mid
    #
    # if K - 1 == first:
    #     return mid
    # if K - 1 < first:
    #     return find_kth(a[:first], first, K)
    #
    # return find_kth(a[first + 1:], n - first - 1, K - first - 1)


if __name__ == "__main__":
    print(find_kth([1, 3, 5, 2, 2], 5, 3))
