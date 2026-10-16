"""56. 合并区间（排序+扫描）
题目：合并所有重叠的区间。
思路：按左端点排序；当前区间与结果最后一个比较——起点 <= 上一终点就合并
     （终点取 max），否则新开一组。
复杂度：时间 O(n log n)，空间 O(n)（结果）。
"""

from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x: x[0])
        merged = []
        for start, end in intervals:
            if merged and start <= merged[-1][1]:
                merged[-1][1] = max(merged[-1][1], end)   # 注意取 max：[1,10],[2,3]
            else:
                merged.append([start, end])
        return merged


if __name__ == "__main__":
    s = Solution()
    assert s.merge([[1, 3], [2, 6], [8, 10], [15, 18]]) == [[1, 6], [8, 10], [15, 18]]
    assert s.merge([[1, 4], [4, 5]]) == [[1, 5]]
    assert s.merge([[1, 10], [2, 3]]) == [[1, 10]]
    print("14 OK")
