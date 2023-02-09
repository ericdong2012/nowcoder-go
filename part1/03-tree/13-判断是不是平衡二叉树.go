package main

import "fmt"

/*
BM36 判断是不是平衡二叉树
https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222?tpId=295&tqId=23250&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。
平衡二叉树（Balanced Binary Tree），具有以下性质：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是平衡二叉树。
https://zhuanlan.zhihu.com/p/56066942

示例1
输入：
{1,2,3,4,5,6,7}
返回值：
true

示例2
输入：
{}
返回值：
true
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IsBalanced_Solution(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}
	return abs(getMaxDepth(pRoot.Left)-getMaxDepth(pRoot.Right)) <= 1
}

func getMaxDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	h1 := getMaxDepth(pRoot.Left)
	h2 := getMaxDepth(pRoot.Right)
	// 加上根节点
	return 1 + max1(h1, h2)
}

func main() {
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := IsBalanced_Solution(&node)
	fmt.Printf("%+v\n", result)
}
