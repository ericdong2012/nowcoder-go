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
	// 方法1：
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

	// 方法2：(方法1 更容易懂)
	// 栈
	if len(s) <= 1 {
		return 0
	}
	// 用来存索引，需要有个元素，不然会报错 比如:  )(
	stack := []int{-1}
	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			// todo: 先弹栈 （为什么先弹栈, 过程还是比较绕的， ‘()’ 通过断点调试，思考该过程）
			stack = stack[:len(stack)-1]
			// 有元素
			if len(stack) != 0 {
				result = max(result, i-stack[len(stack)-1])
			} else {
				// 没有元素, 只有之前添加的 -1
				stack = append(stack, i)
			}
		} else {
			stack = append(stack, i)
		}
	}

	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	parentheses := longestValidParentheses("())")
	fmt.Println(parentheses)
}
