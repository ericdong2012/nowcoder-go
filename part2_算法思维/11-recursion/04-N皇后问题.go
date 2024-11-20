package main

import "fmt"

/*
NC39 N皇后问题
https://www.nowcoder.com/practice/c76408782512486d91eea181107293b6?tpId=117&tqId=37811&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=591&title=

描述
N 皇后问题是指在 n * n 的棋盘上要摆 n 个皇后， 要任何两个皇后不同行，不同列也不在同一条斜线上，
要求：空间复杂度 O(1) ，时间复杂度 O(n!)

例如当输入4时，对应的返回值为2
对应的两种四皇后摆位如下图所示

示例1
输入：
1
返回值：
1

示例2
输入：
8
返回值：
92

*/

/*
递归 + 回溯

n个皇后，不同行不同列，那么肯定棋盘每行都会有一个皇后，每列都会有一个皇后。如果我们确定了第一个皇后的行号与列号，则相当于接下来在n−1n-1n−1行中查找n−1n-1n−1个皇后，这就是一个子问题，因此使用递归：

终止条件： 当最后一行都被选择了位置，说明n个皇后位置齐了，增加一种方案数返回。
返回值： 每一级要将选中的位置及方案数返回。
本级任务： 每一级其实就是在该行选择一列作为该行皇后的位置，遍历所有的列选择一个符合条件的位置加入数组，然后进入下一级。
具体做法：

step 1：对于第一行，皇后可能出现在该行的任意一列，我们用一个数组pos记录皇后出现的位置，下标为行号，元素值为列号。
step 2：如果皇后出现在第一列，那么第一行的皇后位置就确定了，接下来递归地在剩余的n−1n-1n−1行中找n−1n-1n−1个皇后的位置。
step 3：每个子问题检查是否符合条件，我们可以对比所有已经记录的行，对其记录的列号查看与当前行列号的关系：即是否同行、同列或是同一对角线。
*/

/*
# 标准解法
class Solution:
    #判断皇后是否符合条件
    def isValid(self, pos:List[int], row:int, col:int):
        #遍历所有已经记录的行
        for i in range(row):
            #不能同行同列同一斜线
            if row == i or col == pos[i] or abs(row - i) == abs(col - pos[i]):
                return False
        return True

    #递归查找皇后种类
    def recursion(self, n:int, row:int, pos:List[int], res:int):
        #完成全部行都选择了位置
        if row == n:
            res += 1
            return int(res)
        #遍历所有列
        for i in range(n):
            #检查该位置是否符合条件
            if self.isValid(pos, row, i):
                #加入位置
                pos[row] = i
                #递归继续查找
                res = self.recursion(n, row + 1, pos, res)
        return res

    def Nqueen(self, n:int) -> int:
        res = 0
        #下标为行号，元素为列号，记录皇后位置
        pos = [0] * n
        #递归
        result = self.recursion(n, 0, pos, res)
        return result

*/



// 位运算
//func Nqueen(n int) int {
//	limit := 1<<n - 1
//	return process(limit, 0, 0, 0)
//}
//func process(limit, col, right, left int) int {
//	if limit == col {
//		return 1
//	}
//	pos := (^(right | left | col)) & limit
//	mostone, res := 0, 0
//	for pos != 0 {
//		mostone = pos & (^pos + 1)
//		pos -= mostone
//		res += process(limit, col|mostone, (right|mostone)>>1, (left|mostone)<<1)
//	}
//	return res
//}

/*
递归，回溯

1.设置三个集合分别记录不能再被选中的的列col, 正斜线pos, 反斜线neg
2.经规律发现 行号i - 列号j 可确定唯一正斜线，行号i + 列号j 可确定唯一反斜线
3.符合要求的点记录当前点并递归下一个皇后, 最后一个皇后成功安置后将res+1, 然后需回溯回初始点将初始点删除, 将初始点的皇后放置其他位置进行判断
4.不符合要求的需要进行循环
*/

// 迭代
func Nqueen(n int) int {
	// write code here
	// 列
	column := make([]bool, n)
	// 副对角线
	dg := make([]bool, 2*n+1)
	// 主对角线
	udg := make([]bool, 2*n+1)
	cnt := 0

	dfs2(0, n, &cnt, column, dg, udg)
	return cnt
}

func dfs2(row, n int, cnt *int, column, dg, udg []bool) {
	if row == n {
		*cnt++
		return
	}
	for i := 0; i < n; i++ {
		// 主要搞清楚 i-row+n, row+i 意思(i代表的是y, row 代表的x )
		// 副对角线 y-x = b (A[0][0]、A[1][1]、…、A[n-1][n-1]、A[n][n]；) 遍历时搜索到的范围为这条线的下半部分，导致 -x + y < 0，出现负数。为了防止出现负数，将其进行规格化为n - x + y，保证对角线上元素的相对值相同即可。
		// 主对角线 y+x = b (A[0][n]、A[1][n-1]、…、A[n-1][1]、A[n][0]；)
		// https://blog.csdn.net/qq_41094332/article/details/116274425
		// 没有被占用
		if !column[i] && !dg[i-row+n] && !udg[row+i] {
			// 标记，代表这些地方不能放了
			column[i], dg[i-row+n], udg[row+i] = true, true, true
			dfs2(row+1, n, cnt, column, dg, udg)
			column[i], dg[i-row+n], udg[row+i] = false, false, false
		}
	}
}

func main05() {
	nqueen := Nqueen(4)
	fmt.Println(nqueen)
}
