"""
NC170 最长不含重复字符的子字符串
https://www.nowcoder.com/practice/48d2ff79b8564c40a50fa79f9d5fa9c7

描述
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。

示例1
输入："abcabcbb"
返回值：3
说明：因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例2
输入："bbbbb"
返回值：1

示例3
输入："pwwkew"
返回值：3
说明：因为无重复字符的最长子串是 "wke"，所以其长度为 3。
请注意，你的答案必须是子串的长度，"pwke" 是一个子序列，不是子串。
"""

# 连续
# 不重复
# 最长
#
# 前缀和 + hash + 双指针


def length_of_longest_substring(s):
    # 用左右分别来统计存在和不存在
    left, right = 0, 0
    maxlen = 0
    temp = {}
    while left < len(s):
        # 已经走到了最后，防止索引越界
        if right == len(s):
            return maxlen
        # arr[right] 在 temp 中不存在
        if temp.get(s[right], 0) == 0:
            temp[s[right]] = temp.get(s[right], 0) + 1
            right += 1
        else:
            # arr[right] 在 temp 中存在
            # 有重复的将上一个置零(从头部开始置 0)
            temp[s[left]] = 0
            left += 1
        maxlen = max(maxlen, right - left)

    return maxlen


if __name__ == "__main__":
    # substring = length_of_longest_substring("abcabcbb")
    substring = length_of_longest_substring("dvdfv")
    # substring = length_of_longest_substring("Kzz8")
    print(substring)
