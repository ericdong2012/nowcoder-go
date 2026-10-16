"""
NC141 判断是否为回文字符串
https://www.nowcoder.com/practice/e297fdd8e9f543059b0b5f05f3a7f3b2

描述
给定一个长度为 n 的字符串，请编写一个函数判断该字符串是否回文。如果是回文请返回 true，否则返回 false。
字符串回文指该字符串正序与其逆序逐字符一致。

要求：空间复杂度 O(1)，时间复杂度 O(n)

示例1
输入："absba"
返回值：true

示例2
输入："ranko"
返回值：false

备注：
字符串长度不大于 1000000，且仅由小写字母组成
"""


def judge(s):
    """write code here"""
    left, right = 0, len(s) - 1
    while left <= right:
        if s[left] != s[right]:
            return False
        else:
            left += 1
            right -= 1
    return True
