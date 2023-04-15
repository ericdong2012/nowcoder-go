package _3_hash

/*
题⽬⼤意
给出⼀个字符串数组，要求对字符串数组⾥⾯有 Anagrams 关系的字符串进⾏分组。Anagrams 关系是
指两个字符串的字符完全相同，顺序不同，两者是由排列组合组成。

输入: ["eat", "tea", "tan", "ate", "nat", "bat"]

输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]

解题思路
这道题可以将每个字符串都排序，排序完成以后，相同 Anagrams 的字符串必然排序结果⼀样。把排序
以后的字符串当做 key 存⼊到 map 中。遍历数组以后，就能得到⼀个 map，key 是排序以后的字符
串，value 对应的是这个排序字符串以后的 Anagrams 字符串集合。最后再将这些 value 对应的字符串
数组输出即可。

https://www.nowcoder.com/practice/68bda87f91664623a4213abb00e41a5a?tpId=196&tqId=39745&ru=/exam/oj

*/

import (
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	record := map[string][]string{}
	for _, str := range strs {
		//sByte := []rune(str)
		//sort.Sort(sortRunes(sByte))
		//sstrs := record[string(sByte)]
		//sstrs = append(sstrs, str)
		//record[string(sByte)] = sstrs

		// 实际结果有点出入，就是字母的排序上
		sSplit := strings.Split(str, "")
		sort.Strings(sSplit)
		sStr := strings.Join(sSplit, "")
		if _, ok := record[sStr]; !ok {
			record[sStr] = []string{str}
		} else {
			record[sStr] = append(record[sStr], str)
		}
	}

	res := [][]string{}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}
