package main

import "fmt"

/*
NC241 计算器(二)
https://www.nowcoder.com/practice/a9c170bfaf7349e3acb475d786ab1c7d?tpId=117&tqId=39586&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给定一个字符串形式的表达式 s ，请你实现一个计算器并返回结果，除法向下取整。  字符串中包含 + , - , * , / , 保证表达式合法。

示例1
输入：
"1*10"
返回值：
10

示例2
输入：
"8*9-19"
返回值：
53

示例3
输入：
"100000*100*0"
返回值：
0

示例4
输入：
"100000*100/9"
返回值：
1111111

*/

func calculate2(s string) int {
	stack := make([]int, 0)
	// 上一次的符号 + - * /
	sign := '+'
	// 中间获取到的数字
	num := 0
	for i := 0; i < len(s); {
		// 获取具体的数字，  + - * /  ascii 都小于 ‘0’， 此时推出循环，走下面逻辑
		for j := i; j < len(s) && s[j] >= '0'; j, i = j+1, i+1 {
			bi := int(s[j] - '0')
			num = bi + num*10
		}
		switch sign {
		case '+':
			stack = append(stack, num)
		case '-':
			stack = append(stack, -1*num)
		case '*':
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, temp*num)
		case '/':
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, temp/num)
		}
		num = 0
		if i < len(s) {
			// s[i]  byte uint8
			// sign  rune int32   unicode
			sign = rune(s[i])
		}
		i += 1
	}

	result := 0
	for i := 0; i < len(stack); i++ {
		result += stack[i]
	}
	return result
}

func main() {
	i := calculate2("8*10-19")
	fmt.Println(i)
}
