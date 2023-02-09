package main

import (
	"fmt"
	"strconv"
)

/*
NC216 逆波兰表达式求值
https://www.nowcoder.com/practice/885c1db3e39040cbae5cdf59fb0e9382?tpId=117&tqId=39420&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给定一个逆波兰表达式，求表达式的值。

数据范围：表达式长度满足 1 \le n \le 10^4 \1≤n≤10
4
   ，表达式中仅包含数字和 + ，- , * , / ，其中数字的大小满足 |val| \le 200 \∣val∣≤200 。
示例1
输入：
["2","1","+","4","*"]
复制
返回值：
12
复制
示例2
输入：
["2","0","+"]
复制
返回值：
2
*/

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a+b)
		case "-":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a-b)
		case "*":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a*b)
		case "/":
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, a/b)
		default:
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}
	re := stack[len(stack)-1]
	return re
}

func main() {
	rpn := evalRPN([]string{"2", "1", "+", "4", "*"})
	fmt.Println(rpn)
}
