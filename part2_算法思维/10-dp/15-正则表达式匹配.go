package main

/*

正则表达式匹配
https://www.nowcoder.com/practice/28970c15befb4ff3a264189087b99ad4?tpId=196&tqId=37114&rp=1&ru=/exam/company&qru=/exam/company&sourceUrl=%2Fexam%2Fcompany&difficulty=undefined&judgeStatus=undefined&tags=&title=%E6%AD%A3

part2_算法思维\10-dp\16-regular-expression-matching.py
part3_分类\07-dp\08-正则表达式匹配(too hard).go

描述
请实现一个函数用来匹配包括'.'和'*'的正则表达式
1.模式中的字符'.'表示任意一个字符
2.模式中的字符'*'表示它前面的字符可以出现任意次（包含0次）
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但是与"aa.a"和"ab*a"均不匹配

数据范围:
1.str 只包含从 a-z 的小写字母
2.pattern 只包含从 a-z 的小写字母以及字符 . 和 *，无连续的 '*'

示例1
输入：
"aaa", "a*a"
返回值：
true
说明：
中间的*可以出现任意次的a，所以可以出现1次a，能匹配上

示例2
输入：
"aad", "c*a*d"
返回值：
true
说明：
因为这里 c 为 0 个，a被重复一次， * 表示零个或多个a。 因此可以匹配字符串 "aad"

示例3
输入：
"a", ".*"
返回值：
true
说明：
".*" 表示可匹配零个或多个（'*'）任意字符（'.'）

示例4
输入：
"aaab","a*a*a*c"
返回值：
false

*/

// 动态规划
// 设 s 的长度为 m ， p 的长度为 n ；将 s 的第 i 个字符记为 si，p 的第 j 个字符记为 pj，将 s 的前 i 个字符组成的子字符串记为 s[:i] , 同理将 p 的前 j 个字符组成的子字符串记为 p[:j] 。
// 因此，本题可转化为求 s[:m] 是否能和 p[:n] 匹配。

/*
	0 c * a * d
  0 1 0 1 0 1 0
  a 0 0 0 1 1 0
  a 0 0 0 0 1 0
  d 0 0 0 0 0 1

*/

func match(s string, p string) bool {
	// write code here
	m, n := len(s), len(p)
	// dp[i][j] 代表字符串 s 的前 i 个字符和 p 的前 j 个字符能否匹配。
	// 由于 dp[0][0] 代表的是空字符的状态， 因此 dp[i][j] 对应的添加字符是 s[i - 1] 和 p[j - 1] 。
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true //代表两个空字符串能够匹配。
	// 初始化首行
	// 首行 s 为空字符串，因此当 p 的偶数位为 * 时才能够匹配（即让 p 的奇数位出现 0 次，保持 p 是空字符串）。因此，循环遍历字符串 p ，步长为 2（即只看偶数位）。
	for j := 2; j <= n; j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				// 通配符'*'匹配了0次， return dp(i, j + 2)
				// 通配符'*'匹配了多次，return dp(i + 1, j)
				// 通配符'*' 和 '.'匹配了多次， return dp(i + 1, j)
				if dp[i][j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && s[i-1] == p[j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && p[j-2] == '.' {
					dp[i][j] = true
				}
			} else {
				// 匹配 .  或者 其他字符相等   dp[i][j] = dp[i-1][j-1]
				if dp[i-1][j-1] && s[i-1] == p[j-1] {
					dp[i][j] = true
				} else if dp[i-1][j-1] && p[j-1] == '.' {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[m][n]
}


/*
func match(str string, pattern string) bool {
	m := len(str)
	n := len(pattern)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	same := func(i, j int) bool {
		if i < 0 || j < 0 {
			return false
		}
		if pattern[j] == '.' {
			return true
		}
		if pattern[j] == '*' {
			return false
		}
		return str[i] == pattern[j]
	}

	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if pattern[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if same(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if same(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
*/
