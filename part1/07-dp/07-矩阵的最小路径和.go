package main

/*
BM68 矩阵的最小路径和
https://www.nowcoder.com/practice/7d21b6be4c6b429bb92d219341c4f8bb?tpId=295&tqId=1009012&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个 n * m 的矩阵 a，从左上角开始每次只能向右或者向下走，最后到达右下角的位置，路径上所有的数字累加起来就是路径和，输出所有的路径中最小的路径和。

例如：当输入[[1,3,5,9],[8,1,3,4],[5,0,6,1],[8,8,4,0]]时，对应的返回值为12，

示例1
输入：
[[1,3,5,9],[8,1,3,4],[5,0,6,1],[8,8,4,0]]
返回值：
12

示例2
输入：
[[1,2,3],[1,2,3]]
返回值：
7

*/

func minPathSum(matrix [][]int) int {
	// write code here
	m, n := len(matrix), len(matrix[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				matrix[i][j] += matrix[i][j-1]
				continue
			}
			if j == 0 {
				matrix[i][j] += matrix[i-1][j]
				continue
			}
			matrix[i][j] += min3(matrix[i-1][j], matrix[i][j-1])
		}
	}

	return matrix[m-1][n-1]
}

func min3(a, b int) int {
	if a < b {
		return a
	}

	return b
}
