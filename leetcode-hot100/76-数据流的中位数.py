"""295. 数据流的中位数（对顶堆）
题目：addNum 加数，findMedian 返回中位数。
思路：大顶堆存较小一半（存负数模拟），小顶堆存较大一半；
      保持 |len(大顶堆) - len(小顶堆)| ≤ 1。中位数取两堆顶。
复杂度：addNum O(log n)，findMedian O(1)，空间 O(n)。
"""

import heapq


class MedianFinder:
    def __init__(self):
        self.small = []     # 大顶堆（存负数）：较小的一半
        self.large = []     # 小顶堆：较大的一半

    def addNum(self, num: int) -> None:
        heapq.heappush(self.small, -num)
        # 保证 small 堆顶 <= large 堆顶
        heapq.heappush(self.large, -heapq.heappop(self.small))
        # 平衡：small 允许多 1 个
        if len(self.large) > len(self.small):
            heapq.heappush(self.small, -heapq.heappop(self.large))

    def findMedian(self) -> float:
        if len(self.small) > len(self.large):
            return float(-self.small[0])
        return (-self.small[0] + self.large[0]) / 2


if __name__ == "__main__":
    mf = MedianFinder()
    mf.addNum(1)
    mf.addNum(2)
    assert mf.findMedian() == 1.5
    mf.addNum(3)
    assert mf.findMedian() == 2.0
    mf.addNum(4)
    mf.addNum(5)
    assert mf.findMedian() == 3.0
    print("76 OK")
