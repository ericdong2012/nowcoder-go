package main

import "fmt"

/*
BM37 二叉搜索树的最近公共祖先
https://www.nowcoder.com/practice/d9820119321945f588ed6a26f0a6991f?tpId=295&tqId=2290592&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
1.对于该题的最近的公共祖先定义:对于有根树T的两个节点p、q，最近公共祖先LCA(T,p,q)表示一个节点x，满足x是p和q的祖先且x的深度尽可能大。在这里，一个节点也可以是它自己的祖先.
2.二叉搜索树是若它的左子树不空，则左子树上所有节点的值均小于它的根节点的值； 若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值
3.所有节点的值都是唯一的。
4.p、q 为不同节点且均存在于给定的二叉搜索树中。
数据范围:
3<=节点总数<=10000
0<=节点值<=10000

如果给定以下搜索二叉树: {7,1,12,0,4,11,14,#,#,3,5}，如下图:


示例1
输入：
{7,1,12,0,4,11,14,#,#,3,5},1,12
返回值：
7
说明：
节点1 和 节点12的最近公共祖先是7

示例2
输入：
{7,1,12,0,4,11,14,#,#,3,5}, 12,11
返回值：
12
说明：
因为一个节点也可以是它自己的祖先. 所以输出12

*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//func lowestCommonAncestor(root *TreeNode, p int, q int) int {
//	// write code here
//	//node := Dfs2(root, p, q)
//	//return node.Val
//	if root == nil {
//		return -1
//	}
//	if root.Val == p || root.Val == q {
//		return root.Val
//	}
//
//	left := lowestCommonAncestor(root.Left, p, q)
//	right := lowestCommonAncestor(root.Right, p, q)
//	if left != -1 && right != -1 {
//		return root.Val
//	}
//	if left != -1 {
//		return left
//	}
//	if right != -1 {
//		return right
//	}
//	return -1
//}

//func Dfs2(root *TreeNode, p int, q int) *TreeNode {
//	if root == nil || root.Val == p || root.Val == q {
//		return root
//	}
//	if p < root.Val && q < root.Val {
//		return Dfs2(root.Left, p, q)
//	} else if p > root.Val && q > root.Val {
//		return Dfs2(root.Right, p, q)
//	} else {
//		return root
//	}
//}

func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
	// write code here
	node := lCA(root, o1, o2)
	//if node == nil {
	//	return 0
	//}
	return node.Val
}

func lCA(root *TreeNode, o1 int, o2 int) *TreeNode {
	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}
	//  lCA求得o1、o2的最近公共祖先，left去左子树中寻找，right去右子树中寻找
	//  left为空说明左子树中不存在，我们去右子树中寻找即可；right为空说明右子树中不存在，我们去左子树中寻找；当left、right均不为空时，说明两个节点分别坐落于left和right中，返回
	//  此时的根节点root即可
	left := lCA(root.Left, o1, o2)
	right := lCA(root.Right, o1, o2)
	// 左边没有，右边有
	if left == nil {
		return right
		//	右边没有，左边有
	} else if right == nil {
		return left
	} else {
		// left, right都存在， 一个在左子树，一个在右子树，显然root就是答案
		return root
	}
}


func main14() {
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := lowestCommonAncestor(&node, 1, 5)
	fmt.Printf("%+v\n", result)
}
