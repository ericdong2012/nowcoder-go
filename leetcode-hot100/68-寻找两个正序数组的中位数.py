"""4. 寻找两个正序数组的中位数（二分，O(log(m+n))）
题目：两个升序数组的中位数，要求对数复杂度。
思路：在较短的数组上二分分割线 i，j = (m+n+1)//2 - i；
      分割合法 ⇔ 左半最大 ≤ 右半最小（跨数组比较四个边界值）。
      分割线确定后中位数即左半最大（奇数）或左最大/右最小的均值（偶数）。
复杂度：时间 O(log min(m,n))，空间 O(1)。
"""

from typing import List


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        if len(nums1) > len(nums2):           # 保证在短的数组上二分
            nums1, nums2 = nums2, nums1
        m, n = len(nums1), len(nums2)
        half = (m + n + 1) // 2               # 左半应有的大小
        lo, hi = 0, m
        while lo <= hi:
            i = (lo + hi) // 2                # nums1 取左半 i 个
            j = half - i                      # nums2 取左半 j 个
            l1 = nums1[i - 1] if i > 0 else float("-inf")
            r1 = nums1[i] if i < m else float("inf")
            l2 = nums2[j - 1] if j > 0 else float("-inf")
            r2 = nums2[j] if j < n else float("inf")
            if l1 <= r2 and l2 <= r1:         # 分割合法
                if (m + n) % 2:
                    return float(max(l1, l2))
                return (max(l1, l2) + min(r1, r2)) / 2
            if l1 > r2:
                hi = i - 1                    # nums1 拿多了
            else:
                lo = i + 1                    # nums1 拿少了
        raise ValueError("输入非升序")


if __name__ == "__main__":
    s = Solution()
    assert s.findMedianSortedArrays([1, 3], [2]) == 2.0
    assert s.findMedianSortedArrays([1, 2], [3, 4]) == 2.5
    assert s.findMedianSortedArrays([], [1]) == 1.0
    assert s.findMedianSortedArrays([0, 0], [0, 0]) == 0.0
    print("68 OK")
