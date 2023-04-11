# -*- coding: utf-8 -*-
import collections
from typing import List

"""
387. 字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0

s = "loveleetcode"
返回 2
 

注意事项：您可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

"""


class Solution:
    def firstUniqChar(self, s: str) -> int:
        """
        这道题最优的解法就是线性复杂度了，为了保证每个元素是唯一的，至少得把每个字符都遍历一遍。
        算法的思路就是遍历一遍字符串，然后把字符串中每个字符出现的次数保存在一个散列表中。这个过程的时间复杂度为 O(N)，其中 N为字符串的长度
        :param s:
        :return:
        """
        # count = collections.Counter(s)

        # for idx, ch in enumerate(s):
        #     if count[ch] == 1 :
        #         return  idx
        # else:
        #     return  -1

        # or

        # for i, j in collections.Counter(s).items():
        #     if j == 1: return s.index(i)
        # return -1



        # method2:  字典
        # dicts = {}
        # for i in s:
        #     dicts[i] = dicts.get(i, 0) + 1
        # for i in range(len(s)):
        #     if dicts[s[i]] == 1:
        #         return i
        # return -1



        # metho3:
        """
        1.证明字母只出现了一次
        如果一个字符串中的字符在字符串中从左边搜索和从右边搜索得到的index一样，那就证明只有一个了
        2.循环每次是从第一个开始的，保证了执行的顺序
        
        作者：markzhao
        链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string/solution/pythonde-zui-you-xie-fa-by-markzhao/
        来源：力扣（LeetCode）
        著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
        """
        # for c in s:
        #     if s.find(c) == s.rfind(c):
        #         return s.find(c)
        # return -1


        """
        作者：imckl
        链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string/solution/python-4chong-fang-fa-yi-ge-12msde-gai-jin-jie-fa-/
        来源：力扣（LeetCode）
        著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
        """

        min_index = len(s) - 1
        for ch in "abcdefghijklmnopqrstuvwxyz":
            i = s.find(ch)
            if i != -1 and i == s.rfind(ch):
                min_index = min([min_index, i])
        return min_index if min_index < len(s) else  -1


s = Solution()
print(s.firstUniqChar("loveleetcode"))
