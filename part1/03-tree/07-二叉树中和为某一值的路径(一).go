package main

import "fmt"

/*
BM29 二叉树中和为某一值的路径(一)
https://www.nowcoder.com/practice/508378c0823c423baa723ce448cbfd0c?tpId=295&tqId=634&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个二叉树root和一个值 sum ，判断是否有从根节点到叶子节点的节点值之和等于 sum 的路径。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n

例如：
给出如下的二叉树， sum=22，返回true，因为存在一条路径 5→ 4→ 11→ 2 的节点值之和为 22

数据范围：
要求：空间复杂度 O(n)，时间复杂度 O(n)
进阶：空间复杂度 O(树的高度)，时间复杂度 O(n)

示例1
输入：
{5,4,8,1,11,#,9,#,#,2,7},22
返回值：
true

示例2
输入：
{1,2},0
返回值：
false

示例3
输入：
{1,2},3
返回值：
true

示例4
输入：
{},0
返回值：
false

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 多数之和
// 把所有路径找出来，然后判断
// 从根节点往下走，走的时候减去当前节点的值，一直到叶子节点，如果减到最后，值等于叶子节点的值，说明存在这样的结果，直接返回true

func hasPathSum(root *TreeNode, sum int) bool {
	// write code here
	if root == nil {
		return false
	}
	sum -= root.Val
	// 判断叶子节点 还有sum
	if root.Left == nil && root.Right == nil && sum == 0 {
		return true
	}
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)
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

	result := hasPathSum(&node, 30)
	fmt.Printf("%+v\n", result)
}
