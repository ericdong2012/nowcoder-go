package main

import "fmt"

/*
NC38 螺旋矩阵
https://www.nowcoder.com/practice/7edf70f2d29c4b599693dc3aaeea1d31?tpId=117&tqId=37738&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=590,5058,578&title=

描述
给定一个m x n大小的矩阵（m行，n列），按螺旋的顺序返回矩阵中的所有元素。
数据范围：0 \le n,m \le 100≤n,m≤10，矩阵中任意元素都满足 |val| \le 100∣val∣≤100
要求：空间复杂度 O(nm) ，时间复杂度 O(nm)

示例1
输入：
[
[1,2,3],
[4,5,6],
[7,8,9]
]
返回值：
[1,2,3,6,9,8,7,4,5]

示例2
输入：
[]
返回值：
[]

*/

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	if len(matrix[0]) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	left, right, up, down := 0, n-1, 0, m-1

	rets := make([]int, 0, m*n)
	for left <= right && up <= down {
		// left -> right
		for i := left; i <= right; i++ {
			rets = append(rets, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		// up -> down
		for i := up; i <= down; i++ {
			rets = append(rets, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		// right -> left
		for i := right; i >= left; i-- {
			rets = append(rets, matrix[down][i])
		}
		down--
		for i := down; i >= up; i-- {
			rets = append(rets, matrix[i][left])
		}
		left++
		//  最后那个元素在上述的某个地方会加入进去，同时通过break 结束循环
	}
	return rets
}

func main() {
	order := spiralOrder([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
	fmt.Println(order)
}
