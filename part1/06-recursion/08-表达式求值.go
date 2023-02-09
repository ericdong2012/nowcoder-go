package main

import "fmt"

/*
NC137 表达式求值
https://www.nowcoder.com/practice/c215ba61c8b1443b996351df929dc4d4?tpId=117&tqId=37849&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=

描述
请写一个整数计算器，支持加减乘三种运算和括号。

示例1
输入：
"1+2"
返回值：
3

示例2
输入：
"(2*(3-4))*5"
返回值：
-10

示例3
输入：
"3+2*3*4-1"
返回值：
26

*/

func solve(s string) int {
	// write code here
	stack := make([]int, 0)
	num := 0
	var sign byte = '+'
	result := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			bi := int(s[i] - '0')
			num = bi + 10*num
		}

		if s[i] == '(' {
			count := 1
			start, end := i+1, i+1
			for count != 0 {
				if s[end] == '(' {
					count += 1
				}
				if s[end] == ')' {
					count -= 1
				}
				end++
			}
			i = end - 1
			num = solve(s[start : end-1])
		}

		if s[i] < '0' || s[i] > '9' || i == len(s)-1 {
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -1*num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			}
			num = 0
			sign = s[i]
		}

	}

	for len(stack) != 0 {
		result += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return result
}

func main() {
	i := solve("(2*(3-4))*5")
	fmt.Println(i)
}
