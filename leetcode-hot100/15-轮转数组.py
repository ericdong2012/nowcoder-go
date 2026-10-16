"""189. 轮转数组（数组翻转）
题目：数组向右轮转 k 位，原地 O(1) 空间。
思路：三次翻转——整体翻转，再翻转前 k 个、后 n-k 个。
      k %= n 处理 k > n。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        n = len(nums)
        k %= n

        def reverse(l: int, r: int) -> None:
            while l < r:
                nums[l], nums[r] = nums[r], nums[l]
                l += 1
                r -= 1

        reverse(0, n - 1)
        reverse(0, k - 1)
        reverse(k, n - 1)


if __name__ == "__main__":
    s = Solution()
    a = [1, 2, 3, 4, 5, 6, 7]
    s.rotate(a, 3)
    assert a == [5, 6, 7, 1, 2, 3, 4]
    b = [-1]
    s.rotate(b, 2)        # k > n 情形
    assert b == [-1]
    print("15 OK")
