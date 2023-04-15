package main

import "fmt"

/*
NC119 最小的K个数
https://www.nowcoder.com/practice/6a296eb82cf844ca8539b57c23e6e9bf?tpId=117&tqId=37765&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给定一个长度为 n 的可能有重复值的数组，找出其中不去重的最小的 k 个数。例如数组元素是 4,5,1,6,2,7,3,8 这8个数字，则最小的4个数字是1,2,3,4(任意顺序皆可)。
要求：空间复杂度 O(n) ，时间复杂度 O(nlogn)

示例1
输入：
[4,5,1,6,2,7,3,8],4
返回值：
[1,2,3,4]
说明：
返回最小的4个数即可，返回[1,3,2,4]也可以

示例2
输入：
[1],0
返回值：
[]

示例3
输入：
[0,1,2,1,2],3
返回值：
[0,1,1]

*/

func GetLeastNumbers_Solution(input []int, k int) []int {
	length := len(input)
	if k > length || length == 0 {
		return nil
	}
	for i := 0; i < k; i++ {
		for j := length - 1; j > i; j-- {
			// 后一个数比前一个小, 实现升序
			if input[j] < input[j-1] {
				input[j], input[j-1] = input[j-1], input[j]
			}
		}
	}
	//fmt.Println(input)
	return input[:k]
}

func main() {
	solution := GetLeastNumbers_Solution([]int{4, 5, 1, 6, 2, 7, 3, 8}, 4)
	fmt.Println(solution)
}
