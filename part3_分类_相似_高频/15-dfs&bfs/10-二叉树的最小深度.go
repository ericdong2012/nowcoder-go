package main

import "fmt"

/*
NC234 二叉树的最小深度
https://www.nowcoder.com/practice/6a7f40d7696d46f79c74c61179993be6?tpId=117&tqId=39471&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=

描述
给定一颗节点数为N的二叉树，求其最小深度。最小深度是指树的根节点到最近叶子节点的最短路径上节点的数量。
（注：叶子节点是指没有子节点的节点。）
例如当输入{1,2,3,4,5}时，对应的二叉树如下图所示：
可以看出离根节点最近的叶子节点是节点值为3的节点，所以对应的输出为2。

示例1
输入：
{1,2,3,4,5}
返回值：
2

示例2
输入：
{1,#,2,#,3}
返回值：
3

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func run(root *TreeNode) int {
	// 层序遍历 + min
	if root == nil {
		return 0
	}
	res := 0
	q := []*TreeNode{root}

	for len(q) != 0 {
		//newValue := make([]int, 0)
		newQueue := make([]*TreeNode, 0)
		res ++
		for i := 0; i < len(q); i++ {
			//newValue = append(newValue, q[i].Val)
			if q[i].Left == nil && q[i].Right == nil {
				return res
			}
			if q[i].Left != nil {
				newQueue = append(newQueue, q[i].Left)
			}
			if q[i].Right != nil {
				newQueue = append(newQueue, q[i].Right)
			}
		}
		//res = min(res, temp)
		q = newQueue
	}
	return  0
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	node := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	//node := TreeNode{
	//	Val:  1,
	//	Left: nil,
	//	Right: &TreeNode{
	//		Val:  2,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}

	i := run(&node)
	fmt.Println(i)
}
