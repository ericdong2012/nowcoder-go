package main

import "fmt"

/*
NC240 计算器(一)
https://www.nowcoder.com/practice/9b1fca7407954ba5a6f217e7c3537fed?tpId=117&tqId=39585&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=

描述
给定一个字符串形式的表达式 s ，请你实现一个计算器并返回结果。 字符串中包含 + , -  , ( , ) ，保证表达式合法。


示例1
输入：
"1+2"
返回值：
3

示例2
输入：
"1-1"
返回值：
0

示例3
输入：
"-0"
返回值：
0

*/

/*
通过符号把减法看作加上负数的加法来做：
(1) 遇到数字时计算数字；
(2) 遇到运算符时累加结果；
(3) 遇到左括号时把当前计算出来的结果和符号压栈，并清空当前的计算结果，转而计算括号内的表达式；
(4) 遇到右括号时弹栈计算。

注意在每次的计算操作时，需要把当前的数字归零，下次遇到数字字符时重新累加。
*/

func calculate(s string) int {
	stack := make([]int, 0)
	// 记录上一个符号
	sign := 1
	// 记录中间数字
	num := 0
	for i := 0; i < len(s); i++ {
		//for m := i; m < len(s) && s[m] >= '0'; m, i = m+1, i+1 {
		//	bi := int(s[i] - '0')
		//	num = bi + num*10
		//}
		switch s[i] {
		case '+':
			stack = append(stack, sign*num)
			sign = 1
			num = 0
		case '-':
			stack = append(stack, sign*num)
			sign = -1
			num = 0
		case '(':
			// 记录括号个数
			count := 0
			k := i
			for ; k < len(s); k++ {
				if s[k] == '(' {
					count += 1
				}
				if s[k] == ')' {
					count -= 1
				}
				if count == 0 {
					break
				}
			}
			num = calculate(s[i+1 : k])
			// i 直接从最后的k 开始往后执行
			i = k
		default:
			bi := int(s[i] - '0')
			num = bi + num*10
		}

	}
	stack = append(stack, sign*num)

	result := 0
	for i := 0; i < len(stack); i++ {
		result += stack[i]
	}
	return result

}

func main() {
	i := calculate("1+(22-1)+3")
	fmt.Println(i)
}
