package main

import "fmt"

/*
BM25 二叉树的后序遍历
https://www.nowcoder.com/practice/1291064f4d5d4bdeaefbf0dd47d78541?tpId=295&tqId=2291301&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个二叉树，返回他的后序遍历的序列。
后序遍历是值按照 左节点->右节点->根节点 的顺序的遍历。
数据范围：二叉树的节点数量满足 0 \le n \le 100 \0≤n≤100  ，二叉树节点的值满足 1 \le val \le 100 \1≤val≤100  ，树的各节点的值各不相同


示例1
输入：
{1,#,2,3}
返回值：
[3,2,1]

示例2
输入：
{1}
返回值：
[1]

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func postorderTraversal(root *TreeNode) []int {
	// write code here
	var res []int
	helper2(root, &res)
	return res
}

func helper2(nd *TreeNode, res *[]int) {
	if nd == nil {
		return
	}
	helper2(nd.Left, res)
	helper2(nd.Right, res)
	*res = append(*res, nd.Val)
}
func main() {
	// test
	node := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	result := postorderTraversal(&node)
	fmt.Printf("%+v\n", result)
}
