"""
NC1 大数加法
https://www.nowcoder.com/practice/11ae12e8c6fe48f883cad618c2e81475

描述
以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回
字符串仅由 '0'~'9' 构成
时间复杂度 O(n)

示例1
输入："1","99"
返回值："100"
说明：1+99=100

示例2
输入："114514",""
返回值："114514"
"""

# 模拟法或者栈
#
#   设定 i、j 两指针分别指向 s、t 尾部，模拟人工加法；
#   计算进位：计算 carry = tmp // 10，代表当前位相加是否产生进位；
#   添加当前位：计算 tmp = n1 + n2 + carry，并将当前位 tmp % 10 添加至 res 头部；
#   索引溢出处理：当指针 i 或 j 走过数字首部后，给 n1、n2 赋值为 0，相当于给 s、t 中长度较短的数字前面填 0。
#   当遍历完 s、t 后跳出循环，并根据 carry 值决定是否在头部添加进位 1，最终返回 res 即可。


def solve1(s, t):
    """write code here"""
    res = ""
    i, j = len(s) - 1, len(t) - 1
    carry = 0
    # 长度不一，会有一方在后续循环中变成负数
    while i >= 0 or j >= 0:
        # >= 0
        if i >= 0:
            n1 = ord(s[i]) - ord('0')
        else:
            # i 是负数，每个数位数不同，位数走完，需要用 0 填充
            n1 = 0

        if j >= 0:
            n2 = ord(t[j]) - ord('0')
        else:
            n2 = 0
        # 某一位的两数相加加上进位
        temp = n1 + n2 + carry
        # 超过 10 才有进位
        carry = temp // 10
        # str(temp % 10) 在前是因为从个位到十位到百位
        res = str(temp % 10) + res
        j -= 1
        i -= 1

    if carry != 0:
        return "1" + res

    return res


if __name__ == "__main__":
    s = solve1("1", "99")
    print(s)
