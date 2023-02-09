package main

import "fmt"

/*
NC58 找到搜索二叉树中两个错误的节点
https://www.nowcoder.com/practice/4582efa5ffe949cc80c136eeb78795d6?tpId=117&tqId=37820&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=


描述
一棵二叉树原本是搜索二叉树，但是其中有两个节点调换了位置，使得这棵二叉树不再是搜索二叉树，请按升序输出这两个错误节点的值。(每个节点的值各不相同)
搜索二叉树：满足每个节点的左子节点小于当前节点，右子节点大于当前节点。

数据范围：3 \le n \le 1000003≤n≤100000,节点上的值满足 1 \le val \le n1≤val≤n ，保证每个value各不相同
进阶：空间复杂度 O(1)，时间复杂度 O(n)
示例1
输入：
{1,2,3}
返回值：
[1,2]


示例2
输入：
{4,2,5,3,1}
返回值：
[1,3]

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/*func findError(root *TreeNode) []int {
	// write code here
	res := make([]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) != 0 {
		newQ := make([]*TreeNode, 0)
		for i := 0; i < len(q); i++ {
			if q[i].Left == nil && q[i].Right == nil {
				continue
			}
			if q[i].Left != nil && q[i].Right != nil && q[i].Left.Val > q[i].Val && q[i].Right.Val < q[i].Val {
				res = append(res, q[i].Right.Val)
				res = append(res, q[i].Left.Val)

			} else if q[i].Left != nil && q[i].Right != nil && q[i].Left.Val > q[i].Val && q[i].Right.Val > q[i].Val {
				res = append(res, q[i].Val)
				res = append(res, q[i].Left.Val)

			} else if q[i].Left != nil && q[i].Right != nil && q[i].Right.Val < q[i].Val && q[i].Left.Val < q[i].Val {
				res = append(res, q[i].Right.Val)
				res = append(res, q[i].Val)
			}

			newQ = append(newQ, q[i].Left)
			newQ = append(newQ, q[i].Right)
		}
		q = newQ
	}
	return res
}*/

func findError(root *TreeNode) []int {
	// write code here
	res := make([]int, 0)
	dfs(root, &res)
	r := make([]int, 2)
	// 查找两个错误的节点
	/*
		此时的情况必定是一个小的数放到了后面，一个大的数放到了前面
		从左往右找那个大的数，并记录下索引，
		从右往左找小的那个数，并且索引值必大于前面记录下的索引
	*/
	// 从左往右
	for i := 1; i < len(res); i++ {
		// 左侧比根节点大
		if res[i-1] > res[i] {
			r[1] = res[i-1]
			break
		}
	}
	// 从右往左
	for j := len(res) - 2; j >=0; j-- {
		// 右侧比根节点小
		if res[j+1] < res[j] {
			r[0] = res[j+1]
			break
		}
	}
	return r
}

func dfs(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, res)
	*res = append(*res, root.Val)
	dfs(root.Right, res)
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

	//node := TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   1,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   2,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:   5,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	numbers := findError(&node)
	fmt.Println(numbers)
}
