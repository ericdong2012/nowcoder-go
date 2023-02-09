package main

import "fmt"

/*
NC68 跳台阶
https://www.nowcoder.com/practice/8c82a5b80378478f9484d87d1c5f12a4?tpId=117&tqId=37764&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


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

/*

本质是递归问题
针对结果的个数，可以转化为斐波拉切问题

假设f[i]表示在第i个台阶上可能的方法数, 逆向思维。如果我从第n个台阶进行下台阶，下一步有2中可能，一种走到第n-1个台阶，一种是走到第n-2个台阶。
所以f[n] = f[n-1] + f[n-2]. 那么初始条件，f[0] = f[1] = 1。

*/

func jumpFloor(number int) int {
	// write code here
	//if number < 3 {
	//	return number
	//}
	a, b := 0, 1
	for i := 0; i < number; i++ {
		c := a + b
		a = b
		b = c
	}
	return b
}

func main() {
	floor := jumpFloor(1)
	fmt.Println(floor)
}
