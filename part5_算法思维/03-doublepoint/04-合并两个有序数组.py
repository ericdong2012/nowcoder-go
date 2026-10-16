"""
BM87 合并两个有序的数组
https://www.nowcoder.com/practice/89865d4375634fc484f3a24b7fe65665

描述
给出一个有序的整数数组 A 和有序的整数数组 B，请将数组 B 合并到数组 A 中，变成一个有序的升序数组

注意：
1.保证 A 数组有足够的空间存放 B 数组的元素，A 和 B 中初始的元素数目分别为 m 和 n，A 的数组空间大小为 m+n
2.不要返回合并的数组，将数组 B 的数据合并到 A 里面就好了，且后台会自动将合并后的数组 A 的内容打印出来
3.A 数组在 [0,m-1] 的范围也是有序的

示例1
输入：[4,5,6],[1,2,3]
返回值：[1,2,3,4,5,6]
"""


def merge(a, m, b, n):
    """write code here"""
    # 这个代码，线上跑报错
    index = m + n
    result = [0] * (m + n)
    while m > 0 and n > 0:
        # 先比较最后一位
        if a[m - 1] >= b[n - 1]:
            result[index - 1] = a[m - 1]
            m -= 1
        else:
            result[index - 1] = b[n - 1]
            n -= 1
        index -= 1
    # B 数组依然有剩余
    while n > 0 and index > 0:
        result[index - 1] = b[n - 1]
        n -= 1
        index -= 1
    while m > 0 and index > 0:
        result[index - 1] = a[m - 1]
        m -= 1
        index -= 1

    print(result)


if __name__ == "__main__":
    # merge([2, 5, 6], 3, [1, 2, 3], 3)
    merge([], 0, [1], 1)
