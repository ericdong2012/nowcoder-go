"""
387. 字符串中的第一个唯一字符 找出字符串中第一个只出现一次的字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:
s = "leetcode"
返回 0

s = "loveleetcode"
返回 2

注意事项：您可以假定该字符串只包含小写字母。

https://www.nowcoder.com/practice/e896d0f82f1246a3aa7b232ce38029d4
"""


def first_uniq_char(s):
    # 正确解法：使用哈希表记录每个字符出现的次数
    # 注意：原 Go 版本用 "abcdefg...z" 的顺序遍历再比较 find/last_index，
    # 这会返回「字母表中最先出现的唯一字符」而非「字符串中最先出现的唯一字符」，
    # 例如 "leetcode" 会错误地返回 4(c) 而不是 0(l)。
    # 下面用两趟扫描：第一趟统计频次，第二趟按字符串顺序找第一个频次为 1 的字符。
    count = {}
    for ch in s:
        count[ch] = count.get(ch, 0) + 1
    # 第一遍扫描完毕后，再进行一遍扫描，找到第一个只出现一次的字符
    for i, ch in enumerate(s):
        if count[ch] == 1:
            return i
    # 如果找不到只出现一次的字符，则返回 -1
    return -1

    # 原 Go 版本的解法（有 bug，保留作参考）：
    # for v in "abcdefghijkllmnopqrstuvwxyz":
    #     index = s.find(v)
    #     last_index = s.rfind(v)
    #     if index != -1 and last_index != -1 and index == last_index:
    #         return index
    # return -1
