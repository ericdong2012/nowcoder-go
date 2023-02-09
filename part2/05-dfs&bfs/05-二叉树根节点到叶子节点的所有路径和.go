package main

import "fmt"

/*
NC5 二叉树根节点到叶子节点的所有路径和
https://www.nowcoder.com/practice/185a87cd29eb42049132aed873273e83?tpId=117&tqId=37715&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=

描述
给定一个二叉树的根节点root，该树的节点值都在数字\ 0-9 0−9 之间，每一条从根节点到叶子节点的路径都可以用一个数字表示。
1.该题路径定义为从树的根结点开始往下一直到叶子结点所经过的结点
2.叶子节点是指没有子节点的节点
3.路径只能从父节点到子节点，不能从子节点到父节点
4.总节点数目为n

例如根节点到叶子节点的一条路径是1\to 2\to 31→2→3,那么这条路径就用\ 123 123 来代替。
找出根节点到叶子节点的所有路径表示的数字之和
例如：

这颗二叉树一共有两条路径，
根节点到叶子节点的路径 1\to 21→2 用数字\ 12 12 代替
根节点到叶子节点的路径 1\to 31→3 用数字\ 13 13 代替
所以答案为\ 12+13=25 12+13=25

数据范围：节点数 0 \le n \le 1000≤n≤100，保证结果在32位整型范围内
要求：空间复杂度 O(n)，时间复杂度 O(n^2)
进阶：空间复杂度 O(n)，时间复杂度 O(n)

示例1
输入：
{1,2,3}
返回值：
25

示例2
输入：
{1,0}
返回值：
10

示例3
输入：
{1,2,0,3,4}
返回值：
257

*/

/*
从根节点往下走的时候，那么当前节点的值就是父节点的值*10+当前节点的值

*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//var sum int

//func sumNumbers(root *TreeNode) int {
//	// write code here
//	if root == nil {
//		return 0
//	}
//	sum = 0
//	dfs(root, 0)
//	return sum
//}
//
//func dfs(root *TreeNode, path int) {
//	if root == nil {
//		return
//	}
//	path = path*10 + root.Val
//	if root.Left == nil && root.Right == nil {
//		sum += path
//		return
//	}
//	dfs(root.Left, path)
//	dfs(root.Right, path)
//}

func sumNumbers(root *TreeNode) int {
	// write code here
	// 栈
	if root == nil {
		return 0
	}
	nodeStack := make([]*TreeNode, 0)
	valueStack := make([]int, 0)
	res := 0

	nodeStack = append(nodeStack, root)
	valueStack = append(valueStack, root.Val)

	for len(nodeStack) != 0 {
		tempNode := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		tempValue := valueStack[len(valueStack)-1]
		valueStack = valueStack[:len(valueStack)-1]

		if tempNode.Left == nil && tempNode.Right == nil {
			res += tempValue
		} else {
			if tempNode.Left != nil {
				nodeStack = append(nodeStack, tempNode.Left)
				valueStack = append(valueStack, tempValue*10+tempNode.Left.Val)
			}

			if tempNode.Right != nil {
				nodeStack = append(nodeStack, tempNode.Right)
				valueStack = append(valueStack, tempValue*10+tempNode.Right.Val)
			}
		}

	}
	return res
}

func main() {

	node := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	numbers := sumNumbers(&node)
	fmt.Println(numbers)
}
