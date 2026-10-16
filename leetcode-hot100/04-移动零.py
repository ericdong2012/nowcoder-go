"""283. 移动零（双指针）
题目：把数组中所有 0 移到末尾，保持非零元素相对顺序，原地操作。
思路：慢指针 j 指向"下一个非零元素应放的位置"；快指针扫到非零就与 j 交换。
     扫完后 j 之后全是 0。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        j = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[i], nums[j] = nums[j], nums[i]
                j += 1


if __name__ == "__main__":
    s = Solution()
    a = [0, 1, 0, 3, 12]
    s.moveZeroes(a)
    assert a == [1, 3, 12, 0, 0]
    b = [0]
    s.moveZeroes(b)
    assert b == [0]
    print("04 OK")
