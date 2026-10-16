"""
移除K位数字 后最小
https://www.nowcoder.com/practice/0fe685c8272d40f1b9785fedd2499c1c?tpId=196&tqId=39412&ru=/exam/oj

给定一个以字符串表示的数字 num 和一个数字 k ，从 num 中移除 k 位数字，使得剩下的数字最小。如果可以删除全部数字，则结果为 0。
1.num仅有数字组成
2.num是合法的数字，不含前导0
3.删除之后的num，请去掉前导0（不算在移除次数中）

保证 num 中仅包含 0~9 的十进制数

示例1
输入：
"1432219",3
返回值：
"1219"
说明：
移除 4 3 2 后剩下 1219

示例2
输入：
"10",1
返回值：
"0"

示例3
输入：
"100999",3
返回值：
"9"
"""


# 通过此方法搜索，还有一道类似的题，解法类似（一样）
def remove_knums(num, k):
    """write code here"""
    if len(num) <= k:
        return "0"

    stack = []
    for i in range(len(num)):
        # stack不为空，k不为0, 当前值小于栈中最后一个(前一个)，则移除栈中最后一个(弹栈); 因为题目要移除后结果最小，所有要将大的(前一个)弹栈移除掉
        # 单调栈
        # "1432219",3
        # stack: [1219]
        while len(stack) > 0 and k > 0 and num[i] < stack[-1]:
            stack.pop()
            k -= 1
        # 剪枝
        # 如果栈中没有元素并且当前值为'0', 则继续(比如：某个数字后一堆0)  比如: 100999
        # 此处不是必须
        # if len(stack) == 0 and num[i] == '0':
        #     continue
        stack.append(num[i])

    # 取k个数
    # k还没有找完，并且stack长度大于k, 需要从stack中截取, 舍弃后面的数据
    if k != 0 and len(stack) >= k:
        stack = stack[:len(stack) - k]
    # 清除前导0
    tmp = "".join(stack)
    tmp = tmp.lstrip("0")
    if len(tmp) == 0:
        return "0"

    return tmp


if __name__ == "__main__":
    result = remove_knums("100999", 3)
    print(result)
