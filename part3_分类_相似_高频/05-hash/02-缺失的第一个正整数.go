package main

import (
	"fmt"
)

/*
NC30 缺失的第一个正整数
https://www.nowcoder.com/practice/50ec6a5b0e4e45348544348278cdcee5?tpId=117&tqId=37800&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=585&title=

描述
给定一个无重复元素的整数数组nums，请你找出其中没有出现的最小的正整数
进阶： 空间复杂度 O(1)，时间复杂度 O(n)

示例1
输入：
[1,0,2]
返回值：
3

示例2
输入：
[-2,3,4,1,5]
返回值：
2

示例3
输入：
[4,5,6,8,9]
返回值：
1
*/


// 排序，看首位元素
// 如果大于等于2， 则最小的是1
// 否则, 从1开始遍历，判断是否在数组中，不在，则该数就是最小正整数

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


	/*
		也不错

	   sort.Ints(nums)
		if nums[0] >= 2 {
			return 1
		}
		temp := make(map[int]bool, nums[len(nums)-1])
		for i := 0; i < len(nums); i++ {
			temp[nums[i]] = true
		}

		// 从1 到nums 中的最大数
		for i := 1; i <= nums[len(nums)-1]; i++ {
			if _, ok := temp[i]; ok {
				continue
			} else {
				return i
			}
		}
		// 如果全部遍历完还没有找到，最是最大的加上1
		return nums[len(nums)-1] + 1
	*/
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
	nums := []int{1, 23, 3, 4}
	disappeared := minNumberDisappeared(nums)
	fmt.Println(disappeared)

}
