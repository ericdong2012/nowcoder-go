"""3. 无重复字符的最长子串（滑动窗口）
题目：不含重复字符的最长子串长度。
思路：右指针扩张；出现重复时左指针收缩（用 set/last 记录）。
      优化：last[ch] 记字符最后位置，左指针直接跳到"重复字符的下一位"。
复杂度：时间 O(n)，空间 O(字符集)。
"""

from typing import List


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        last = {}          # 字符 -> 最后出现的下标
        left = 0
        best = 0
        for right, ch in enumerate(s):
            if ch in last and last[ch] >= left:
                left = last[ch] + 1        # 直接跳过重复
            last[ch] = right
            best = max(best, right - left + 1)
        return best


if __name__ == "__main__":
    s = Solution()
    assert s.lengthOfLongestSubstring("abcabcbb") == 3
    assert s.lengthOfLongestSubstring("bbbbb") == 1
    assert s.lengthOfLongestSubstring("pwwkew") == 3
    assert s.lengthOfLongestSubstring("") == 0
    print("08 OK")
