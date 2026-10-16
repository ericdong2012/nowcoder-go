"""
题目大意
给出一个字符串数组，要求对字符串数组里面有 Anagrams 关系的字符串进行分组。Anagrams 关系是指
两个字符串的字符完全相同，顺序不同，两者是由排列组合组成。

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]

输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

解题思路
这道题可以将每个字符串都排序，排序完成以后，相同 Anagrams 的字符串必然排序结果一样。
把排序以后的字符串当做 key 存入到 map 中。遍历数组以后，就能得到一个 map，key 是排序以后的字符串，
value 对应的是这个排序字符串以后的 Anagrams 字符串集合。最后再将这些 value 对应的字符串数组输出即可。

https://www.nowcoder.com/practice/68bda87f91664623a4213abb00e41a5a
"""
from collections import defaultdict


def group_anagrams(strs):
    record = defaultdict(list)
    for s in strs:
        # sByte = sorted(s)
        # record["".join(sByte)].append(str)
        # 实际结果有点出入，就是字母的排序上
        key = "".join(sorted(s))
        record[key].append(s)

    res = []
    for v in record.values():
        res.append(v)
    return res
