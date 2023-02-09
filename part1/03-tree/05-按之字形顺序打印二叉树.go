package main

import "fmt"

/*
BM27 按之字形顺序打印二叉树
https://www.nowcoder.com/practice/91b69814117f4e8097390d107d2efbe0?tpId=295&tqId=23454&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个二叉树，返回该二叉树的之字形层序遍历，（第一层从左向右，下一层从右向左，一直这样交替）

数据范围：0 \le n \le 15000≤n≤1500,树上每个节点的val满足 |val| <= 1500∣val∣<=1500
要求：空间复杂度：O(n)O(n)，时间复杂度：O(n)O(n)
例如：
给定的二叉树是{1,2,3,#,#,4,5}

该二叉树之字形层序遍历的结果是
[
[1],
[3,2],
[4,5]
]
示例1
输入：
{1,2,3,#,#,4,5}
复制
返回值：
[[1],[3,2],[4,5]]
复制
说明：
如题面解释，第一层是根节点，从左到右打印结果，第二层从右到左，第三层从左到右。
示例2
输入：
{8,6,10,5,7,9,11}
复制
返回值：
[[8],[10,6],[5,7,9,11]]
复制
示例3
输入：
{1,2,3,4,5}
复制
返回值：
[[1],[3,2],[4,5]]

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func Print(pRoot *TreeNode) [][]int {
	// write code here
	/*
		方式一： 将层序遍历中的奇偶做个判断,  层序遍历 + 双端队列
		方式二： 栈 + 层序遍历
	*/
	if pRoot == nil {
		return nil
	}
	res := [][]int{}
	q := []*TreeNode{pRoot}
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
		height += 1
		if height%2 == 0 {
			// reverse
			reverse(temp)
		}
		res = append(res, temp)
		q = newQueue
	}
	return res
}

func reverse(ans []int) {
	i, j := 0, len(ans)-1
	for i < j {
		p := ans[i]
		ans[i] = ans[j]
		ans[j] = p
		j--
		i++
	}
}

func main() {
	// test
	node := TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 10,
			Left: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   11,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := Print(&node)
	fmt.Printf("%+v\n", result)
}
