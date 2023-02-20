package main

import (
	"fmt"
)

/*
NC73 数组中出现次数超过一半的数字
https://www.nowcoder.com/practice/e8a1b01a2df14cb2b228b30ee6a92163?tpId=117&tqId=37770&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=


描述
给一个长度为 n 的数组，数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。
例如输入一个长度为9的数组[1,2,3,2,2,2,5,4,2]。由于数字2在数组中出现了5次，超过数组长度的一半，因此输出2。

要求：空间复杂度：O(1)，时间复杂度 O(n)
输入描述：
保证数组输入非空，且保证有解
示例1
输入：
[1,2,3,2,2,2,5,4,2]
返回值：
2

示例2
输入：
[3,3,3,3,2,2,2]
返回值：
3

示例3
输入：
[1]
返回值：
1

*/

func MoreThanHalfNum_Solution(numbers []int) int {
	// write code here
	result := map[int]int{}
	for _, v := range numbers {
		result[v] += 1
		// 不能相等，因为存在一种可能，数组中只有两种数据
		if result[v] > len(numbers)/2 {
			return v
		}
	}
	return 0

}

func main() {
	solution := MoreThanHalfNum_Solution([]int{1, 2, 3, 2, 2, 2, 5, 4, 2})
	fmt.Println(solution)
}
