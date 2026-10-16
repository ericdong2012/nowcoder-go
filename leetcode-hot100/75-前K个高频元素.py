"""347. 前 K 个高频元素（哈希 + 堆/桶排序）
题目：出现频率前 k 高的元素。
思路 A：计数 + 小顶堆维护前 k，O(n log k)。
思路 B（更优）：计数 + 桶排序（频率→数字列表），O(n)。
复杂度：桶版时间 O(n)，空间 O(n)。
"""

from typing import List
from collections import Counter


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = Counter(nums)
        buckets = [[] for _ in range(len(nums) + 1)]   # 下标=频率
        for num, freq in count.items():
            buckets[freq].append(num)
        res = []
        for freq in range(len(buckets) - 1, 0, -1):    # 从高频往低收
            for num in buckets[freq]:
                res.append(num)
                if len(res) == k:
                    return res
        return res


if __name__ == "__main__":
    s = Solution()
    assert set(s.topKFrequent([1, 1, 1, 2, 2, 3], 2)) == {1, 2}
    assert s.topKFrequent([1], 1) == [1]
    print("75 OK")
