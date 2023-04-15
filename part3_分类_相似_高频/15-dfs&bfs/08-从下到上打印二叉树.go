package main

/*
NC224 从下到上打印二叉树
https://www.nowcoder.com/practice/ed982e032ff04d6a857b4cb4e6369d04?tpId=117&tqId=39449&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=


描述
给定一棵二叉树，返回齐自底向上的层序遍历。

数据范围：二叉树上节点数满足 1 \le n \le 1000 \1≤n≤1000  ，二叉树上的值满足 |val| \le 10^9 \∣val∣≤10
9


样例图：

示例1
输入：
{1,2,3,4,#,5,6}
返回值：
[[4,5,6],[2,3],[1]]

示例2
输入：
{1,2}
返回值：
[[2],[1]]

*/

func levelOrderBottom(root *TreeNode) [][]int {
	// write code here
	if root == nil {
		return nil
	}

	res := [][]int{}
	q := []*TreeNode{root}

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
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}
