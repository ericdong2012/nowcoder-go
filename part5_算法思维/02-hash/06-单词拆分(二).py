"""
https://www.nowcoder.com/practice/bb40e0aae84a4491a35f93b322516cb5

描述
给定一个字符串 s 和一个字符串数组 dic，在字符串的任意位置拆分任意次后得到的字符串集合
是否是给定字符串数组的子集。（即拆分后的字符串集合中的所有字符串都在 dic 数组中）

示例1
输入："nowcoder", ["now","coder","no","wcoder"]
返回值：["no wcoder", "now coder"]

示例4
输入："nownowcoder", ["now", "coder"]
返回值：["now now coder"]

说明：你可以重复使用 dic 数组中的字符串
"""

# 思路：以 dic 为基础，做排列组合（全排列），然后看是否等于 s，等于存到数组中
# 很多边界没有考虑情况
# def word_div2(s, dic):
#     arr = []
#     for i in range(len(dic)):
#         for j in range(len(dic)):
#             if i == j and i != 0:
#                 continue
#             if len(dic) > 1:
#                 arr.append(dic[i] + " " + dic[j])
#             else:
#                 arr.append(dic[i])
#     result = []
#     for i in range(len(arr)):
#         strs = arr[i].split(" ")
#         if len(strs) > 1 and strs[0] + strs[1] == s:
#             result.append(arr[i])
#         elif len(strs) == 1 and strs[0] == s:
#             result.append(arr[i])
#     return result

"""
nowcoder1  nowcoder1    ==> nowcoder1
nowcoder11  nowcoder1 1  ==> nowcoder1 1 1
"""


# 解法1
def word_div(s, dic):
    """write code here"""
    m = {v: True for v in dic}
    ans = []
    path = []

    # 递归 + 回溯
    def dfs(idx):
        if idx == len(s):
            ans.append(" ".join(path))
            return
        for i in range(idx + 1, len(s) + 1):
            if m.get(s[idx:i], False):
                path.append(s[idx:i])
                dfs(i)
                path.pop()

    dfs(0)
    return ans


# 解法2
def word_break(s, word_dict):
    word_set = {word: True for word in word_dict}
    n = len(s)
    ans = []

    def dfs(start, end, sb):
        # 最终，做拼接
        if start >= n:
            ans.append(" ".join(sb))
            return
        # 弹栈
        if end >= n:
            return
        # 滑动窗口，当字符串符合要求，将下次递归的首位更新
        tmp = s[start:end + 1]
        if word_set.get(tmp, False):
            dfs(end + 1, end + 1, sb + [tmp])
        dfs(start, end + 1, sb)

    dfs(0, 0, [])
    return ans


"""
解法2扩展 - 字典树
class Trie:  # trie类
    def __init__(self):
        self.children = dict()  # 用dict存储类
        self.words = list()

class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        n = len(s)
        ans = []
        root = Trie()  # 初始化trie树
        for word in wordDict:  # 构建trie树
            cur = root
            for c in word:
                if c not in cur.children:
                    cur.children[c] = Trie()
                cur = cur.children[c]
            cur.words.append(word)

        def dfs(i, node, sb):
            if i == n:
                if node == root:  # 回到根
                    ans.append(" ".join(sb))
                return
            if s[i] not in node.children:  # 如果不匹配立刻结束
                return
            node = node.children[s[i]]
            if node.words:  # 断开
                dfs(i + 1, root, sb + [node.words[0]])
            dfs(i + 1, node, sb)  # 继续

        node = root
        dfs(0, node, [])
        return ans
"""

if __name__ == "__main__":
    div1 = word_div("nownowcoder", ["now", "coder"])
    div2 = word_break("nownowcoder", ["n", "now", "owcoder", "coder"])
    print(div1)
    print(div2)
