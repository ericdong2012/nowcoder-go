"""75. 颜色分类（三路指针 / 荷兰国旗）
题目：只有 0/1/2 的数组原地排序，一遍扫描。
思路：三指针——p0 是 0 的右边界（放 0），p2 是 2 的左边界（放 2），
      cur 扫描；遇 0 与 p0 交换（p0 换来的一定是 1，cur 可前进），
      遇 2 与 p2 交换（换来的未知，cur 不前进）。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        p0, cur, p2 = 0, 0, len(nums) - 1
        while cur <= p2:
            if nums[cur] == 0:
                nums[p0], nums[cur] = nums[cur], nums[p0]
                p0 += 1
                cur += 1
            elif nums[cur] == 2:
                nums[p2], nums[cur] = nums[cur], nums[p2]
                p2 -= 1                  # 换来的数未检查，cur 原地
            else:
                cur += 1


if __name__ == "__main__":
    s = Solution()
    a = [2, 0, 2, 1, 1, 0]
    s.sortColors(a)
    assert a == [0, 0, 1, 1, 2, 2]
    b = [2, 0, 1]
    s.sortColors(b)
    assert b == [0, 1, 2]
    print("97 OK")
