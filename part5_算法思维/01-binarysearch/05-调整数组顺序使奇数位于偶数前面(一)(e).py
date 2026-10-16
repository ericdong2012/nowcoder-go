"""
NC77 调整数组顺序使奇数位于偶数前面(一)
https://www.nowcoder.com/practice/ef1f53ef31ca408cada5093c8780f44b

描述
输入一个长度为 n 整数数组，实现一个函数来调整该数组中数字的顺序，使得所有的奇数位于数组的前面部分，
所有的偶数位于数组的后面部分，并保证奇数和奇数，偶数和偶数之间的相对位置不变。

示例1
输入：[1,2,3,4]
返回值：[1,3,2,4]

示例2
输入：[2,4,6,5,7]
返回值：[5,7,2,4,6]
"""


def re_order_array(array):
    """write code here"""
    # 注意：原 Go 代码变量名是反的，odd 实际存偶数，even 实际存奇数，这里保持一致
    odd = []   # 实际收集偶数
    even = []  # 实际收集奇数
    for i in range(len(array)):
        if array[i] % 2 == 0:
            odd.append(array[i])
        else:
            even.append(array[i])

    even = even + odd
    return even
