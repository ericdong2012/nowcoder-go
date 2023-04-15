package main

import "fmt"

/*
NC190 字符串的全部子序列
https://www.nowcoder.com/practice/92e6247998294f2c933906fdedbc6e6a?tpId=117&tqId=39361&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=


描述
给定一个字符串s，长度为n，求s的所有子序列
1.子序列: 指一个字符串删掉部分字符（也可以不删）形成的字符串，可以是不连续的，比如"abcde"的子序列可以有"ace", "ad"等等
2.将所有的子序列的结果返回为一个字符串数组
3.字符串里面可能有重复字符，但是返回的子序列不能有重复的子序列，比如"aab"的子序列只有"","a","aa","aab","ab","b"，不能存在2个相同的"ab"
4.返回字符串数组里面的顺序可以不唯一

要求:时间复杂度为O(2^n)

示例1
输入：
"ab"
返回值：
["","a","ab","b"]
说明：
返回["","b","a","ab"]也是可以的，视为正确，顺序不唯一

示例3
输入：
"aab"
返回值：
["","a","aa","aab","ab","b"]

示例2
输入：
"dbcq"
返回值：
["","b","bc","bcq","bq","c","cq","d","db","dbc","dbcq","dbq","dc","dcq","dq","q"]

说明：
返回的字符串数组里面不能存在"ab", "ab"这样2个相同的子序列
*/

// 该题和11-08集合的所有子集 类似， 但是条件限制没有那么多
func generatePermutation(s string) []string {
	filter := make(map[string]bool)
	// 添加了一个“”
	var ans = []string{""}
	for _, v := range s {
		var tmp []string
		// 能把ans中前面有的数据跟当前结合起来
		for _, a := range ans {
			if !filter[a+string(v)] {
				tmp = append(tmp, a+string(v))
			}
			filter[a+string(v)] = true
		}
		ans = append(ans, tmp...)
	}
	return ans
}

func main() {
	permutation := generatePermutation("ab")
	fmt.Println(permutation)
}
