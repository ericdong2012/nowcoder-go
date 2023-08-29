package main

import (
	"fmt"
)

/*

NC6 二叉树中的最大路径和
https://www.nowcoder.com/practice/da785ea0f64b442488c125b441a4ba4a?tpId=117&tqId=37716&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=583&title=

描述
二叉树里面的路径被定义为:从该树的任意节点出发，经过父=>子 或者  子=>父的连接，达到任意节点的序列。
注意:
1. 同一个节点在一条二叉树路径里中最多出现一次
2. 一条路径至少包含一个节点，且不一定经过根节点

给定一个二叉树的根节点root，请你计算它的最大路径和
给出以下的二叉树，
	2
1       3
最优路径是:2=>1=>3，或者3=>1=>2，最大路径和= 2+1+3=6

示例1
输入：
{1,2,3}
 	1
2      3
返回值：
6

示例2
输入：
{-20,8,20,#,#,15,6}
   -20
8       20
	 15    6
返回值：
41
说明：
其中一条最大路径为:15=>20=>6，路径和为15+20+6=41

示例3
输入：
{-2,#,-3}
    -2
		-3
返回值：
-2

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/*

public int backtrace(TreeNode root){
    if(root == null) return 0;
    // 后序遍历，最后计算和，避免根节点值重复计算
    // 计算左、右子结点最大路径和
    int left = backtrace(root.left);
    int right = backtrace(root.right);
    int cur = root.val;
    if(left > 0){
        // 获取左子节点值
        cur += left;
    }
    if(right > 0){
        // 获取右子节点值
        cur += right;
    }
    // 统计出现的最大路径和
    max = Math.max(max, cur);
    // 返回当前最大路径和，不能直接返回cur，因为cur包含了两条路径，而返回只能一条路径
    return Math.max(root.val, Math.max(left, right) + root.val);
}

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
	// 后序遍历
	// 分别求出左右子树的最大深度
	left := max(dfs(root.Left), 0)
	right := max(dfs(root.Right), 0)
	// 统计出现的最大路径和 =  左子树的最大深度加右子树的最大深度加本节点
	res = max(res, root.Val+left+right)
	// 返回当前节点的最大路径
	num := max(left, right) + root.Val
	return num
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}



func main19() {
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
