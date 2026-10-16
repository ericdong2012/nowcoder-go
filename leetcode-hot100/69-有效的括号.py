"""20. 有效的括号（栈）
题目：'()[]{}' 组成的字符串是否有效（正确闭合与嵌套顺序）。
思路：栈——左括号入栈；右括号必须与栈顶配对并弹出；结束时栈空。
复杂度：时间 O(n)，空间 O(n)。
"""


class Solution:
    def isValid(self, s: str) -> bool:
        pairs = {")": "(", "]": "[", "}": "{"}
        stack = []
        for ch in s:
            if ch in "([{":
                stack.append(ch)
            elif not stack or stack.pop() != pairs[ch]:
                return False
        return not stack


if __name__ == "__main__":
    s = Solution()
    assert s.isValid("()[]{}") is True
    assert s.isValid("(]") is False
    assert s.isValid("([)]") is False
    assert s.isValid("{[]}") is True
    assert s.isValid("(") is False
    print("69 OK")
