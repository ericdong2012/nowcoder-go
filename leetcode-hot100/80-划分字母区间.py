"""763. 划分字母区间（贪心 + 最后出现位置）
题目：切成尽可能多的段，使每个字母只出现在一段中，返回各段长度。
思路：记录每个字母最后出现的位置；扫描时用"当前段内字母的最远末位"
      扩张段边界；到达边界即切一刀。
复杂度：时间 O(n)，空间 O(26)。
"""

from typing import List


class Solution:
    def partitionLabels(self, s: str) -> List[int]:
        last = {ch: i for i, ch in enumerate(s)}
        res = []
        start = end = 0
        for i, ch in enumerate(s):
            end = max(end, last[ch])
            if i == end:                    # 本段所有字母的最远末位已到
                res.append(end - start + 1)
                start = i + 1
        return res


if __name__ == "__main__":
    s = Solution()
    assert s.partitionLabels("ababcbacadefegdehijhklij") == [9, 7, 8]
    assert s.partitionLabels("eccbbbbdec") == [10]
    print("80 OK")
