package main

import "fmt"

/*
BM77 最长的括号子串
https://www.nowcoder.com/practice/45fd68024a4c4e97a8d6c45fc61dc6ad?tpId=295&tqId=715&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给出一个长度为 n 的，仅包含字符 '(' 和 ')' 的字符串，计算最长的格式正确的括号子串的长度。

例1: 对于字符串 "(()" 来说，最长的格式正确的子串是 "()" ，长度为 2 .
例2：对于字符串 ")()())" , 来说, 最长的格式正确的子串是 "()()" ，长度为 4 .

要求时间复杂度 O(n),空间复杂度 O(n)


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
栈


动态规划
class Solution:
    def longestValidParentheses(self , s: str) -> int:
        res = 0
        #长度为0的串，返回0
        if len(s) == 0:
            return res
        #dp[i]表示以下标为i的字符为结束点的最长合法括号长度
        dp = [0 for i in range(len(s))]
        #第一位不管是左括号还是右括号都是0，因此不用管
        for i in range(1, len(s)):
            #取到左括号记为0，有右括号才合法
            if s[i] == ')':
                #如果该右括号前一位就是左括号
                if s[i - 1] == '(':
                    #计数+2
                    if i >= 2:
                        dp[i] = dp[i - 2] + 2
                    else:
                        dp[i] = 2
                #找到这一段连续合法括号序列前第一个左括号做匹配
                elif i - dp[i - 1] > 0 and s[i - dp[i - 1] - 1] == '(':
                    if i - dp[i - 1] > 1:
                        dp[i] = dp[i - dp[i - 1] - 2] + dp[i - 1] + 2
                    else:
                        dp[i] = dp[i - 1] + 2
            #维护最大值
            res = max(res, dp[i])
        return res

*/

func longestValidParentheses(s string) int {
	res := 0
	// 存左括号索引
	stack := []int{}
	start := -1
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			//如果右括号时栈为空，不合法，设置为结束位置
			if len(stack) == 0 {
				start = i
			} else {
				//弹出左括号
				stack = stack[:len(stack)-1]
				//栈中还有左括号，说明右括号不够，减去栈顶位置就是长度
				if len(stack) != 0 {
					index := stack[len(stack)-1]
					res = max3(res, i-index)
				} else {
					//栈中没有括号，说明左右括号行号，减去上一次结束的位置就是长度
					/* ()()   /   )()  */
					res = max3(res, i-start)
				}
			}
		} else {
			stack = append(stack, i)
		}
	}
	return res
}

func max3(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func main() {
	//parentheses := longestValidParentheses("(()()")
	parentheses := longestValidParentheses("()")
	fmt.Println(parentheses)
}
