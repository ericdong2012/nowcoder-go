package main

/*
NC109 岛屿数量
https://www.nowcoder.com/practice/0c9664d1554e466aa107d899418e814e?tpId=117&tqId=37833&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=


描述
给一个01矩阵，1代表是陆地，0代表海洋， 如果两个1相邻，那么这两个1属于同一个岛。我们只考虑上下左右为相邻。
岛屿: 相邻陆地可以组成一个岛屿（相邻:上下左右） 判断岛屿个数。
例如：
输入
[
[1,1,0,0,0],
[0,1,0,1,1],
[0,0,0,1,1],
[0,0,0,0,0],
[0,0,1,1,1]
]
对应的输出为3
(注：存储的01数据其实是字符'0','1')
示例1
输入：
[[1,1,0,0,0],[0,1,0,1,1],[0,0,0,1,1],[0,0,0,0,0],[0,0,1,1,1]]
复制
返回值：
3
复制
示例2
输入：
[[0]]
复制
返回值：
0
复制
示例3
输入：
[[1,1],[1,1]]
复制
返回值：
1
复制
备注：
01矩阵范围<=200*200

*/


/*

矩阵中多处聚集着1，要想统计1聚集的堆数而不重复统计，那我们可以考虑每次找到一堆相邻的1，就将其全部改成0，而将所有相邻的1改成0的步骤又可以使用深度优先搜索（dfs）：
当我们遇到矩阵的某个元素为1时，首先将其置为了0，然后查看与它相邻的上下左右四个方向，如果这四个方向任意相邻元素为1，则进入该元素，进入该元素之后我们发现又回到了刚刚的子问题，
又是把这一片相邻区域的1全部置为0，因此可以用递归实现。

*/

func solve(grid [][]byte) int {
	row := len(grid)
	col := len(grid[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		// 退出条件
		if x < 0 || x >= row || y < 0 || y >= col || grid[x][y] == '0' {
			return
		}
		//保证之后再访问的时候，面积不增加
		grid[x][y] = '0'
		// 将左 上  右  下
		dfs(x-1, y)
		dfs(x, y-1)
		dfs(x+1, y)
		dfs(x, y+1)
	}
	res := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				// 会不会有重复的？？？
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

func max1(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
