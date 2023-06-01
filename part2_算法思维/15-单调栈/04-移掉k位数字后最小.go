package _5_单调栈

import (
	"fmt"
	"strings"
)

/*
https://www.nowcoder.com/practice/0fe685c8272d40f1b9785fedd2499c1c?tpId=196&tqId=39412&ru=/exam/oj


给定一个以字符串表示的数字 num 和一个数字 k ，从 num 中移除 k 位数字，使得剩下的数字最小。如果可以删除全部数字，则结果为 0。
1.num仅有数字组成
2.num是合法的数字，不含前导0
3.删除之后的num，请去掉前导0（不算在移除次数中）

保证 num 中仅包含 0~9 的十进制数

示例1
输入：
"1432219",3
返回值：
"1219"
说明：
移除 4 3 2 后剩下 1219

示例2
输入：
"10",1
返回值：
"0"

示例3
输入：
"100999",3
返回值：
"9"

*/

func removeKnums(num string, k int) string {
	// write code here
	if len(num) <= k {
		return "0"
	}

	stack := make([]byte, 0, len(num))
	j := 0
	// 100999 => 9
	for i := 0; i < len(num); i++ {
		for len(stack) >= 1 && num[i] < stack[len(stack)-1] && j < k {
			stack = stack[:len(stack)-1]
			j++
		}
		stack = append(stack, num[i])
	}
	if j < k {
		stack = stack[:len(stack)-(k-j)]
	}

	tmp := string(stack)
	if len(tmp) == 0 {
		return "0"
	}
	tmp = strings.TrimLeft(tmp, "0")
	if tmp == "" {
		return "0"
	}
	return tmp
}

func main() {
	result := removeKnums("100999", 3)
	fmt.Println(result)
}
