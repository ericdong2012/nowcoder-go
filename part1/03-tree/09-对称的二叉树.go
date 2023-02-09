package main

import "fmt"

/*
BM31 对称的二叉树
https://www.nowcoder.com/practice/ff05d44dfdb04e1d83bdbdab320efbcb?tpId=295&tqId=23452&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一棵二叉树，判断其是否是自身的镜像（即：是否对称）

要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
备注：
你可以用递归和迭代两种方法解决这个问题

示例1
输入：
{1,2,2,3,4,4,3}
返回值：
true

示例2
输入：
{8,6,9,5,7,7,5}
返回值：
false

*/

// 层序遍历 + 左右交换后 是否相等

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func isSymmetrical(pRoot *TreeNode) bool {
	// write code here
	if pRoot == nil {
		return true
	}

	return helper9(pRoot.Right, pRoot.Left)
}

func helper9(right, left *TreeNode) bool {
	if right == nil && left == nil {
		return true
	}
	if right == nil || left == nil {
		return false
	}

	if right.Val != left.Val {
		return false
	}

	return helper9(right.Left, left.Right) && helper9(right.Right, left.Left)
}

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

	result := isSymmetrical(&node)
	fmt.Printf("%+v\n", result)
}
