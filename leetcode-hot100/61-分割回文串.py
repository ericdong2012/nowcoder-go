"""131. 分割回文串（回溯 + 预处理）
题目：把 s 分割成若干回文子串，返回所有方案。
思路：回溯切割点；为 O(1) 判回文，先 DP 预处理 is_pal[i][j]
     （is_pal[i][j] = s[i]==s[j] 且 is_pal[i+1][j-1]）。
复杂度：时间 O(n·2^n)，空间 O(n²)。
"""

from typing import List


class Solution:
    def partition(self, s: str) -> List[List[str]]:
        n = len(s)
        pal = [[True] * n for _ in range(n)]
        for i in range(n - 1, -1, -1):            # 按长度递增填表
            for j in range(i + 1, n):
                pal[i][j] = s[i] == s[j] and pal[i + 1][j - 1]

        res, path = [], []

        def backtrack(start: int) -> None:
            if start == n:
                res.append(path[:])
                return
            for end in range(start, n):
                if pal[start][end]:
                    path.append(s[start:end + 1])
                    backtrack(end + 1)
                    path.pop()

        backtrack(0)
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.partition("aab") == [["a", "a", "b"], ["aa", "b"]]
    assert s.partition("a") == [["a"]]
    print("61 OK")
