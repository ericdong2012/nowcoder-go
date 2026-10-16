"""287. 寻找重复数（Floyd 判圈 on 序列）
题目：n+1 个数都在 [1, n] 内，恰有一个重复，找出它；不改数组、O(1) 空间。
思路：把 nums[i] 当作"i 位置的下一个指针"，重复数 = 环入口。
      快慢指针（f=f[f], s=f[s]）先判相遇，再一个从头走找入口。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        slow = fast = nums[0]
        while True:                    # 判圈：快走2步慢走1步
            slow = nums[slow]
            fast = nums[nums[fast]]
            if slow == fast:
                break
        slow = nums[0]                 # 找入口：一个从头，一个从相遇点，同速
        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]
        return slow


if __name__ == "__main__":
    s = Solution()
    assert s.findDuplicate([1, 3, 4, 2, 2]) == 2
    assert s.findDuplicate([3, 1, 3, 4, 2]) == 3
    assert s.findDuplicate([2, 2, 2, 2]) == 2
    print("99 OK")
