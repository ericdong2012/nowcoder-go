package main

import (
	"fmt"
)

/*
BM95 分糖果问题
https://www.nowcoder.com/practice/76039109dd0b47e994c08d8319faa352?tpId=295&tqId=1008104&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
一群孩子做游戏，现在请你根据游戏得分来发糖果，要求如下：
1. 每个孩子不管得分多少，起码分到一个糖果。
2. 任意两个相邻的孩子之间，得分较多的孩子必须拿多一些糖果。(若相同则无此限制)

给定一个数组 arr 代表得分数组，请返回最少需要多少糖果。

示例1
输入：
[1,1,2]
返回值：
4
说明：
最优分配方案为1,1,2

示例2
输入：
[1,1,1]
返回值：
3
说明：
最优分配方案是1,1,1
*/

/*
step 1：使用一个辅助数组记录每个位置的孩子分到的糖果，全部初始化为1.
step 2：从左到右遍历数组，如果右边元素比相邻左边元素大，意味着在递增，糖果数就是前一个加1，否则保持1不变。
step 3：从右到左遍历数组，如果左边元素比相邻右边元素大，意味着在原数组中是递减部分，如果左边在上一轮中分到的糖果数更小，则更新为右边的糖果数+1，否则保持不变。
step 4：将辅助数组中的元素累加求和。
*/
func candy(arr []int) int {
	// write code here
	res := make([]int, len(arr))
	for i := 0; i < len(res); i++ {
		res[i] = 1
	}
	// 排序
	//sort.Ints(arr)
	// 从左到右遍历, 右边在递增, 对复合条件的加1
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			res[i] = res[i-1] + 1
		}
	}
	// 有个误区，不能对整体做比较，只能局部比较，任意两个相邻的孩子之间，得分较多的孩子必须拿多一些糖果， 所以不能上来就排序
	//result := 0
	//for i := 0; i < len(res); i++ {
	//	result += res[i]
	//}
	result := res[len(res)-1]
	// 从右到左遍历,
	for i := len(arr) - 2; i >= 0; i-- {
		// 如果左边更大 并且 糖果数更小或相等
		if arr[i] > arr[i+1] && res[i] <= res[i+1] {
			res[i] = res[i+1] + 1
		}
		result += res[i]
	}

	return result
}

func main() {
	//i2 := candy([]int{1, 1, 2, 3})
	i2 := candy([]int{1, 4, 3, 2, 7, 9}) // 1,2,1,2,3
	fmt.Println(i2)
}
