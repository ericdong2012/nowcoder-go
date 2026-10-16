"""438. 找到字符串中所有字母异位词（滑动窗口+计数）
题目：在 s 中找 p 的所有异位词的起始下标。
思路：固定窗口大小 len(p) 的滑动窗口，比较窗口内字符计数与 p 的计数。
      优化：维护 diff 计数而非每次全量比较，O(26) 检查。
复杂度：时间 O(n·26)，空间 O(26)。
"""

from typing import List
from collections import Counter


class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        m, n = len(p), len(s)
        if m > n:
            return []
        need = Counter(p)
        window = Counter(s[:m])
        res = [0] if window == need else []
        for i in range(m, n):
            window[s[i]] += 1                     # 右边进
            left_ch = s[i - m]
            window[left_ch] -= 1                  # 左边出
            if window[left_ch] == 0:
                del window[left_ch]               # 计数为0要删，否则 == 比较失败
            if window == need:
                res.append(i - m + 1)
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.findAnagrams("cbaebabacd", "abc") == [0, 6]
    assert s.findAnagrams("abab", "ab") == [0, 1, 2]
    print("09 OK")
