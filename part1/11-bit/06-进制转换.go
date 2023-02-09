package main

import "fmt"

/*
NC112 进制转换
https://www.nowcoder.com/practice/2cc32b88fff94d7e8fd458b8c7b25ec1?tpId=117&tqId=37836&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=


描述
给定一个十进制数 M ，以及需要转换的进制数 N 。将十进制数 M 转化为 N 进制数。
当 N 大于 10 以后， 应在结果中使用大写字母表示大于 10 的一位，如 'A' 表示此位为 10 ， 'B' 表示此位为 11 。
若 M 为负数，应在结果中保留负号。


示例1
输入：
7,2
返回值：
"111"

示例2
输入：
10,16
返回值：
"A"

备注：
M是32位整数，2<=N<=16.

*/

func solve(M int, N int) string {
	// write code here
	sign := false
	if M < 0 {
		sign = true
		M = -M
	}
	var ret string
	for M > 0 {
		mod := M % N
		var a byte
		if mod <= 9 {
			a = '0' + byte(mod)
		} else {
			a = 'A' + byte(mod-10)
		}

		ret = string(a) + ret
		M = M / N
	}
	// 符号更改回去
	if sign {
		ret = "-" + ret
	}
	return ret
}

func main() {
	//s := solve(20, 16)
	s := solve(10, 16)
	fmt.Println(s)
}
