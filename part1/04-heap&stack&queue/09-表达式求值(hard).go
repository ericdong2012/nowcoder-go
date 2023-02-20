package main

import "fmt"

/*
NC137 表达式求值
https://www.nowcoder.com/practice/c215ba61c8b1443b996351df929dc4d4?tpId=117&tqId=37849&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=

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

/*
写一个支持+ - *三种符号的运算器，其中优先级+ - 是一级，*更高一级
支持括号运算

栈 + 递归

step 1：使用栈辅助处理优先级，默认符号为加号。
step 2：遍历字符串，遇到数字，则将连续的数字字符部分转化为int型数字。
step 3：遇到左括号，则将括号后的部分送入递归，处理子问题；遇到右括号代表已经到了这个子问题的结尾，结束继续遍历字符串，将子问题的加法部分相加为一个数字，返回。
step 4：当遇到符号的时候如果是+，得到的数字正常入栈，如果是-，则将其相反数入栈，如果是*，则将栈中内容弹出与后一个元素相乘再入栈。
step 5：最后将栈中剩余的所有元素，进行一次全部相加。

*/

func solve(s string) int {
	//// 把每次运算的结果放到该栈中
	//stack := make([]int, 0)
	//
	//// 字符串转换为字符切片
	//data := /*[]byte(s)*/ s
	//
	//// 表示字符串中每一个数字
	//number := 0
	//// 初识状态操作符
	//operation := byte('+')
	//
	//// 处理字符串
	//for i := 0; i < len(data); i++ {
	//	// 当前字符
	//	c := data[i]
	//
	//	// 0~9 字符是数字
	//	if c >= '0' && c <= '9' {
	//		number = 10*number + int(c-'0')
	//		//continue 不能写continue，因为如果是最后一个字符的话，需要特殊处理
	//	}
	//
	//	// '('
	//	// 遇到左括号特殊处理，还会进入到递归中
	//	if c == '(' {
	//		// 遇到的第一个左括号
	//		count := 1
	//		// 进入递归需要处理的开始于截止
	//		start, end := i+1, i+1
	//
	//		// 每一次遇到括号，就是一层递归
	//		// 每一次递归都是去掉最外层的一对括号，直到没有括号
	//		// count == 0 就是最有一个右括号
	//		for count != 0 {
	//			// 遇到左括号 ++
	//			if data[end] == '(' {
	//				count++
	//			}
	//			//遇到右括号 --
	//			if data[end] == ')' {
	//				count--
	//			}
	//
	//			end++
	//		}
	//
	//		// 已经处理好的字符位置
	//		i = end - 1
	//
	//		// 进入递归，并顺便传入需要处理的字符串
	//		// start, end-1 为需要处理的字符串的开始于截止位置
	//		number = solve(s[start : end])
	//	}
	//
	//	// c < '0' || c > '9' 表示 '+' / '-' / '*'
	//	// i == len(data)-1 表示已经是最后一个字符了，此时需要进行最后的一次运算
	//	if c == '+' || c == '-' || c == '*' || i == len(data)-1 {
	//		switch operation {
	//		case '+':
	//			stack = append(stack, number)
	//		case '-':
	//			stack = append(stack, -number)
	//		case '*':
	//			stack[len(stack)-1] = stack[len(stack)-1] * number
	//		}
	//
	//		// 重置数字
	//		number = 0
	//
	//		// 当前的操作符号，前面已经对括号进行了处理，这里不会有括号
	//		//  i == len(data)-1 时，operation 无意义了
	//		operation = c
	//	}
	//}
	//
	//// 累加栈中每一个数字
	//result := 0
	//for len(stack) != 0 {
	//	result += stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//}
	//
	//return result

	// write code here
	stack := make([]int, 0)
	data := []byte(s)
	number := 0
	var operation byte = '+'
	result := 0
	for i := 0; i < len(data); i++ {
		// 0~9
		if data[i] >= '0' && data[i] <= '9' {
			number = 10*number + int(data[i]-'0')
		}

		// '('
		if data[i] == '(' {
			count := 1
			start, end := i+1, i+1
			for count != 0 {
				if data[end] == '(' {
					count++
				}
				if data[end] == ')' {
					count--
				}
				end++
			}
			i = end - 1
			number = solve(s[start : end-1])
		}

		// '+'/'-'/'*'
		if data[i] < '0' || data[i] > '9' || i == len(data)-1 {
			switch operation {
			case '+':
				stack = append(stack, number)
			case '-':
				stack = append(stack, -number)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * number
			}
			number = 0
			operation = data[i]
		}
	}

	for len(stack) != 0 {
		result += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return result

}

func main() {
	i := solve("2*(3-4)+5")
	fmt.Println(i)
}
