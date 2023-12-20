package main

import "fmt"

/*
BM36 判断是不是平衡二叉树
https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222?tpId=295&tqId=23250&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
输入一棵节点数为 n 二叉树，判断该二叉树是否是平衡二叉树。
平衡二叉树（Balanced Binary Tree）：
它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是平衡二叉树。  https://zhuanlan.zhihu.com/p/56066942

示例1
输入：
{1,2,3,4,5,6,7}
返回值：
true

示例2
输入：
{}
返回值：
true
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
思路1：
1.判断以根结点的树是否为二叉平衡树。求出左右子树的高度，判断它们的高度差是否超过了1。
2.递归判断根的左子树是否为平衡二叉树
3.递归判断根的右子树是否为平衡二叉树
注意：空树也是平衡二叉树

思路2：
step 1：第一个函数递归遍历二叉树所有节点(左右节点)。
step 2：对于每个节点判断，调用第二个函数获取子树深度。
step 3：第二个函数递归获取子树深度，只需要不断往子节点深度遍历，累加左右深度的较大值, left > right ? left +1 : right+1
step 4：根据深度判断该节点下的子树是否为平衡二叉树，即: abs(left - right ) <=1

  1
  2 二层的深度需要至上而下一直遍历
  3 三层的深度需要至上而下一直遍历
  4

https://blog.csdn.net/qq_44750696/article/details/124513713
https://blog.csdn.net/hansionz/article/details/82745625
// 刷题模板
https://www.nowcoder.com/practice/8b3b95850edb4115918ecebdf1b4d222?tpId=295&tqId=23250&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295
*/
func IsBalanced_Solution(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	// {1,2,3,4,#,#,5,6}  防止单侧层级不满足要求
	if !IsBalanced_Solution(pRoot.Left) {
		return false
	}
	if !IsBalanced_Solution(pRoot.Right) {
		return false
	}
	// 判断左右子树的高度差 是否小于1
	return abs(getMaxDepth(pRoot.Left)-getMaxDepth(pRoot.Right)) <= 1
}

// 获取左子树或者右子树的深度
func getMaxDepth(pRoot *TreeNode) int {
	if pRoot == nil {
		return 0
	}
	h1 := getMaxDepth(pRoot.Left)
	h2 := getMaxDepth(pRoot.Right)
	// 加上根节点
	return 1 + max1(h1, h2)
}

func main() {
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
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
	}

	result := IsBalanced_Solution(&node)
	fmt.Printf("%+v\n", result)
}
