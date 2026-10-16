"""31. 下一个排列（手动模拟）
题目：把数组变成字典序的下一个排列；已是最大则变成最小（升序）。
思路：从右找第一个升序对 (i, i+1)（nums[i] < nums[i+1]）；
      在右侧找最后一个 > nums[i] 的数交换；再把 i+1 之后整段反转。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        n = len(nums)
        i = n - 2
        while i >= 0 and nums[i] >= nums[i + 1]:    # 找第一个下降点
            i -= 1
        if i >= 0:
            j = n - 1
            while nums[j] <= nums[i]:               # 找最后一个大于 nums[i] 的
                j -= 1
            nums[i], nums[j] = nums[j], nums[i]
        nums[i + 1:] = reversed(nums[i + 1:])       # 后缀反转 → 最小化


if __name__ == "__main__":
    s = Solution()
    a = [1, 2, 3]
    s.nextPermutation(a)
    assert a == [1, 3, 2]
    b = [3, 2, 1]
    s.nextPermutation(b)
    assert b == [1, 2, 3]
    c = [1, 1, 5]
    s.nextPermutation(c)
    assert c == [1, 5, 1]
    d = [1, 5, 1]
    s.nextPermutation(d)
    assert d == [5, 1, 1]
    print("98 OK")
