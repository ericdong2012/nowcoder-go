package main

/*
BM91 反转字符串
https://www.nowcoder.com/practice/c3a6afee325e472386a1c4eb1ef987f3?tpId=295&tqId=1024337&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
写出一个程序，接受一个字符串，然后输出该字符串反转后的字符串。（字符串长度不超过1000）

数据范围： 0 \le n \le 10000≤n≤1000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
示例1
输入：
"abcd"
复制
返回值：
"dcba"
复制
示例2
输入：
""
复制
返回值：
""

*/

func solve(str string) string {
	// write code here
	//strs := []byte(str)
	//left, right:= 0, len(str) -1
	//for left <= right{
	//	strs[left], strs[right] = strs[right], strs[left]
	//}
	//return string(strs)
	res := []byte{}
	for i := len(str) - 1; i > 0; i-- {
		res = append(res, str[i])
	}
	return string(res)
}
