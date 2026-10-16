"""17. 电话号码的字母组合（回溯/笛卡尔积）
题目：数字串 2-9 对应手机九键字母，求所有可能的字母组合。
思路：回溯逐位展开；数字→字母映射表。
复杂度：时间 O(3^m · 4^n)，空间 O(m+n)。
"""

from typing import List


class Solution:
    LETTERS = {
        "2": "abc", "3": "def", "4": "ghi", "5": "jkl",
        "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz",
    }

    def letterCombinations(self, digits: str) -> List[str]:
        if not digits:
            return []
        res, path = [], []

        def backtrack(i: int) -> None:
            if i == len(digits):
                res.append("".join(path))
                return
            for ch in self.LETTERS[digits[i]]:
                path.append(ch)
                backtrack(i + 1)
                path.pop()

        backtrack(0)
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.letterCombinations("23") == ["ad", "ae", "af", "bd", "be", "bf",
                                          "cd", "ce", "cf"]
    assert s.letterCombinations("") == []
    print("57 OK")
