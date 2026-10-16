"""32. 最长有效括号（DP / 栈）
题目：'(' ')' 组成的串中最长的合法括号子串长度。
思路（栈版）：栈底放"最后一个未匹配位置"哨兵；
      '(' 入栈下标；')' 弹栈——空了说明')'本身失配（重设哨兵），
      否则当前合法长度 = i - 新栈顶。
复杂度：时间 O(n)，空间 O(n)。
"""


class Solution:
    def longestValidParentheses(self, s: str) -> int:
        stack = [-1]                  # 哨兵：合法子串的左边界-1
        best = 0
        for i, ch in enumerate(s):
            if ch == "(":
                stack.append(i)
            else:
                stack.pop()
                if not stack:
                    stack.append(i)   # ')' 失配：成为新的边界哨兵
                else:
                    best = max(best, i - stack[-1])
        return best


if __name__ == "__main__":
    s = Solution()
    assert s.longestValidParentheses("(()") == 2
    assert s.longestValidParentheses(")()())") == 4
    assert s.longestValidParentheses("") == 0
    assert s.longestValidParentheses("()(())") == 6
    print("90 OK")
