# -*- coding: utf-8 -*-
from typing import List

"""
https://leetcode-cn.com/problems/power-of-two/

231. 2的幂

给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

示例 1:
输入: 1
输出: true
解释: 2^0= 1

示例 2:
输入: 16
输出: true
解释: 2^4= 16

示例 3:
输入: 218
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-two
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
"""


class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        # method1: 迭代，如果有结果不是0， 不是2的幂

        # method2: 判断最右边的1
        if n == 0:
            return False
        # n; 8     0 1000
        # -n: -8   1 1000    (1 0111 + 1)

        # 7 & -7
        # 0000 0111
        # 1111 1001 (00000111这个结果是7的结果，然后各位取反 11111000 末尾加1 11111001这个就是-7的二进制数了)
        # 0000 0111
        return n & (-n) == n

        # method3: 去除最右边的1
        # https://leetcode-cn.com/problems/power-of-two/solution/power-of-two-er-jin-zhi-ji-jian-by-jyd/
        # if n == 0 :
        #     return False

        # 8      01000
        # 7      00111
        # 8 & 7  00000

        # 7      111
        # 6      110
        # 7 & 6  110 不等与0
        # return n & (n -1 ) == 0


s = Solution()
print(s.isPowerOfTwo(16))
