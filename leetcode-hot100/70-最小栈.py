"""155. 最小栈（辅助栈）
题目：支持 push/pop/top/min 的栈，min O(1)。
思路：辅助栈与主栈同步压入"当前最小值"；弹出时同步弹出。
复杂度：所有操作 O(1)，空间 O(n)。
"""


class MinStack:
    def __init__(self):
        self.stack = []
        self.min_stack = []           # min_stack[i] = 前 i+1 个元素的最小值

    def push(self, val: int) -> None:
        self.stack.append(val)
        cur_min = min(val, self.min_stack[-1]) if self.min_stack else val
        self.min_stack.append(cur_min)

    def pop(self) -> None:
        self.stack.pop()
        self.min_stack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        return self.min_stack[-1]


if __name__ == "__main__":
    st = MinStack()
    st.push(-2)
    st.push(0)
    st.push(-3)
    assert st.getMin() == -3
    st.pop()
    assert st.top() == 0
    assert st.getMin() == -2
    print("70 OK")
