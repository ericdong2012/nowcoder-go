package main

import "fmt"

/*
NC138 矩阵最长递增路径
https://www.nowcoder.com/practice/7a71a88cdf294ce6bdf54c899be967a2?tpId=117&tqId=37850&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=



描述
给定一个 n 行 m 列矩阵 matrix ，矩阵内所有数均为非负整数。 你需要在矩阵中找到一条最长路径，使这条路径上的元素是递增的。并输出这条最长路径的长度。
这个路径必须满足以下条件：

1. 对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外。
2. 你不能走重复的单元格。即每个格子最多只能走一次。

数据范围：1 \le n,m \le 10001≤n,m≤1000，0 \le matrix[i][j] \le 10000≤matrix[i][j]≤1000
进阶：空间复杂度 O(nm)O(nm) ，时间复杂度 O(nm)O(nm)

例如：当输入为[[1,2,3],[4,5,6],[7,8,9]]时，对应的输出为5，
其中的一条最长递增路径如下图所示：

示例1
输入：
[[1,2,3],[4,5,6],[7,8,9]]
复制
返回值：
5
复制
说明：
1->2->3->6->9即可。当然这种递增路径不是唯一的。
示例2
输入：
[[1,2],[4,3]]
复制
返回值：
4
复制
说明：
 1->2->3->4
备注：
矩阵的长和宽均不大于1000，矩阵内每个数不大于1000
*/

var temp [][]int = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}

func solve2(matrix [][]int) int {
	// write code here
	/*
			矩阵内是非负数，求最长的递增路径的长度
			移动方向可以是上下左右，不能超出边界，这将是递归的判定条件
			同一条路径不能有重复的单元格，需要有记忆

			终止条件： 进入路径最后一个点后，四个方向要么是矩阵边界，要么没有递增的位置，路径不能再增长，返回上一级。
			返回值： 每次返回的就是本级之后的子问题中查找到的路径长度加上本级的长度。
			本级任务： 每次进入一级子问题，先初始化后续路径长度为0，然后遍历四个方向（可以用数组表示，下标对数组元素的加减表示去往四个方向），
					进入符合不是边界且在递增的邻近位置作为子问题，查找子问题中的递增路径长度。因为有四个方向，所以最多有四种递增路径情况，因此要维护当级子问题的最大值。

			step 1：使用一个dp数组记录iii，jjj处的单元格拥有的最长递增路径，这样在递归过程中如果访问到就不需要重复访问。
		    step 2：遍历矩阵每个位置，都可以作为起点，并维护一个最大的路径长度的值。
		    step 3：对于每个起点，使用dfs查找最长的递增路径：只要下一个位置比当前的位置数字大，就可以深入，同时累加路径长度。

	*/
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	res := 0
	var n int = len(matrix)
	var m int = len(matrix[0])

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			res = max3(res, dfs4(matrix, dp, i, j))
		}
	}

	return res
}

func dfs4(matix [][]int, dp [][]int, i, j int) int {
	// 访问过
	if dp[i][j] != 0 {
		return dp[i][j]
	}
	dp[i][j] += 1
	for k := 0; k < 4; k++ {
		nexti := i + temp[k][0]
		nextj := j + temp[k][1]

		if nextj >= 0 && nextj < len(matix[0]) && nexti >= 0 && nexti < len(matix) && matix[nexti][nextj] > matix[i][j] {
			dp[i][j] = max3(dp[i][j], dfs4(matix, dp, nexti, nextj)+1)
		}
	}
	return dp[i][j]
}

func max3(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}

}

func main() {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, 2)
	}

	//for i := 0; i < 2; i++ {
	//	dp[0][i] = 0
	//}

	fmt.Println(dp)
}
