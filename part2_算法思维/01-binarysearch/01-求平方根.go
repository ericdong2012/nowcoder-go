package main

import "fmt"

/*
NC32 求平方根
https://www.nowcoder.com/practice/09fbfb16140b40499951f55113f2166c?tpId=117&tqId=37734&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
实现函数 int sqrt(int x).
计算并返回 x 的平方根（向下取整）


https://connect.applecross.link/clash/930835/mLieDowiymML

示例1
输入：
2
返回值：
1

示例2
输入：
2143195649
返回值：
46294

*/


func sqrt(x int) int {
	// write code here
	if x == 0 || x == 1 {
		return x
	}
	mid := x / 2
	fmt.Println(mid)
	// 大于10
	for mid*mid > x {
		mid = mid / 2
	}
	// 小于10
	for mid*mid <= x {
		mid = mid + 1
	}
	return mid - 1
}

func main() {
	fmt.Printf("%+v", sqrt(10))
	//fmt.Printf("%+v", sqrt(3))
}
