package main

import "fmt"

/*
NC237 最大矩形
https://www.nowcoder.com/practice/5720efc1bdff4ca3a7dad37ca012cb60?tpId=117&tqId=39474&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


描述
给定一个仅包含 0 和 1 ，大小为 n*m 的二维二进制矩阵，找出仅包含 1 的最大矩形并返回面积。

数据范围：1 \le n,m \le 200 \1≤n,m≤200  保证输入的矩形中仅含有 0 和 1

例如输入[[1,0,1,0,0],[1,1,1,1,0],[1,1,1,1,1],[1,0,0,1,0]]时，对应的输出为8，所形成的最大矩形如下图所示：
示例1
输入：
[[1]]
复制
返回值：
1
复制
示例2
输入：
[[1,1],[0,1]]
复制
返回值：
2
复制
说明：
第一列的两个 1 组成一个矩形
示例3
输入：
[[0,0],[0,0]]
复制
返回值：
0
复制
示例4
输入：
[[1,0,1,0,0],[1,1,1,1,0],[1,1,1,1,1],[1,0,0,1,0]]
复制
返回值：
8

*/

/*
https://www.nowcoder.com/practice/5720efc1bdff4ca3a7dad37ca012cb60?tpId=117&tqId=39474&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=581,582,586&title=


def maximalRectangle(self , matrix: List[List[int]]) -> int:
	if not matrix:return 0
	m, n=len(matrix), len(matrix[0])
	# 记录当前位置上方连续“1”的个数
	pre=[0]*(n+1)
	res=0
	for i in range(m):
		for j in range(n):
			# 前缀和
			pre[j]=pre[j]+1 if matrix[i][j]==1 else 0

		# 单调栈
		stack=[-1]
		for k,num in enumerate(pre):
			while stack and pre[stack[-1]]>num:
				index=stack.pop()
				res=max(res,pre[index]*(k-stack[-1]-1))
			stack.append(k)

	return res

*/

/*

本质上是对矩阵中的每行依次执行直方图最大矩阵算法。

获得该行的直方图之后，就可以在heights上求解该行的直方图最大矩阵面积，例如题目中的第三行，得到的heights = [3, 2, 3, 2, 1]
*/

func maximalRectangle(matrix [][]int) (result int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	height := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := range height {
			height[j] = (height[j] + 1) * matrix[i][j] // update height
			minHeight := height[j]
			for k := j; k >= 0 && height[k] > 0; k-- { // loop valid width
				minHeight = min(minHeight, height[k])
				s := minHeight * (j - k + 1)
				if result < s {
					result = s
				}
			}
		}
	}
	return
}

func min(a, b int) (result int) {
	if a < b {
		result = a
	} else {
		result = b
	}
	return
}

func main() {
	rectangle := maximalRectangle([][]int{[]int{1, 0, 1, 0, 0}, []int{1, 1, 1, 1, 0}, []int{1, 1, 1, 1, 1}, []int{1, 0, 0, 1, 0}})
	fmt.Println(rectangle)
}
