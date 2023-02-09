package main

import (
	"fmt"
)

/*
NC6 二叉树中的最大路径和
https://www.nowcoder.com/practice/da785ea0f64b442488c125b441a4ba4a?tpId=117&tqId=37716&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=583&title=

描述
二叉树里面的路径被定义为:从该树的任意节点出发，经过父=>子或者子=>父的连接，达到任意节点的序列。
注意:
1.同一个节点在一条二叉树路径里中最多出现一次
2.一条路径至少包含一个节点，且不一定经过根节点

给定一个二叉树的根节点root，请你计算它的最大路径和

例如：
给出以下的二叉树，

最优路径是:2=>1=>3，或者3=>1=>2，最大路径和=2+1+3=6

数据范围：节点数满足 0 \le n \le 10^50≤n≤10
5
  ，节点上的值满足 |val| \le 1000∣val∣≤1000
要求：空间复杂度 O(1)O(1)，时间复杂度 O(n)O(n)
示例1
输入：
{1,2,3}
复制
返回值：
6
复制
示例2
输入：
{-20,8,20,#,#,15,6}
复制
返回值：
41
复制
说明：

其中一条最大路径为:15=>20=>6，路径和为15+20+6=41
示例3
输入：
{-2,#,-3}
复制
返回值：
-2


*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/*
	leftchild + rightchild + root-> val;
	leftchild + root->val;
	rightchild + root->val; 那么result的最大值就可能为上述三者中的最大值，依次执行后续遍历即可得出最优解。
*/

var res = -10000000

func maxPathSum(root *TreeNode) int {
	dfs(root)
	return res
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(dfs(root.Left), 0)
	right := max(dfs(root.Right), 0)
	sum := root.Val + left + right
	res = max(res, sum)
	return root.Val + max(left, right)
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

//var sum int = 0
//
//func maxPathSum(root *TreeNode) int {
//
//	// write code here
//	if root == nil {
//		return 0
//	}
//	sum += sum + root.Val
//	// 判断叶子节点 还有sum
//	if root.Left == nil && root.Right == nil {
//		mx := root.Val
//		return max(mx, sum)
//	}
//	return max( max(maxPathSum(root.Left),0), max(maxPathSum(root.Right),0))
//}
//
//func max(a int, b int) int {
//	if a >= b {
//		return a
//	}
//	return b
//}

func main() {
	// test
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := maxPathSum(&node)
	fmt.Printf("%+v\n", result)
}
