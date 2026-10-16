"""
NC103 反转字符串
https://www.nowcoder.com/practice/c3a6afee325e472386a1c4eb1ef987f3

描述
写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过 1000）

数据范围：0 <= n <= 1000
要求：空间复杂度 O(n)，时间复杂度 O(n)

示例1
输入："abcd"
返回值："dcba"

示例2
输入：""
返回值：""
"""


def solve(s):
    """write code here"""
    res = ""
    for i in range(len(s) - 1, -1, -1):
        res += s[i]
    return res
