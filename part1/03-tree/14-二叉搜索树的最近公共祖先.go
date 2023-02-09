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
复制
返回值：
7
复制
说明：
节点1 和 节点12的最近公共祖先是7
示例2
输入：
{7,1,12,0,4,11,14,#,#,3,5},12,11
复制
返回值：
12
复制
说明：
因为一个节点也可以是它自己的祖先.所以输出12

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	// write code here
	if root.Val > p && root.Val > q {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p && root.Val < q {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root.Val
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
