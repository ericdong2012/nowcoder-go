package main

/*
BM63 跳台阶
https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4?tpId=295&tqId=23261&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个 n 级的台阶总共有多少种跳法（先后次序不同算不同的结果）。
要求：时间复杂度：O(n) ，空间复杂度： O(1)

示例1
输入：
2
返回值：
2

说明：
青蛙要跳上两级台阶有两种跳法，分别是：先跳一级，再跳一级或者直接跳两级。因此答案为2

示例2
输入：
7
返回值：
21
*/

func jumpFloor(number int) int {
	// write code here
	step1, step2 := 0, 1
	for i := 0; i < number; i++ {
		//c := step1 + step2
		//step1 = step2
		//step2 = c
		step1, step2 = step2, step1 + step2
	}
	return step2
}