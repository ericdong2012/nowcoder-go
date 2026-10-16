"""22. 括号生成（回溯 + 剪枝）
题目：生成 n 对括号的所有合法组合。
思路：回溯构建——右括号数必须 ≤ 左括号数；
      左 < n 可加左；右 < 左可加右。长度到 2n 收集。
复杂度：时间 O(4^n / √n)（卡特兰数），空间 O(n)。
"""

from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res, path = [], []

        def backtrack(opened: int, closed: int) -> None:
            if len(path) == 2 * n:
                res.append("".join(path))
                return
            if opened < n:
                path.append("(")
                backtrack(opened + 1, closed)
                path.pop()
            if closed < opened:               # 剪枝：右不能多于左
                path.append(")")
                backtrack(opened, closed + 1)
                path.pop()

        backtrack(0, 0)
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.generateParenthesis(3) == ["((()))", "(()())", "(())()", "()(())", "()()()"]
    assert s.generateParenthesis(1) == ["()"]
    print("59 OK")
