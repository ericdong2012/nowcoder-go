package main

import (
	"fmt"
)

/*
NC84 完全二叉树结点数
https://www.nowcoder.com/practice/512688d2ecf54414826f52df4e4b5693?tpId=117&tqId=37786&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5058&title=

描述
给定一棵完全二叉树的头节点head，返回这棵树的节点个数。

完全二叉树指：设二叉树的深度为h，则 [1,h-1] 层的节点数都满足 2^{i-1}

示例1
输入：
{1,2,3}
返回值：
3

示例2
输入：
{}
返回值：
0

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func nodeNum(root *TreeNode) int {
	// write code here
	if root == nil {
		return 0
	}

	leftNodes := nodeNum(root.Left)
	rightNodes := nodeNum(root.Right)

	return leftNodes + rightNodes + 1

	//stack := []*TreeNode{root}
	//flag := false
	//count := 1
	//for stack != nil {
	//	var newStack []*TreeNode
	//	for _, v := range stack {
	//		if (flag && (v.Left != nil || v.Right != nil)) || (v.Left == nil && v.Right != nil) {
	//			count++
	//			flag = false
	//		}
	//		if v.Left != nil {
	//			newStack = append(newStack, v.Left)
	//			count++
	//		}
	//
	//		if v.Right != nil {
	//			newStack = append(newStack, v.Right)
	//			count++
	//		}
	//		// 产生了叶子节点
	//		if v.Left == nil || v.Right == nil {
	//			flag = true
	//		}
	//	}
	//	stack = newStack
	//}
	//
	//return count
}

func main() {
	//node := TreeNode{
	//	Val: 3,
	//	Left: &TreeNode{
	//		Val:   1,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 7,
	//		Left: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   8,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}

	node := TreeNode{}

	result := nodeNum(&node)
	fmt.Printf("%+v\n", result)
}
