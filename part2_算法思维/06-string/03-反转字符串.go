package main

/*
NC103 反转字符串
https://www.nowcoder.com/practice/c3a6afee325e472386a1c4eb1ef987f3?tpId=117&tqId=37827&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=579&title=


描述
写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

数据范围： 0 \le n \le 10000≤n≤1000
要求：空间复杂度 O(n)，时间复杂度 O(n)

示例1
输入：
"abcd"
返回值：
"dcba"

示例2
输入：
""
返回值：
""
*/

func solve(str string) string {
	// write code here
	res := ""
	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return res
}
