"""49. 字母异位词分组（哈希）
题目：把互为字母异位词的字符串分到一组。
思路：排序后的字符串作为异位词的"指纹"（anagram 排序结果相同），按指纹分组。
复杂度：时间 O(n·k·log k)（k 为单词长度），空间 O(n·k)。
"""

from typing import List
from collections import defaultdict


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(list)
        for s in strs:
            key = "".join(sorted(s))
            groups[key].append(s)
        return list(groups.values())


if __name__ == "__main__":
    s = Solution()
    out = s.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
    assert sorted([sorted(g) for g in out]) == sorted(
        [sorted(g) for g in [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]])
    assert s.groupAnagrams([""]) == [[""]]
    print("02 OK")
