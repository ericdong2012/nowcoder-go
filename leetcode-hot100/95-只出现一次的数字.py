"""136. 只出现一次的数字（异或）
题目：其他数都出现两次，只有一个出现一次，找它。
思路：异或性质 a^a=0、a^0=a、交换律——全部异或，成对抵消。
复杂度：时间 O(n)，空间 O(1)。
"""

from typing import List
from functools import reduce
import operator


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        return reduce(operator.xor, nums)


if __name__ == "__main__":
    s = Solution()
    assert s.singleNumber([2, 2, 1]) == 1
    assert s.singleNumber([4, 1, 2, 1, 2]) == 4
    assert s.singleNumber([1]) == 1
    print("95 OK")
