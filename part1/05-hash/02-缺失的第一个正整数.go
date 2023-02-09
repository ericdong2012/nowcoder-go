package main

import (
	"fmt"
)

/*
NC30 缺失的第一个正整数
https://www.nowcoder.com/practice/50ec6a5b0e4e45348544348278cdcee5?tpId=117&tqId=37800&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

描述
给定一个无重复元素的整数数组nums，请你找出其中没有出现的最小的正整数

进阶： 空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)

数据范围:
-231<=nums[i]<=231-1
0<=len(nums)<=5*105
示例1
输入：
[1,0,2]
复制
返回值：
3
复制
示例2
输入：
[-2,3,4,1,5]
复制
返回值：
2
复制
示例3
输入：
[4,5,6,8,9]
复制
返回值：
1
*/

func minNumberDisappeared(nums []int) int {

	// write code here
	// 构造数据
	temp := make([]int, 5e5)
	// 更改temp 中的value
	temp[0] = 1
	for _, v := range nums {
		if v > 0 {
			temp[v] += 1
		}
	}
	// 遍历第一个0
	for k, v := range temp {
		if v == 0  {
			return k
		}
	}
	return 1
}

//func findIndex(result map[int]int, num int) int {
//	for k, _ := range result {
//		if k == num {
//			return k
//		}
//	}
//	return -1
//}

func main() {
	// 构造一个足够大的map
	nums := []int{1, 23, 3, 4}
	//temp := map[int]int{1: 0, 2: 0, 3: 0}
	//for _, v := range nums {
	//	if v >= 0 {
	//		index := findIndex(temp, v)
	//		if index > 0 {
	//			temp[index] = 1
	//		}
	//	}
	//}
	//
	//for _, v := range temp {
	//	if v == 0 {
	//		return v
	//	}
	//}

	disappeared := minNumberDisappeared(nums)
	fmt.Println(disappeared)

}
