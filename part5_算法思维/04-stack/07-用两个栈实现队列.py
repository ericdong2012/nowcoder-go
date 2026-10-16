"""
NC76 用两个栈实现队列
https://www.nowcoder.com/practice/54275ddae22f475981afa2244dd448c6

描述
用两个栈来实现一个队列，使用 n 个元素来完成 n 次在队列尾部插入整数(push)和 n 次在队列头部删除整数(pop)的功能。
队列中的元素为 int 类型。保证操作合法，即保证 pop 操作时队列内已有元素。

数据范围：n <= 1000
要求：存储 n 个元素的空间复杂度为 O(n)，插入与删除的时间复杂度都是 O(1)

示例1
输入：["PSH1","PSH2","POP","POP"]
返回值：1,2
"""

# 头在前，尾在后
stack1 = []

# 尾在前，头在后
stack2 = []


def push1(node):
    stack1.append(node)


def pop1():
    # 如果 stack2 存在，从 stack2 中取
    if len(stack2) != 0:
        tmp = stack2[-1]
        stack2.pop()
        return tmp
    # 如果 stack2 为空，stack1 存在
    if len(stack1) != 0:
        # 获取 stack1 头部的元素
        tmp = stack1[0]
        # 将 stack1 中的元素从尾部取出，添加到 stack2 中，作为头部
        for i in range(len(stack1) - 1, 0, -1):
            stack2.append(stack1[i])
        # 清空 stack1
        stack1.clear()
        # 返回 stack1 头部的元素
        return tmp
    # 如果 stack1、stack2 都不存在
    return -1


"""
class Solution:
    def __init__(self):
        self.A = []
        self.B = []
    def push(self, num):
        self.A.append(num)
    def pop(self):
        if self.B:
            return self.B.pop()
        if not self.A:
            return None
        else:
            while self.A:
                self.B.append(self.A.pop())
            return self.B.pop()
"""
