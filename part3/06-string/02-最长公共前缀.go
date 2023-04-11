package main

/*
NC55 最长公共前缀
https://www.nowcoder.com/practice/28eb3175488f4434a4a6207f6f484f47?tpId=117&tqId=37752&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=

描述
给你一个大小为 n 的字符串数组 strs ，其中包含n个字符串 , 编写一个函数来查找字符串数组中的最长公共前缀，返回这个公共前缀。

示例1
输入：
["abca","abc","abca","abc","abcc"]
返回值：
"abc"

示例2
输入：
["abc"]
返回值：
"abc"
*/

/*
(排序后)纵向扫描
横向扫描

class Solution:
    def longestCommonPrefix(self , strs: List[str]) -> str:
        n = len(strs)
        #空字符串数组
        if n == 0:
            return ""
        #遍历第一个字符串的长度
        for i in range(len(strs[0])):
            temp = strs[0][i]
            #遍历后续的字符串
            for j in range(1,n):
                #比较每个字符串该位置是否和第一个相同
                if i == len(strs[j]) or strs[j][i] != temp:
                    #不相同则结束
                    return strs[0][0:i]
        #后续字符串有整个字符串的前缀
        return strs[0]

*/

func longestCommonPrefix(strs []string) string {
	// write code here
	// ["abca","abc","abca","abc","abcc"]
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(string(strs[0])); i++ {
		temp := string(strs[0])[i]
		for j := 1; j < len(strs); j++ {
			// 如果i 和 strs[j] 长度相等(某个字符串走到头了) 或者  strs[j] 的第i位 和 strs[0]的第i位 不相等， 则strs[0] 要缩短
			if i == len(string(strs[j])) || string(strs[j])[i] != temp {
				return string(strs[0])[:i]
			}
		}
	}
	// 遍历完，则说明strs[0] 就是最长公共前缀
	return string(strs[0])
}
