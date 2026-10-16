"""
BM63 跳台阶
https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4?tpId=295&tqId=23261

描述
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
要求：时间复杂度：O(n) ，空间复杂度： O(1)

示例1
输入：2
返回值：2

示例2
输入：7
返回值：21
"""


def jumpFloor(number):
    # write code here
    step1, step2 = 0, 1
    for i in range(number):
        # c = step1 + step2
        # step1 = step2
        # step2 = c
        step1, step2 = step2, step1 + step2
    return step2


if __name__ == "__main__":
    print(jumpFloor(7))
