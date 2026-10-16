"""15. 三数之和（排序+双指针）
题目：找出所有不重复的三元组使 a+b+c==0。
思路：排序后固定 i，左右双指针找两数之和 == -nums[i]。
      跳过重复：i 跳过相同值；找到一组后 l/r 各自跳过相同值。
复杂度：时间 O(n²)，空间 O(log n)（排序栈）。
"""

from typing import List


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
        n = len(nums)
        for i in range(n - 2):
            if nums[i] > 0:                       # 最小数>0，后面不可能凑出 0
                break
            if i > 0 and nums[i] == nums[i - 1]:  # 跳过重复的固定数
                continue
            l, r = i + 1, n - 1
            while l < r:
                total = nums[i] + nums[l] + nums[r]
                if total < 0:
                    l += 1
                elif total > 0:
                    r -= 1
                else:
                    res.append([nums[i], nums[l], nums[r]])
                    while l < r and nums[l] == nums[l + 1]:
                        l += 1
                    while l < r and nums[r] == nums[r - 1]:
                        r -= 1
                    l += 1
                    r -= 1
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.threeSum([-1, 0, 1, 2, -1, -4]) == [[-1, -1, 2], [-1, 0, 1]]
    assert s.threeSum([0, 0, 0, 0]) == [[0, 0, 0]]
    assert s.threeSum([1, 2, -2, -1]) == []
    print("06 OK")
