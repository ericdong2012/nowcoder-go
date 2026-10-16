"""
NC68 跳台阶
https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4

描述
一只青蛙一次可以跳上 1 级台阶，也可以跳上 2 级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法
（先后次序不同算不同的结果）。

要求：时间复杂度：O(n) ，空间复杂度： O(1)

示例1
输入：2
返回值：2
说明：青蛙要跳上两级台阶有两种跳法，分别是：先跳一级，再跳一级或者直接跳两级。因此答案为 2

示例2
输入：7
返回值：21
"""

# 本质是递归问题
# 针对结果的个数，可以转化为斐波拉切问题
#
# 假设 f[i] 表示在第 i 个台阶上可能的方法数，逆向思维。如果我从第 n 个台阶进行下台阶，
# 下一步有 2 种可能，一种走到第 n-1 个台阶，一种是走到第 n-2 个台阶。
# 所以 f[n] = f[n-1] + f[n-2]。那么初始条件，f[0] = f[1] = 1。


def jump_floor(number):
    """write code here"""
    # if number < 3:
    #     return number
    step1, step2 = 0, 1
    for _ in range(number):
        # c = step1 + step2
        # step1 = step2
        # step2 = c
        step1, step2 = step2, step1 + step2
    return step2


if __name__ == "__main__":
    floor = jump_floor(1)
    print(floor)
