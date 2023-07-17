package main

import "fmt"

/*
BM23 二叉树的前序遍历
https://www.nowcoder.com/practice/5e2135f4d2b14eb8a5b06fab4c938635?tpId=295&tqId=2291302&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给你二叉树的根节点root, 返回它节点值的前序遍历。

示例1
输入:
{1,#,2,3}
返回值:
[1,2,3]
*/

/*
https://blog.csdn.net/u013834525/article/details/80421684

对于整棵树而言，A是根，A分别在最前面、中间、后面被遍历到
前序遍历的话，一棵树的根永远在左子树前面，左子树又永远在右子树前面，你就找他的起点好了。


终止条件： 当子问题到达叶子节点后，后一个不管左右都是空，因此遇到空节点就返回。
返回值： 每次处理完子问题后，就是将子问题访问过的元素返回，依次存入了数组中。
本级任务： 每个子问题优先访问这棵子树的根节点，然后递归进入左子树和右子树。
具体做法：

step 1：准备数组用来记录遍历到的节点值，Java可以用List，C++可以直接用vector。
step 2：从根节点开始进入递归，遇到空节点就返回，否则将该节点值加入数组。
step 3：依次进入左右子树进行递归。


		 A
	 B      C
  D    E  F   G

前序遍历： ABDECFG
中序遍历： DBEAFCG
后序遍历： DEBFGCA
层序遍历： ABCDEFG

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// write code here
	var res []int
	helper(root, &res)
	return res
}

func helper(nd *TreeNode, res *[]int) {
	if nd == nil {
		return
	}
	*res = append(*res, nd.Val)
	helper(nd.Left, res)
	helper(nd.Right, res)
}

func main01() {
	// test
	node := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	result := preorderTraversal(&node)
	fmt.Printf("%+v\n", result)
}
