"""5. 最长回文子串（中心扩展）
题目：s 中最长的回文子串。
思路：对每个位置做两种中心扩展：奇数长（单中心）与偶数长（双中心），
      取最长。O(n²) 但常数小、好写；Manacher 可到 O(n)（面试不常要求）。
复杂度：时间 O(n²)，空间 O(1)。
"""


class Solution:
    def longestPalindrome(self, s: str) -> str:
        def expand(l: int, r: int) -> tuple:
            while l >= 0 and r < len(s) and s[l] == s[r]:
                l -= 1
                r += 1
            return l + 1, r                 # 回文区间 [l+1, r)

        best_l, best_r = 0, 0
        for i in range(len(s)):
            l1, r1 = expand(i, i)           # 奇数长
            l2, r2 = expand(i, i + 1)       # 偶数长
            if r1 - l1 > best_r - best_l:
                best_l, best_r = l1, r1
            if r2 - l2 > best_r - best_l:
                best_l, best_r = l2, r2
        return s[best_l:best_r]


if __name__ == "__main__":
    s = Solution()
    assert s.longestPalindrome("babad") in ("bab", "aba")
    assert s.longestPalindrome("cbbd") == "bb"
    assert s.longestPalindrome("a") == "a"
    assert s.longestPalindrome("aacabdkacaa") in ("aca",)
    print("94 OK")
