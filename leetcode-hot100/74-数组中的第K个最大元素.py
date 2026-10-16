"""215. 数组中的第K个最大元素（快速选择）
题目：第 k 大的元素。
思路：快速选择（partition 后看基准落点与第 k 大的位置关系，只递归一侧）；
      面试常要求手写 partition。堆排（heapq.nlargest）作对比。
复杂度：快速选择平均 O(n) 最坏 O(n²)；堆 O(n log k)。
"""

import random
from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        target = len(nums) - k          # 升序排序后目标下标
        l, r = 0, len(nums) - 1
        while True:
            pivot = nums[random.randint(l, r)]
            i, j = l, r
            while i <= j:
                while nums[i] < pivot:
                    i += 1
                while nums[j] > pivot:
                    j -= 1
                if i <= j:
                    nums[i], nums[j] = nums[j], nums[i]
                    i += 1
                    j -= 1
            if target <= j:
                r = j
            elif target >= i:
                l = i
            else:
                return nums[j + 1]      # j < target < i，该位置即基准


if __name__ == "__main__":
    s = Solution()
    assert s.findKthLargest([3, 2, 1, 5, 6, 4], 2) == 5
    assert s.findKthLargest([3, 2, 3, 1, 2, 4, 5, 5, 6], 4) == 4
    assert s.findKthLargest([1], 1) == 1
    print("74 OK")
