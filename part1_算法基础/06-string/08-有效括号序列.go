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
该题和 heap&stack&queue 中的第十题一样的， 并且解法更好
*/
func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, c := range s {
		switch c {
		// 以下case 只压栈
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		default:
			// s 是个“]” 等特殊情况
			if len(stack) == 0 {
				return false
			}
			// 当前字符跟col中最后一个做比较，相等则弹栈， 否则false
			//if rune(c) == stack[len(stack)-1] {
			//	stack = stack[:len(stack)-1]
			//} else {
			//	return false
			//}
			lastV := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c != lastV {
				return false
			}
		}
	}

	// 全部弹栈完成，如果为空，说明是符合要求的
	return len(stack) == 0
}

func main() {
	valid := isValid("()[]{}")
	fmt.Println(valid)
}
