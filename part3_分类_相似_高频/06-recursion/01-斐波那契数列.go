package main

import "fmt"

/*
NC65 斐波那契数列
https://www.nowcoder.com/practice/c6c7742f5ba7442aada113136ddea0c3?tpId=117&tqId=37760&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


描述
大家都知道斐波那契数列，现在要求输入一个正整数 n ，请你输出斐波那契数列的第 n 项。
斐波那契数列是一个满足 fib(x)=\left\{ \begin{array}{rcl} 1 & {x=1,2}\\ fib(x-1)+fib(x-2) &{x>2}\\ \end{array} \right.fib(x)={

要求：空间复杂度 O(1)，时间复杂度 O(n) ，本题也有时间复杂度 O(logn) 的解法



示例1
输入：
4
返回值：
3

说明：
根据斐波那契数列的定义可知，fib(1)=1,fib(2)=1,fib(3)=fib(3-1)+fib(3-2)=2,fib(4)=fib(4-1)+fib(4-2)=3，所以答案为3。

示例2
输入：
1
返回值：
1

示例3
输入：
2
返回值：
1

*/


// 一定查看下官方的题解，非常经典
func Fibonacci(n int) int {
	//if n == 1 || n == 2 {
	//	return 1
	//}
	//return Fibonacci(n-1)  + Fibonacci(n-2)

	//if n == 1 || n == 2 {
	//	return 1
	//}
	//pre, curr, temp := 1, 1, 0
	//for idx := 3; idx <= n; idx++ {
	//	// 保存下一次结果
	//	temp = pre + curr
	//	pre = curr
	//	curr = temp
	//}
	//return temp

	a, b := 0, 1
	for i := 1; i < n; i++ {
		temp := a + b
		a = b
		b = temp
	}
	return b
}

func main01() {
	fibonacci := Fibonacci(5)
	fmt.Println(fibonacci)
}
