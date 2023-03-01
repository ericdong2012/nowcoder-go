package main

import "fmt"

/*
BM24 二叉树的中序遍历
https://www.nowcoder.com/practice/0bf071c135e64ee2a027783b80bf781d?tpId=295&tqId=1512964&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个二叉树的根节点root，返回它的中序遍历结果。
进阶：空间复杂度 O(n)，时间复杂度 O(n)


示例1
输入：
{1,2,#,#,3}
返回值：
[2,3,1]

示例2
输入：
{}
返回值：
[]

示例3
输入：
{1,2}
返回值：
[2,1]

示例4
输入：
{1,#,2}
返回值：
[1,2]


*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func inorderTraversal(root *TreeNode) []int {
	// write code here
	var res []int
	helper1(root, &res)
	return res
}

func helper1(nd *TreeNode, res *[]int) {
	if nd == nil {
		return
	}
	helper1(nd.Left, res)
	*res = append(*res, nd.Val)
	helper1(nd.Right, res)
}
func main() {
	// test
	node := TreeNode{
		Val:   1,
		Right: nil,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Left: nil,
		},
	}

	result := inorderTraversal(&node)
	fmt.Printf("%+v\n", result)
}
