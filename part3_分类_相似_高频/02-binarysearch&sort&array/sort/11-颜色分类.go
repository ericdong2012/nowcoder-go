package main

import "fmt"

/*
NC212 颜色分类
https://www.nowcoder.com/practice/52e04ddb7b5640a8869c2d3da2ad3344?tpId=117&tqId=39416&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590&title=

描述
给定一个包含红色，白色，蓝色，一同 n 个元素的数组，对其进行排序使得相同的颜色相邻并且按照 红色，白色，蓝色的顺序排序。
数组中 0 代表红色，1 代表白色，2 代表蓝色。 数组中只包含 0 1 2。


示例1
输入：
[0,2,1]
返回值：
[0,1,2]

示例2
输入：
[0,0,2,0]
返回值：
[0,0,0,2]
*/

func sortColor(colors []int) []int {
	// write code here
	stack0 := []int{}
	stack1 := []int{}
	stack2 := []int{}
	for i := 0; i < len(colors); i++ {
		if colors[i] == 0 {
			stack0 = append(stack0, colors[i])
		}
		if colors[i] == 1 {
			stack1 = append(stack1, colors[i])
		}
		if colors[i] == 2 {
			stack2 = append(stack2, colors[i])
		}
	}
	stack0 = append(stack0, stack1...)
	stack0 = append(stack0, stack2...)
	return stack0

	// nb 解法
	//less, more, curr := -1, len(colors), 0
	//
	//for curr < more {
	//	switch v := colors[curr]; v {
	//	case 2:
	//		more--
	//		colors[curr], colors[more] = colors[more], colors[curr]
	//	case 0:
	//		less++
	//		colors[curr], colors[less] = colors[less], colors[curr]
	//		fallthrough
	//	default:
	//		curr++
	//	}
	//}
	//
	//return colors
}

func main() {
	color := sortColor([]int{2, 1, 2})
	fmt.Println(color)
}
