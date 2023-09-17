package _3_hash

import (
	"fmt"
)

/*
https://www.nowcoder.com/practice/bb40e0aae84a4491a35f93b322516cb5?tpId=196&tqId=39340&ru=/exam/oj

描述
给定一个字符串 s 和一个字符串数组 dic, 在字符串的任意位置拆分任意次后得到的字符串集合 是否 是给定字符串数组的子集。（即拆分后的字符串集合中的所有字符串都在 dic 数组中），

示例1
输入：
"nowcoder", ["now","coder","no","wcoder"]
返回值：
["no wcoder", "now coder"]

示例2
输入：
"nowcoder", ["now","wcoder"]
返回值：
[] / false

示例3
输入：
"nowcoder", ["nowcoder"]
返回值：
["nowcoder"]  / true

示例4
输入：
"nownowcoder", ["now", "coder"]
返回值：
["now now coder"]  / true

说明：
你可以重复使用 dic 数组中的字符串

*/

// 单词拆分一
func wordDiv(s string, dic []string) bool {
	//对标的 dic 建立hash表， 并初始化
	m := make(map[string]bool)
	for _, v := range dic {
		m[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]


	// 单词拆分 二
	//var (
	//	ans  []string // 最终结果
	//	path []string // 中间结果
	//	dfs  func(idx int)
	//)
	//// "nownowcoder",  ["n", "now", "owcoder", "coder"]
	//dfs = func(idx int) {
	//	// 当 idx 到达s最后, 最后结果加上“ ”,  退出（终止条件）
	//	if idx == len(s) {
	//		ans = append(ans, strings.Join(path, " "))
	//		return
	//	}
	//	// i 从 idx 下一位开始, i之所以能等于是因为 s[idx:i]  i可以到len(s)
	//	for i := idx + 1; i <= len(s); i++ {
	//		// i - idx 在dic中时
	//		if m[s[idx:i]] {
	//			// path添加一位
	//			path = append(path, s[idx:i])
	//			dfs(i)
	//			// path去掉最后一位
	//			path = path[:len(path)-1]
	//		}
	//	}
	//}
	//dfs(0)
	//// 此处和单词拆分2 不同
	//return len(ans) > 0
}

func main() {
	//div := wordDiv("nownowcoder", []string{"now", "coder"})
	div := wordDiv("nownowcoder", []string{"n", "now", "owcoder", "coder"})
	fmt.Println(div)
}
