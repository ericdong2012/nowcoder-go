package main

import (
	"fmt"
	"strings"
)

/*
NC182 单词拆分(二)
https://www.nowcoder.com/practice/bb40e0aae84a4491a35f93b322516cb5?tpId=117&tqId=39349&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

描述
给定一个字符串s和一个字符串数组dic ，在字符串s的任意位置添加任意多个空格后得到的字符串集合是给定字符串数组 dic 的子集（即拆分后的字符串集合中的所有字符串都在 dic 数组中）

示例1
输入：
"nowcoder",["now","coder","no","wcoder"]
返回值：
["no wcoder","now coder"]

示例2
输入：
"nowcoder",["now","wcoder"]
返回值：
[]

示例3
输入：
"nowcoder",["nowcoder"]
返回值：
["nowcoder"]

示例4
输入：
"nownowcoder",["now","coder"]
返回值：
["now now coder"]

*/
func wordDiv(s string, dic []string) []string {
	// 对标的dic 建立hash表， 并初始化
	m := make(map[string]bool)
	for _, v := range dic {
		m[v] = true
	}
	var (
		ans  []string // 最终结果
		path []string // 中间结果
		dfs  func(idx int)
	)
	// "nownowcoder",  ["n", "now", "owcoder", "coder"]
	dfs = func(idx int) {
		// 退出(终止条件):  当idx 到达s最后, 最后结果加上“ ” 就是最终结果, 退出
		if idx == len(s) {
			ans = append(ans, strings.Join(path, " "))
			return
		}
		// i 从idx 下一位开始，之所以 能等于 s长度  是因为 s[idx:i]  i 可以到 len(s)
		for i := idx + 1; i <= len(s); i++ {
			// s[ i - idx] 在hash表中存在时
			if m[s[idx:i]] {
				// path添加一位
				path = append(path, s[idx:i])
				dfs(i)
				// path去掉最后一位
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return ans
}

func main() {
	//div := wordDiv("nownowcoder", []string{"now", "coder"})
	div := wordDiv("nownowcoder", []string{"n", "now", "owcoder", "coder"})
	fmt.Println(div)
}
