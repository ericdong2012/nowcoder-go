"""76. 最小覆盖子串（滑动窗口）
题目：s 中包含 t 全部字符（含重复数）的最小子串。
思路：右扩窗口直到覆盖（need 计数），然后左收缩到最小；重复"扩张-收缩"。
复杂度：时间 O(n+|t|+26·n)，空间 O(26)。
"""

from collections import Counter


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        need = Counter(t)                 # t 中每个字符还差多少
        missing = len(t)                  # 总共还差多少字符
        left = 0
        best = (float("inf"), 0, 0)       # (长度, 左, 右)
        for right, ch in enumerate(s, 1):
            if need[ch] > 0:
                missing -= 1
            need[ch] -= 1
            if missing == 0:              # 已覆盖，开始左收缩
                while need[s[left]] < 0:  # 左端字符是多余的
                    need[s[left]] += 1
                    left += 1
                if right - left < best[0]:
                    best = (right - left, left, right)
                # 收缩掉一个必需字符，让 right 继续
                need[s[left]] += 1
                missing += 1
                left += 1
        return "" if best[0] == float("inf") else s[best[1]:best[2]]


if __name__ == "__main__":
    s = Solution()
    assert s.minWindow("ADOBECODEBANC", "ABC") == "BANC"
    assert s.minWindow("a", "a") == "a"
    assert s.minWindow("a", "aa") == ""
    print("12 OK")
