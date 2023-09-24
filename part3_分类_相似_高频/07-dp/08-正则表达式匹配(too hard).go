package main

import "fmt"

/*
NC122 正则表达式匹配
https://www.nowcoder.com/practice/28970c15befb4ff3a264189087b99ad4?tpId=117&tqId=37780&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=

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

/*
	0 c * a * d
  0 1 0 1 0 1 0
  a 0 0 0 1 1 0
  a 0 0 0 0 1 0
  d 0 0 0 0 0 1

step 1：设dp[i][j]表示str前i个字符和pattern前j个字符是否匹配。（需要注意这里的i，j是长度，比对应的字符串下标要多1）
step 2：（初始条件） 首先，毋庸置疑，两个空串是直接匹配，因此dp[0][0]=true。然后我们假设str字符串为空，那么pattern要怎么才能匹配空串呢？答案是利用'*'字符出现0次的特性。遍历pattern字符串，如果遇到'*'意味着它前面的字符可以出现0次，要想匹配空串也只能出现0，那就相当于考虑再前一个字符是否能匹配，因此dp[0][i]=dp[0][i−2]。
step 3：（状态转移） 然后分别遍历str与pattern的每个长度，开始寻找状态转移。首先考虑字符不为'*'的简单情况，只要遍历到的两个字符相等，或是pattern串中为'.'即可匹配，因此最后一位匹配，即查看二者各自前一位是否能完成匹配，即dp[i][j]=dp[i−1][j−1] 。
然后考虑'*'出现的情况：pattern[j - 2] == '.' || pattern[j - 2] == str[i - 1]：即pattern前一位能够多匹配一位，可以用'*'让它多出现一次或是不出现，
因此有转移方程 dp[i][j] = dp[i−1][j] ∣∣ dp[i][j−2]

不满足上述条件，只能不匹配，让前一个字符出现0次，dp[i][j]=dp[i][j−2]

*/

// 画图看几个大的地方，几个关键的值，上下左右，对角线，i-2
func match(str string, pattern string) bool {
	m, n := len(str), len(pattern)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			/*
				0 a a d
			  0 1 0 0 0
			  c 0 0
			  * 1 0         (此处a 要和 * 相等，则a 和他的前一位相等)
			  a 0 1 0
			  * 1 1 1
			  d 0

			  dp[i][j] = dp[i-1][j-1]
			  dp[i][j] = dp[i-1][j] || dp[i][j-2]   纵 和 横

			*/
			if pattern[j-1] != '*' {
				// 相等或者是 pattern[j-1] == '.' 则转化为判断子问题  dp[i-1][j-1]
				if i > 0 && (str[i-1] == pattern[j-1] || pattern[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			} else {
				// 有 '*' 的情况下

				// * 当做0使用( 把前面字符消除掉 ) 横向，也适用于i==0
				if j > 1 {
					dp[i][j] = dp[i][j-2]
				}
				// * 当做1次或者多次用(前面的字符出现1次或者多次)
				// i>0 && j >1 后续用在i-1, j-2
				// dp[i-1][j]  前一位匹配
				// 首先判断 * 与 当前i 对应的字符进行匹配，然后保留 b* 这个整体，然后i 转移到i-1，再看i-i 与 b* 是否匹配
				if i > 0 && j > 1 && dp[i-1][j] && (str[i-1] == pattern[j-2] || pattern[j-2] == '.') {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}
	return dp[m][n]
}

func main08() {
	b := match("aad", "c*a*d")
	fmt.Println(b)
}
