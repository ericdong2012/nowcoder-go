"""
NC90 包含 min 函数的栈
https://www.nowcoder.com/practice/4c776177d2c04c2494f2555c9fcc1e49

描述
定义栈的数据结构，请在该类型中实现一个能够得到栈中所含最小元素的 min 函数，
输入操作时保证 pop、top 和 min 函数操作时，栈中一定有元素。

此栈包含的方法有：
push(value): 将 value 压入栈中
pop():       弹出栈顶元素
top():       获取栈顶元素
min():       获取栈中最小元素

数据范围：操作数量满足 0 <= n <= 300，输入的元素满足 |val| <= 10000
进阶：栈的各个操作的时间复杂度是 O(1)，空间复杂度是 O(n)

示例:
输入:    ["PSH-1","PSH2","MIN","TOP","POP","PSH1","TOP","MIN"]
输出:    -1,2,1,-1
"""

# 后进先出
L = []


def push(node):
    """write code here"""
    L.append(node)


def pop():
    """write code here"""
    del L[-1]


def top():
    """write code here"""
    return L[-1]


def min():
    """write code here"""
    m = L[0]
    for v in L:
        if v < m:
            m = v
    return m
