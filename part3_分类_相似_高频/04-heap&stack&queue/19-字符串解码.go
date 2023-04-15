package main

import (
	"fmt"
	"strconv"
)

/*
NC199 字符串解码
https://www.nowcoder.com/practice/4e008fd863bb4681b54fb438bb859b92?tpId=117&tqId=39388&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=

描述
给一个加密过的字符串解码，返回解码后的字符串。
加密方法是：k[c] ，表示中括号中的 c 字符串重复 k 次，例如 3[a] 解码结果是 aaa ，保证输入字符串符合规则。不会出现类似 3a , 3[3] 这样的输入。

示例1
输入：
"3[a]"
返回值：
"aaa"

示例2
输入：
"abc"
返回值：
"abc"

示例3
输入：
"3[3[b]]"
返回值：
"bbbbbbbbb"

*/

// 栈
/*
class Solution:
    def decodeString(self , s: str) -> str:
        # write code here
        stack, digits, digit, cur = [], [], 0, ""
        for c in s:
            if "0" <= c <= "9":
                digit += 10 * digit + int(c)\


            elif c == "[":
                digits.append(digit)
                digit = 0
                stack.append(cur)
                cur = ""
            elif c == "]":
                stack[-1] += cur * digits.pop()
                cur = stack.pop()
            else:
                cur += c
        return stack[-1] + cur if stack else cur
*/

func decodeString(s string) string {
	// write code here
	digits := []int{}
	digit := 0
	cur := ""
	stack := []string{}
	// 3[3[b]]
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			temp, _ := strconv.Atoi(string(s[i]))
			digit += temp
		} else if s[i] == '[' {
			// 把之前的数字添加进，然后清空
			digits = append(digits, digit)
			digit = 0
			// 把之前的字符添加进，然后清空
			stack = append(stack, cur)
			cur = ""
		} else if s[i] == ']' {
			// 如果digits 不为空，则去除digits 的最后一位， 重复添加到stack最后一位
			for j := 0; j < digits[len(digits)-1]; j++ {
				stack[len(stack)-1] += cur
			}
			// 拿到stack 最后一位
			cur = stack[len(stack)-1]
			// 移除到stack 和 digits最后一位
			stack = stack[:len(stack)-1]
			digits = digits[:len(digits)-1]
		} else {
			// 具体的字符
			cur += string(s[i])
		}
	}
	// stack还有剩余
	if len(stack) > 0 {
		// todo 如何确定只需要最后一位呢
		return stack[len(stack)-1] + cur
	} else {
		return cur
	}
}

func main() {
	//s := decodeString("3[a]")
	//s := decodeString("abc")
	s := decodeString("3[3[b]]")
	//s := decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")
	fmt.Println(s)
}
