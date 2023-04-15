package main

/*
NC226 被围绕的区域
https://www.nowcoder.com/practice/3946670643fe4ec2aedcc2be45aed1a9?tpId=117&tqId=39451&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=

描述
给定一个 n*m 大小的的矩阵，矩阵中由 ‘X' 和 'O' 构成，找到所有被 'X' 围绕的区域，并将其用 'X' 填充。

例如：
[['X','X','X','X'],
['X','O','O','X'],
['X','O','X','X'],
['X','X','O','X']]
中间的三个 ‘O’ 被 'X'围绕，因此将其填充为 'X' ，但第四行的 'O' 下方没有被 'X' 围绕，因此不改变，结果为
[['X','X','X','X'],
['X','X','X','X'],
['X','X','X','X'],
['X','X','O','X']]

数据范围： 1 \le n , m \le 200 \1≤n,m≤200  ，矩阵中保证只含有 'X' 和 'O'
示例1
输入：
[[X,X,X,X],[X,O,O,X],[X,O,X,X],[X,X,O,X]]
复制
返回值：
[[X,X,X,X],[X,X,X,X],[X,X,X,X],[X,X,O,X]]
复制
示例2
输入：
[[O]]
复制
返回值：
[[O]]

*/

func check(i, j int, target [][]byte) bool {
	nums := 0
	// 上
	ii := i
	for ii >= 0 {
		if target[ii][j] == 'X' {
			nums++
			break
		}
		ii--
	}
	// 左
	jj := j
	for jj >= 0 {
		if target[i][jj] == 'X' {
			nums++
			break
		}
		jj--
	}
	// 下
	ii = i
	for ii <= len(target)-1 {
		if target[ii][j] == 'X' {
			nums++
			break
		}
		ii++
	}
	// 右
	jj = j
	for jj <= len(target[0])-1 {
		if target[i][jj] == 'X' {
			nums++
			break
		}
		jj++
	}

	if nums == 4 {
		return true
	}
	return false
}

func surroundedArea(board [][]byte) [][]byte {
	// write code here
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' && check(i, j, board) {
				board[i][j] = 'X'
			}
		}
	}
	return board

}
