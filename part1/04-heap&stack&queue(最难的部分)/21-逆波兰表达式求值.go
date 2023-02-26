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
表达式中仅包含数字和 + ，- , * , / ，


示例1
输入：
["2","1","+","4","*"]
返回值：
12

示例2
输入：
["2","0","+"]
返回值：
2
*/

func evalRPN(tokens []string) int {
	// 保存数字及中间结果
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			// 不抽取出来是因为前面的数据取值会有问题
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
		// 默认添加数据
		default:
			num, _ := strconv.Atoi(tokens[i])
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	rpn := evalRPN([]string{"2", "1", "+", "4", "*"})
	fmt.Println(rpn)
}
