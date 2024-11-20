package _3_hash

import (
	"fmt"
)

/*
https://www.nowcoder.com/practice/c0d32c1ce5744472a01b2351a2c2767f?tpId=117&tqId=39348&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj&difficulty=undefined&judgeStatus=undefined&tags=&title=%E5%8D%95%E8%AF%8D%E6%8B%86%E5%88%86

描述
给定一个字符串 s 和一个字符串数组 dic, 在字符串的任意位置拆分任意次后得到的字符串集合 是否 是给定字符串数组的子集。（即拆分后的字符串集合中的所有字符串都在 dic 数组中），

示例1
输入：
"nowcoder", ["now","coder","no","wcoder"]
返回值：
true

示例2
输入：
"nowcoder", ["now","wcoder"]
返回值：
false

示例3
输入：
"nowcoder", ["nowcoder"]
返回值：
true

示例4
输入：
"nownowcoder", ["now", "coder"]
返回值：
true

说明：
你可以重复使用 dic 数组中的字符串

*/


func wordDiv1(s string, dic []string) bool {
	// 对标的 dic 建立hash表， 并初始化
	m := make(map[string]bool)
	for _, v := range dic {
		m[v] = true
	}
	stack := make([]bool, len(s)+1)
	stack[0] = true
	// 滑动窗口
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			// 调试看看
			// s[j:i] 这个窗口中的字符串在m中，也就是dic中，
			// 并且stack[j] 存在, 当i走到3， stack[0]==true  , m["now"] ==> stack[2] = true
			// 并且stack[j] 存在, 当i走到10， stack[3]==true  , m["coder"] ==> stack[10] = true
			if stack[j] && m[s[j:i]] {
				stack[i] = true
				// 优化时间
				break
			}
		}
	}

	return stack[len(s)]
}

func main() {
	//div := wordDiv("nownowcoder", []string{"now", "coder"})
	div := wordDiv1("nownowcoder", []string{"n", "now", "owcoder", "coder"})
	fmt.Println(div)
}
