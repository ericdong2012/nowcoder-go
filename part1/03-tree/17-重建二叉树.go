package main

/*
BM40 重建二叉树
https://www.nowcoder.com/practice/8a19cbe657394eeaac2f6ea9b0f6fcf6?tpId=295&tqId=23282&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定节点数为 n 的二叉树的前序遍历和中序遍历结果，请重建出该二叉树并返回它的头结点。
例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}，则重建出如下图所示。


提示:
1.vin.length == pre.length
2.pre 和 vin 均无重复元素
3.vin出现的元素均出现在 pre里
4.只需要返回根结点，系统会自动输出整颗树做答案对比
数据范围：n \le 2000n≤2000，节点的值 -10000 \le val \le 10000−10000≤val≤10000
要求：空间复杂度 O(n)O(n)，时间复杂度 O(n)O(n)
示例1
输入：
[1,2,4,7,3,5,6,8],[4,7,2,1,5,3,8,6]
复制
返回值：
{1,2,3,4,#,5,6,#,7,#,#,8}
复制
说明：
返回根节点，系统会输出整颗二叉树对比结果，重建结果如题面图示
示例2
输入：
[1],[1]
复制
返回值：
{1}
复制
示例3
输入：
[1,2,3,4,5,6,7],[3,2,4,1,6,5,7]
复制
返回值：
{1,2,5,3,4,6,7}

*/

func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	if len(pre) == 0 || len(vin) == 0 {
		return nil
	}

	root := &TreeNode{Val: pre[0]}
	idx := findIndex(vin, pre[0])

	if idx > 0 {
		root.Left = reConstructBinaryTree(pre[1:idx+1], vin[0:idx])
	}
	if idx+1 < len(vin) {
		root.Right = reConstructBinaryTree(pre[idx+1:], vin[idx+1:])
	}

	return root
}

func findIndex(a []int, val int) int {
	for i, v := range a {
		if v == val {
			return i
		}
	}
	return -1
}
