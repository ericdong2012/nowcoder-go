"""
NC79 丑数
https://www.nowcoder.com/practice/6aa9e04fc3794f68acf8778237ba065b

描述
把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。例如 6、8 都是丑数，但 14 不是，
因为它包含质因子 7。习惯上我们把 1 当做是第一个丑数。求按从小到大的顺序的第 n 个丑数。

示例1
输入：7
返回值：8
"""

# 题目的意思是：一个数等于 2^x * 3^y * 5^z 相乘的结果
# 1, 2, 3, 4, 5, 6, 8, 9, 10, 12   7,11,13,14
# 传参是个数，结果是最后的那个值

# 最小堆
# step 1：第一个丑数 1 加入数组。
# step 2：使用 i、j、k 三个索引表示该数字有无被乘 2、乘 3、乘 5。
# step 3：后续继续找 n−1 个丑数，每次取当前丑数索引乘 2、乘 3、乘 5 的最小值加入数组，并计数。
# step 4：若是该丑数为相应索引乘上某个数字，则对应的索引往后一位。


def min3(a, b):
    if a < b:
        return a
    else:
        return b


def get_ugly_number(index):
    """write code here"""
    if index <= 1:
        return index
    num = [0] * index
    num[0] = 1
    # i、j、k 分别为三个队列的指针，temp 从队列头选出来的最小数
    i = 0
    j = 0
    k = 0
    temp = 0
    for ii in range(1, index):
        # 第一次是 2、3、5 比较，得到最小的是 2
        """
        画图
        ii = 1
        i  = 1
        num = [1, 2]
        """
        # 最小堆，选出三个队列头最小的数
        a, b, c = num[i] * 2, num[j] * 3, num[k] * 5
        temp = min3(a, min3(b, c))

        # 这三个 if 有可能进入一个或者多个，进入多个是三个队列头最小的数有多个的情况
        if temp == a:
            i += 1
        if temp == b:
            j += 1
        if temp == c:
            k += 1

        num[ii] = temp
        """
        temp = min(2, 3, 5) = 2, i=1, j=0, k=0, num[1] = 2
        temp = min(4, 3, 5) = 3, i=1, j=1, k=0, num[2] = 3
        temp = min(4, 6, 5) = 4, i=1, j=1, k=1, num[3] = 4
        temp = min(4, 6, 10) = 4, i=2, j=1, k=1, num[4] = 4
        """

    return num[index - 1]
    # return temp


if __name__ == "__main__":
    solution = get_ugly_number(7)
    print(solution)
