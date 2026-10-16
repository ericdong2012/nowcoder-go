"""39. 组合总和（回溯）
题目：candidates（无重复、可重复使用）中凑出和为 target 的所有组合。
思路：回溯；为避免重复组合，递归时传 start=i（可重复用自己，但不能回头）。
复杂度：时间 O(n^(target/min))，空间 O(target)。
"""

from typing import List


class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        res, path = [], []

        def backtrack(start: int, remain: int) -> None:
            if remain == 0:
                res.append(path[:])
                return
            if remain < 0:
                return
            for i in range(start, len(candidates)):
                path.append(candidates[i])
                backtrack(i, remain - candidates[i])   # i：可重复使用自己
                path.pop()

        backtrack(0, target)
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.combinationSum([2, 3, 6, 7], 7) == [[2, 2, 3], [7]]
    assert s.combinationSum([2, 3, 5], 8) == [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
    print("58 OK")
