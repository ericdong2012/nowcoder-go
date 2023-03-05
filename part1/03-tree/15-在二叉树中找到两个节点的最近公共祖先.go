package main

/*
BM38 在二叉树中找到两个节点的最近公共祖先
https://www.nowcoder.com/practice/e0cc33a83afe4530bcec46eba3325116?tpId=295&tqId=1024325&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点。

如当输入{3,5,1,6,2,0,8,#,#,7,4},5,1时，二叉树{3,5,1,6,2,0,8,#,#,7,4}如下图所示：
所以节点值为5和节点值为1的节点的最近公共祖先节点的节点值为3，所以对应的输出为3。
节点本身可以视为自己的祖先

示例1
输入：
{3,5,1,6,2,0,8,#,#,7,4},5,1
返回值：
3

示例2
输入：
{3,5,1,6,2,0,8,#,#,7,4},2,7
返回值：
2

*/

/*
和14 题不同的地方在于 不是二叉搜索树

重建二叉树， dfs 找根节点
*/

func lowestCommonAncestor2(root *TreeNode, o1 int, o2 int) int {
	// write code here
	node := lCA2(root, o1, o2)
	//if node == nil {
	//	return 0
	//}
	return node.Val
}

func lCA2(root *TreeNode, o1 int, o2 int) *TreeNode {
	if root == nil || root.Val == o1 || root.Val == o2 {
		return root
	}
	//  lCA求得o1、o2的最近公共祖先，left去左子树中寻找，right去右子树中寻找
	//  left为空说明左子树中不存在，我们去右子树中寻找即可；right为空说明右子树中不存在，我们去左子树中寻找；当left、right均不为空时，说明两个节点分别坐落于left和right中，返回
	//  此时的根节点root即可
	left := lCA2(root.Left, o1, o2)
	right := lCA2(root.Right, o1, o2)
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
