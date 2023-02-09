
"""
NC97 字符串出现次数的TopK问题
https://www.nowcoder.com/practice/fd711bdfa0e840b381d7e1b82183b3ee?tpId=117&tqId=37809&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=


描述
给定一个字符串数组，再给定整数 k ，请返回出现次数前k名的字符串和对应的次数。
返回的答案应该按字符串出现频率由高到低排序。如果不同的字符串有相同出现频率，按字典序排序。
对于两个字符串，大小关系取决于两个字符串从左到右第一个不同字符的 ASCII 值的大小关系。
比如"ah1x"小于"ahb"，"231"<”32“
字符仅包含数字和字母

数据范围：字符串数满足 0 \le n \le 1000000≤n≤100000，每个字符串长度 0 \le n \le 100≤n≤10，0 \le k \le 25000≤k≤2500
要求：空间复杂度 O(n)O(n)，时间复杂度O(n \log k)O(nlogk)

示例1
输入：
["a","b","c","b"],2
复制
返回值：
[["b","2"],["a","1"]]
复制
说明：
"b"出现了2次，记["b","2"]，"a"与"c"各出现1次，但是a字典序在c前面，记["a","1"]，最后返回[["b","2"],["a","1"]]

示例2
输入：
["123","123","231","32"],2
复制
返回值：
[["123","2"],["231","1"]]
复制
说明：
 "123"出现了2次，记["123","2"]，"231"与"32"各出现1次，但是"231"字典序在"32"前面，记["231","1"]，最后返回[["123","2"],["231","1"]]
示例3
输入：
["abcd","abcd","abcd","pwb2","abcd","pwb2","p12"],3
复制
返回值：
[["abcd","4"],["pwb2","2"],["p12","1"]]
复制
备注：
1 \leq N \leq 10^51≤N≤10
5


"""

# from typing import List
#
#
# class Solution:
#     def topKstrings(self, strings: List[str], k: int) -> List[List[str]]:
        # write code here
#         d = {}
#         for item in strings:
#             if item in d:
#                 d[item] += 1
#             else:
#                 d[item] = 1
#         v_d = {}
#         for key, val in d.items():
#             if val in v_d:
#                 v_d[val].append(key)
#             else:
#                 v_d[val] = [key]
#         keys = sorted(v_d.keys(), reverse=True)
#         res = []
#         count = 0
#         for v in keys:
#             for key in sorted(v_d[v]):
#                 res.append([key, v])
#                 count += 1
#             if count == k:
#                 return res
#         return res[:k]

from typing import List


class Solution:
    def topKstrings(self, strings: List[str], k: int) -> List[List[str]]:
        # write code here
        dict = {}
        for i in range(len(strings)):
            if strings[i] not in dict:
                dict[strings[i]] = 1
            else:
                dict[strings[i]] += 1
        return (sorted(dict.items(), key=lambda x: (-x[1], x[0]))[:k])


if __name__ == '__main__':
    result = Solution().topKstrings(["a", "b", "c", "b"], 2)
    print(result)

