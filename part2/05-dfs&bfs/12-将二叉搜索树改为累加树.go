package main

/*
NC215 将二叉搜索树改为累加树
https://www.nowcoder.com/practice/19bff16ca0d64d6da38ed24fd2903ce9?tpId=117&tqId=39419&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=


描述
给定一个二叉搜索树，树上的节点各不相同，请你将其修改为累加树，使每个节点的值变成原树中更大节点之和。
二叉搜索树的定义是任一节点的左子树的任意节点的值小于根节点的值，右子树则相反。

示例1
输入：
{0,#,1}
返回值：
{1,#,1}

示例2
输入：
{1,0,2}
返回值：
{3,3,2}

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res2 int   = 0
func convertBST(root *TreeNode) *TreeNode {
	// write code here
	/*
		由于二叉搜索树的左子树节点值总是小于当前节点，右子树节点值总是大于当前节点。
		所以按右、根、左顺序遍历时，树中的节点值一定是从大到小变化的，只要在遍历的过程中记录当前累加和，即可得到当前节点即将被替换的值
	*/
	if root == nil {
		return nil
	}

	dfs3(root)
	return root
}

func dfs3(root *TreeNode) {
	if root == nil {
		return
	}
	//按右、根、左顺序遍历整棵树
	dfs3(root.Right)
	res2 += root.Val
	root.Val = res
	dfs3(root.Left)
}
