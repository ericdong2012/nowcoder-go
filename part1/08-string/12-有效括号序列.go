package main

import "fmt"

/*

NC52 有效括号序列
https://www.nowcoder.com/practice/37548e94a270412c8b9fb85643c8ccc2?tpId=117&tqId=37749&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=

描述
给出一个仅包含字符'(', ')', '{', '}', '[' 和 ']' 的字符串， 判断给出的字符串是否是合法的括号序列
括号必须以正确的顺序关闭，"()"和"()[]{}"都是合法的括号序列， 但"(]"和"([)]"不合法

数据范围：字符串长度 0\le n \le 100000≤n≤100000
要求：空间复杂度 O(n)，时间复杂度 O(n)

示例1
输入：
"()[]{}"
返回值：
true

示例2
输入：
"[]"
返回值：
true

示例3
输入：
"([)]"
返回值：
false

*/

/*
可以是左括号入站，也可以是右括号入站

*/
func isValid(s string) bool {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

func main() {
	valid := isValid("()[]{}")
	fmt.Println(valid)
}
