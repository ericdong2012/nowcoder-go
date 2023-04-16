package main

import "fmt"

/*
NC49 最长的括号子串
https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad?tpId=117&tqId=37745&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=

描述
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。

例1: 对于字符串 "(()" 来说，最长的格式正确的子串是 "()" ，长度为 2 .
例2：对于字符串 ")()())" , 来说, 最长的格式正确的子串是 "()()" ，长度为 4 .

要求时间复杂度 O(n), 空间复杂度 O(n).

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

/*
栈或者动态规划

栈
step 1：可以使用栈来记录左括号下标。
step 2：遍历字符串，左括号入栈，每次遇到右括号则弹出左括号的下标。
step 3：然后长度则更新为当前下标与栈顶下标的距离。
step 4：遇到不符合的括号，可能会使栈为空，因此需要使用start记录上一次结束的位置，这样用当前下标减去start即可获取长度，即得到子串。
step 5：循环中最后维护子串长度最大值。

*/

func longestValidParentheses(s string) int {
	// 记录索引
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			// 移除一个
			stack = stack[:len(stack)-1]
			// 有元素
			if len(stack) != 0 {
				index := stack[len(stack)-1]
				result = max(result, i-index)
			} else {
				// 没有元素, 记录当前位置
				result = max(result, i)
			}
		} else {
			//  '('， 记录此时的索引
			stack = append(stack, i)
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	//parentheses := longestValidParentheses(")(()")
	parentheses := longestValidParentheses("()()")
	fmt.Println(parentheses)
}
