package main

import "fmt"

/*
NC231 只出现一次的数字
https://www.nowcoder.com/practice/c04bd25f0396471b90dfc30d96b9109b?tpId=117&tqId=39456&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5074&title=

描述
给定一个整数数组，数组中有一个数出现了一次，其他数出现了两次，请找出只出现了一次的数。
数据范围：数组中元素个数满足 1 \le n \le 100000 \1≤n≤100000  ，数组中的元素大小满足 |val| \le 10^9 \∣val∣≤10
9

示例1
输入：
[1]
返回值：
1

示例2
输入：
[1,2,2]
返回值：
1


*/

func singleNumber1(nums []int) int {
	// write code here
	temp := nums[0]
	for i := 1; i < len(nums); i++ {
		temp = temp ^ nums[i]
	}
	return temp
}

func main() {
	number1 := singleNumber1([]int{1, 2, 2})
	fmt.Println(number1)
}
