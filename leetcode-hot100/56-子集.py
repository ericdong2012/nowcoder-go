"""78. 子集（回溯）
题目：返回数组的所有子集（幂集），元素互不相同。
思路：回溯——每个元素"选或不选"；或每层从 start 开始收集当前 path。
复杂度：时间 O(n·2^n)，空间 O(n)。
"""

from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []

        def backtrack(start: int, path: List[int]) -> None:
            res.append(path[:])              # 每个节点都是一个子集
            for i in range(start, len(nums)):
                path.append(nums[i])
                backtrack(i + 1, path)
                path.pop()

        backtrack(0, [])
        return res


if __name__ == "__main__":
    s = Solution()
    assert sorted(s.subsets([1, 2, 3])) == sorted(
        [[], [1], [2], [3], [1, 2], [1, 3], [2, 3], [1, 2, 3]])
    assert s.subsets([]) == [[]]
    print("56 OK")
