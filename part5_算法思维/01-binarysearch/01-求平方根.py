"""
NC32 求平方根
https://www.nowcoder.com/practice/09fbfb16140b40499951f55113f2166c

描述
实现函数 int sqrt(int x).
计算并返回 x 的平方根（向下取整）

示例1
输入：2
返回值：1

示例2
输入：2143195649
返回值：46294
"""


def sqrt(x):
    """write code here"""
    if x == 0 or x == 1:
        return x
    mid = x // 2
    # print(mid)

    while mid * mid > x:
        mid = mid // 2

    while mid * mid <= x:
        mid = mid + 1

    return mid - 1


if __name__ == "__main__":
    print(sqrt(10))
    # print(sqrt(3))
