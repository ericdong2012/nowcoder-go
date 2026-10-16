"""46. 全排列（回溯）
题目：不含重复数字的数组的所有全排列。
思路：回溯——每层从未用过的数里选一个；用 used 标记，进入递归前标记、
      返回后撤销。
复杂度：时间 O(n·n!)，空间 O(n)。
"""

from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res, path = [], []
        used = [False] * len(nums)

        def backtrack() -> None:
            if len(path) == len(nums):
                res.append(path[:])          # 快照
                return
            for i, x in enumerate(nums):
                if used[i]:
                    continue
                used[i] = True
                path.append(x)
                backtrack()
                path.pop()                   # 撤销
                used[i] = False

        backtrack()
        return res


if __name__ == "__main__":
    s = Solution()
    assert sorted(s.permute([1, 2, 3])) == sorted([[1, 2, 3], [1, 3, 2], [2, 1, 3],
                                                   [2, 3, 1], [3, 1, 2], [3, 2, 1]])
    assert s.permute([0, 1]) == [[0, 1], [1, 0]]
    print("55 OK")
