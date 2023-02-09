package main

import "fmt"

/*
NC200 移动 0
https://www.nowcoder.com/practice/102586387caa4afcbad6f96affce9780?tpId=117&tqId=39389&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5054&title=


描述
给定一个数组，请你实现将所有 0 移动到数组末尾并且不改变其他数字的相对顺序。

示例1
输入：
[1,2,0,3]
返回值：
[1,2,3,0]

示例2
输入：
[1,2,3]
返回值：
[1,2,3]

示例3
输入：
[0,0]
返回值：
[0,0]

*/

func moveZeroes(nums []int) []int {
	// write code here
	// 实现0往后滚动
	x := 0
	for i, v := range nums {
		if v == 0 {
			x++
		} else if x != 0 {
			nums[i-x] = nums[i]
			nums[i] = 0
		}
	}
	return nums

}

func main() {
	zeroes := moveZeroes([]int{1, 2, 0, 0, 3})
	fmt.Println(zeroes)
}
