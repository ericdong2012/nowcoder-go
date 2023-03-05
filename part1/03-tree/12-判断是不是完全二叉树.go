package main

import (
	"fmt"
)

/*
BM35 判断是不是完全二叉树
https://www.nowcoder.com/practice/8daa4dff9e36409abba2adbe413d6fae?tpId=295&tqId=2299105&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个二叉树，确定他是否是一个完全二叉树。

完全二叉树的定义：若二叉树的深度为h，除第h 层外，其它各层的结点数都达到最大个数，第 h 层所有的叶子结点都连续集中在最左边，这就是完全二叉树。（第 h 层可能包含 [1~2h] 个节点）

首先我们需要知道什么是完全二叉树：叶子节点只能出现在最下层和次下层，且最下层的叶子节点集中在树的左部。需要注意的是，满二叉树肯定是完全二叉树，而完全二叉树不一定是满二叉树。

示例1
输入：
{1,2,3,4,5,6}
返回值：
true

示例2
输入：
{1,2,3,4,5,6,7}
返回值：
true

示例3
输入：
{1,2,3,4,5,#,6}
返回值：
false

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/*
	层序遍历, 队列   从上到下遍历所有层，每层从左到右，只有次下层和最下层才有叶子节点，其他层出现叶子节点就意味着不是完全二叉树。
*/

func isCompleteTree(root *TreeNode) bool {
	// write code here
	stack := []*TreeNode{root}
	flag := false
	for stack != nil {
		var newStack []*TreeNode
		for _, v := range stack {
			if (flag && (v.Left != nil || v.Right != nil)) || (v.Left == nil && v.Right != nil) {
				return false
			}
			if v.Left != nil {
				newStack = append(newStack, v.Left)
			}
			if v.Right != nil {
				newStack = append(newStack, v.Right)
			}
			// 产生了叶子节点
			if v.Left == nil || v.Right == nil {
				flag = true
			}
		}
		stack = newStack
	}

	return true
}

func main12() {
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 7,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := isCompleteTree(&node)
	fmt.Printf("%+v\n", result)
}
