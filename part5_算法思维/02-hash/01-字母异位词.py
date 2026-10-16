"""
题目大意
给出 2 个字符串 s 和 t，如果 t 中的字母在 s 中都存在，输出 true，否则输出 false

解题思路
这道题可以用打表的方式做。先把 s 中的每个字母都存在一个 26 个容量的数组里面，每个下标依次对应
26 个字母。s 中每个字母都对应表中一个字母，每出现一次就加 1。然后再扫字符串 t，每出现一个字母
就在表里面减一。如果都出现了，最终表里面的值肯定都是 0。最终判断表里面的值是否都是 0 即可，
有非 0 的数都输出 false
"""
from collections import defaultdict


# 解法一
def is_anagram(s, t):
    hash_map = defaultdict(int)
    for value in s:
        hash_map[value] += 1
    for value in t:
        hash_map[value] -= 1
    for value in hash_map.values():
        if value != 0:
            return False
    return True


# def is_anagram(s, t):
#     alphabet = [0] * 26
#     if len(s) != len(t):
#         return False
#     for ch in s:
#         alphabet[ord(ch) - ord('a')] += 1
#     for ch in t:
#         alphabet[ord(ch) - ord('a')] -= 1
#     for i in range(26):
#         if alphabet[i] != 0:
#             return False
#     return True


# 解法二
# def is_anagram1(s, t):
#     if s == "" and t == "":
#         return True
#     if s == "" or t == "":
#         return False
#     if len(s) != len(t):
#         return False
#     s_sorted = sorted(s)
#     t_sorted = sorted(t)
#     for i in range(len(s_sorted)):
#         if s_sorted[i] != t_sorted[i]:
#             return False
#     return True
