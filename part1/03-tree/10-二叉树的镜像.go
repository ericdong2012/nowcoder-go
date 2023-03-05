package main

import "fmt"

/*
BM33 二叉树的镜像
https://www.nowcoder.com/practice/a9d0ecbacef9410ca97463e4a5c83be7?tpId=295&tqId=1374963&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
操作给定的二叉树，将其变换为源二叉树的镜像。
要求： 空间复杂度 O(n) 。本题也有原地操作，即空间复杂度 O(1)的解法，时间复杂度 O(n)

比如：
源二叉树
镜像二叉树

示例1
输入：
{8,6,10,5,7,9,11}
返回值：
{8,10,6,11,9,7,5}
说明：
如题面所示

示例2
输入：
{}
返回值：
{}

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func Mirror(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := Mirror(root.Left)
	right := Mirror(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func main10() {
	// test
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
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

	result := Mirror(&node)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Left.Val)
	fmt.Printf("%+v\n", result.Right.Val)
	fmt.Printf("%+v\n", result.Left.Left.Val)
	fmt.Printf("%+v\n", result.Left.Right.Val)
	fmt.Printf("%+v\n", result.Right.Left.Val)
	fmt.Printf("%+v\n", result.Right.Right.Val)
}
