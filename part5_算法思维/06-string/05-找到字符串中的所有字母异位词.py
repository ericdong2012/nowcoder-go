"""
https://www.nowcoder.com/practice/9ff491c910e5427fab6ae67745929085

找到字符串中的所有字母异位词
给定一个字符串 s 和一个非空字符串 p，找出 s 中的所有 p 的异位字符串的子串，返回这些子串的起始索引。

解题思路
这道题是一道考"滑动窗口"的题目。和第 3 题，第 76 题，第 567 题类似的。解法也是用 freq[256] 记录每个字符的出现的频次次数。
滑动窗口左边界往右滑动的时候，划过去的元素释放次数(即次数 ++)，滑动窗口右边界往右滑动的时候，划过去的元素消耗次数(即次数 --)。
右边界和左边界相差 len(p) 的时候，需要判断每个元素是否都用过一遍了。具体做法是每经过一个符合规范的元素，count 就 --，
count 初始值是 len(p)，当每个元素都符合规范的时候，右边界和左边界相差 len(p) 的时候，count 也会等于 0。
当区间内有不符合规范的元素(freq < 0 或者是不存在的元素)，那么当区间达到 len(p) 的时候，count 无法减少到 0，
区间右移动的时候，左边界又会开始 count ++，只有当左边界移出了这些不合规范的元素以后，
才可能出现 count = 0 的情况，即找到 Anagrams 的情况。

"abab", "ab"
[0, 1, 2]
"""

# 解法3 最好
# 解法1
# def find_anagrams(s, p):
#     p_count = {}
#     for i in range(len(p)):
#         p_count[p[i]] = p_count.get(p[i], 0) + 1
#
#     result = []
#     s_count = {}
#     for i in range(len(s)):
#         # 将端点加入窗口
#         s_count[s[i]] = s_count.get(s[i], 0) + 1
#         # 如果当前窗口长度大于模式串长度，将左端点对应元素减 1；如果左端点对应的元素个数为 0，则删除
#         # 比如 abc, ab
#         if i >= len(p):
#             s_count[s[i - len(p)]] -= 1
#             if s_count[s[i - len(p)]] == 0:
#                 del s_count[s[i - len(p)]]
#         # 判断窗口内的字符与模式串是否相同
#         if s_count == p_count:
#             result.append(i - len(p) + 1)
#     return result


"""
解法2：
    size = len(p)
    stack = []
    sp = "".join(sorted(p))
    for i in range(len(s) - size + 1):
        ss = "".join(sorted(s[i:i + size]))
        if ss == sp:
            stack.append(i)
    return stack
"""


# 解法3
# 没有用排序，用的是 hash
def find_word(s, p):
    length = len(p)
    target = [0] * 26
    for v in p:
        target[ord(v) - ord('a')] += 1
    res = []
    for i in range(len(s) - length + 1):
        sub_str = s[i:i + length]
        tmp_arr = [0] * 26
        for v in sub_str:
            tmp_arr[ord(v) - ord('a')] += 1
        if tmp_arr == target:
            res.append(i)
    return res
