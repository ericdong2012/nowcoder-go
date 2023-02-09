package main

/*
NC80 把二叉树打印成多行
https://www.nowcoder.com/practice/445c44d982d04483b04a54f298796288?tpId=117&tqId=37781&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=


描述
给定一个节点数为 n 二叉树，要求从上到下按层打印二叉树的 val 值，同一层结点从左至右输出，每一层输出一行，将输出的结果存放到一个二维数组中返回。
例如：
给定的二叉树是{1,2,3,#,#,4,5}

该二叉树多行打印层序遍历的结果是
[
[1],
[2,3],
[4,5]
]

输入描述：
给定一个二叉树的根节点
示例1
输入：
{1,2,3,#,#,4,5}
返回值：
[[1],[2,3],[4,5]]

示例2
输入：
{8,6,10,5,7,9,11}
返回值：
[[8],[6,10],[5,7,9,11]]

示例3
输入：
{1,2,3,4,5}
返回值：
[[1],[2,3],[4,5]]

示例4
输入：
{}
返回值：
[]


*/

func Print(pRoot *TreeNode) [][]int {
	// write code here
	// 层序遍历 打印

	if pRoot == nil {
		return nil
	}

	res := [][]int{}
	q := []*TreeNode{pRoot}

	for len(q) != 0 {
		newValue := make([]int, 0)
		newQueue := make([]*TreeNode, 0)

		for i := 0; i < len(q); i++ {
			newValue = append(newValue, q[i].Val)
			if q[i].Left != nil {
				newQueue = append(newQueue, q[i].Left)
			}
			if q[i].Right != nil {
				newQueue = append(newQueue, q[i].Right)
			}
		}

		res = append(res, newValue)
		q = newQueue
	}

	return res

}
