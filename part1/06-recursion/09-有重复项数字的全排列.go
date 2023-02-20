package main

import (
	"fmt"
	"sort"
)

/*
NC42 有重复项数字的全排列
https://www.nowcoder.com/practice/a43a2b986ef34843ac4fdd9159b69863?tpId=117&tqId=37739&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


描述
给出一组可能包含重复项的数字，返回该组数字的所有排列。结果以字典序升序排列。

数据范围： 0 < n \le 80<n≤8 ，数组中的值满足 -1 \le val \le 5−1≤val≤5
要求：空间复杂度 O(n!)，时间复杂度 O(n!)

示例1
输入：
[1,1,2]
返回值：
[[1,1,2],[1,2,1],[2,1,1]]

示例2
输入：
[0,1]
返回值：
[[0,1],[1,0]]

*/

func permuteUnique(num []int) [][]int {
	// write code here

	if len(num) == 0 {
		return nil
	}
	sort.Ints(num)

	result := [][]int{}
	dfs1(num, 0, &result)

	return result

}

func dfs1(nums []int, index int, result *[][]int) {
	// 排列到最后一位，无需再处理，直接生成字符序列
	if index == len(nums)-1 {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*result = append(*result, temp)
		return
	}

	for i := index; i < len(nums); i++ {
		//// 剪枝, 下一次循环的，字符相同
		if i != index && nums[i] == nums[index] {
			continue
		}
		// 递归处理下一层
		nums[i], nums[index] = nums[index], nums[i]
		dfs1(nums, index+1, result)
	}
	for i := len(nums) - 1; i > index; i-- {
		nums[i], nums[index] = nums[index], nums[i]
	}
}

func main() {
	unique := permuteUnique([]int{1, 1, 2})
	fmt.Println(unique)
}
