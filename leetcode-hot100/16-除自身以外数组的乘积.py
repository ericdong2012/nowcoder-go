"""238. 除自身以外数组的乘积（前缀×后缀）
题目：answer[i] = 除 nums[i] 之外所有元素乘积；不许用除法。
思路：answer[i] = (i 左边全部乘积) × (i 右边全部乘积)。
      一遍从左填前缀，一遍从右乘后缀，空间 O(1)（不算结果）。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List


class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        answer = [1] * n
        prefix = 1
        for i in range(n):                 # answer[i] = 左边乘积
            answer[i] = prefix
            prefix *= nums[i]
        suffix = 1
        for i in range(n - 1, -1, -1):     # 再乘上右边乘积
            answer[i] *= suffix
            suffix *= nums[i]
        return answer


if __name__ == "__main__":
    s = Solution()
    assert s.productExceptSelf([1, 2, 3, 4]) == [24, 12, 8, 6]
    assert s.productExceptSelf([-1, 1, 0, -3, 3]) == [0, 0, 9, 0, 0]
    print("16 OK")
