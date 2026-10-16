"""128. 最长连续序列（哈希/集合）
题目：无序数组中最长连续数字序列长度，要求 O(n)。
思路：全部入 set；只从"序列起点"（x-1 不在 set 里）向右数。
      非起点直接跳过，保证每个数最多被访问两次 → O(n)。
复杂度：时间 O(n)，空间 O(n)。
"""

from typing import List


class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        num_set = set(nums)
        best = 0
        for x in num_set:
            if x - 1 in num_set:          # 不是起点，跳过
                continue
            length = 1
            while x + length in num_set:
                length += 1
            best = max(best, length)
        return best


if __name__ == "__main__":
    s = Solution()
    assert s.longestConsecutive([100, 4, 200, 1, 3, 2]) == 4
    assert s.longestConsecutive([0, 3, 7, 2, 5, 8, 4, 6, 0, 1]) == 9
    assert s.longestConsecutive([]) == 0
    print("03 OK")
