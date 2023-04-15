package main

import "fmt"

/*
NC185 岛屿的最大面积
https://www.nowcoder.com/practice/5568943d3a08403f932a5e54ec3ece71?tpId=117&tqId=39352&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=


描述
给定一个用 n*m 矩阵表示的群岛的地图，其中 1 表示岛屿， 0 表示海洋，每个岛屿的水平或竖直方向相邻的岛屿可以视为连在一起的岛屿，每一块岛屿视为面积为 1 ，请问面积最大的岛屿是多少。

例如：
当输入[[1,0],[0,1]]时，对应的地图为：
只有在水平或竖直方向相邻的岛屿可以连在一起，所以每个岛屿互相独立。最大面积是1
当输入[[1,1],[1,0]]时，对应的地图为：
三块岛屿可以连在一起，最大面积是3


示例1
输入：
[[1,0],[0,1]]
返回值：
1

示例2
输入：
[[1,1],[1,0]]
返回值：
3

示例3
输入：
[[0]]
返回值：
0

*/

func maxAreaIsland(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	var dfs func(int, int) int
	dfs = func(x, y int) int {
		// 退出条件
		if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == 0 {
			return 0
		}
		//保证之后再访问的时候，面积不增加
		grid[x][y] = 0
		// 将左 上  右  下 还有自身的结果 累加
		return dfs(x-1, y) + dfs(x, y-1) + dfs(x+1, y) + dfs(x, y+1) + 1
	}
	maxN := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				temp := dfs(i, j)
				maxN = max(maxN, temp)
			}
		}
	}
	return maxN
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	island := maxAreaIsland([][]int{[]int{1, 1}, []int{1, 0}})
	fmt.Println(island)
}
