package main

import "fmt"

/*
最长的括号子串
https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad?tpId=117&tqId=37745&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%25E6%259C%2580%25E9%2595%25BF%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=&title=%E6%9C%80%E9%95%BF

描述
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。

示例1
输入：
 "(()"
返回值：
2

示例2
输入：
"(())"
返回值：
4

*/

func longestValidParentheses(s string) int {
	// write code here
	//stack := make([]rune, 0)
	//count := 0
	//for i := 0; i < len(s); i++ {
	//	if s[i] == '(' {
	//		stack = append(stack, rune(s[i]))
	//	} else {
	//		if len(stack) != 0 {
	//			stack = stack[:len(stack)-1]
	//			count += 2
	//		}
	//	}
	//}
	//
	//return count

	//if len(s) == 0 {
	//	return 0
	//}
	//stack := make([]int, len(s)+1)
	//top := 0
	//stack[top] = -1
	//res := 0
	//for i := 0; i < len(s); i++ {
	//	if s[i] == '(' {
	//		top++
	//		stack[top] = i
	//	} else {
	//		top--
	//		if top == -1 {
	//			top++
	//			stack[top] = i
	//		} else {
	//			res = max(res, i-stack[top])
	//		}
	//	}
	//}
	//
	//return res

	// dp
	/*
		dp[i]表示以i结尾最长合法字符串。如果s[i]=='('时该字符串一定不合法；当s[i]==')'时，假设存在解，那么该右括号与其对应的左括号之间的字符串一定是合法的。因此对于i-1的位置，以i-1结尾的合法字符串的开头下标为i - dp[i - 1]，当其前一个位置s[i - 1 - dp[i - 1]] == '('时，可以与s[i]进行匹配，dp[i]更新为dp[i - 1] + 2。
		此时还需要注意，如果在与当前右括号匹配的左括号的前一个位置(i - 1 - dp[i - 1]) - 1，以该处为结尾的最长合法字符串不为0，也需要加到结果上。例如()()
	*/
	dp := make([]int, len(s))
	maxValue := 0
	// 起始点是1
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			preIndex := i - 1 - dp[i-1]
			if preIndex >= 0 && s[preIndex] == '(' {
				dp[i] = dp[i-1] + 2
				if preIndex-1 >= 0 {
					dp[i] += dp[preIndex-1]
				}
			}
		}
		maxValue = max(maxValue, dp[i])
	}

	return maxValue
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	parentheses := longestValidParentheses(
		"()()")
	fmt.Println(parentheses)
}
