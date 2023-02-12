package main

import "fmt"

/*
NC86 矩阵元素查找
https://www.nowcoder.com/practice/3afe6fabdb2c46ed98f06cfd9a20f2ce?tpId=117&tqId=37788&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
已知一个有序矩阵mat，同时给定矩阵的大小n和m以及需要查找的元素x，且矩阵的行和列都是从小到大有序的。设计查找算法返回所查找元素的二元数组，代表该元素的行号和列号(均从零开始)。保证元素互异。


示例1
输入：
[[1,2,3],[4,5,6]], 6
返回值：
[1,2]

示例2
输入：
[[1,2,3]],1,3,2
返回值：
[0,1]

*/


// 该题和array 中二维数组的查找类似
func findElement(mat [][]int, n int, m int, x int) []int {
	for i := 0; i < n; i++ {
		left, right := 0, m-1
		for left <= right {
			mid := (left + right) / 2
			if mat[i][mid] == x {
				return []int{i, mid}
			} else if mat[i][mid] < x {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return []int{-1, -1}

}

func main() {
	element := findElement([][]int{[]int{1, 2, 3}, []int{4, 5, 6}}, 2, 3, 6)
	fmt.Println(element)
}
