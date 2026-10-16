"""394. 字符串解码（栈）
题目："3[a2[c]]" → "accaccacc"，k[encoded_string] 形式解码。
思路：栈存"上一层（前缀字符串, 待重复次数）"；
      遇 ']' 弹出并拼接：前缀 + 次数×当前段；数字要累计（可能多位）。
复杂度：时间 O(n·maxK·maxN)，空间 O(n)。
"""


class Solution:
    def decodeString(self, s: str) -> str:
        stack = []                # [(前缀, 次数)]
        cur, num = "", 0
        for ch in s:
            if ch.isdigit():
                num = num * 10 + int(ch)
            elif ch == "[":
                stack.append((cur, num))
                cur, num = "", 0
            elif ch == "]":
                prefix, k = stack.pop()
                cur = prefix + cur * k
            else:
                cur += ch
        return cur


if __name__ == "__main__":
    s = Solution()
    assert s.decodeString("3[a]2[bc]") == "aaabcbc"
    assert s.decodeString("3[a2[c]]") == "accaccacc"
    assert s.decodeString("2[abc]3[cd]ef") == "abcabccdcdcdef"
    assert s.decodeString("abc") == "abc"
    print("71 OK")
