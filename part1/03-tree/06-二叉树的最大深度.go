package main

import "fmt"

/*
BM28 二叉树的最大深度
https://www.nowcoder.com/practice/8a2b2bf6c19b4f23a9bdb9b233eefa73?tpId=295&tqId=642&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
求给定二叉树的最大深度，
深度是指树的根节点到任一叶子节点路径上节点的数量。
最大深度是所有叶子节点的深度的最大值。
（注：叶子节点是指没有子节点的节点。）

示例1
输入：
{1,2}
返回值：
2


示例2
输入：
{1,2,3,4,#,#,5}
返回值：
3

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func maxDepth(root *TreeNode) int {
	// write code here
	// 层序遍历，队列
	if root == nil {
		return 0
	}
	res := [][]int{}
	q := []*TreeNode{root}
	height := 0
	for q != nil {
		// 存储每层的节点值
		temp := []int{}
		// 存储每层的节点
		var newQueue []*TreeNode
		for i := 0; i < len(q); i++ {
			temp = append(temp, q[i].Val)
			if q[i].Left != nil {
				newQueue = append(newQueue, q[i].Left)
			}
			if q[i].Right != nil {
				newQueue = append(newQueue, q[i].Right)
			}
		}
		res = append(res, temp)
		q = newQueue
		height += 1
	}
	return height
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

	result := maxDepth(&node)
	fmt.Printf("%+v\n", result)
}
