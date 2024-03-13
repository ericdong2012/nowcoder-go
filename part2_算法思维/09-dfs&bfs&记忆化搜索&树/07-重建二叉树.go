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

示例1
输入：
[1,2,4,7,3,5,6,8],[4,7,2,1,5,3,8,6]
返回值：
{1,2,3,4,#,5,6,#,7,#,#,8}
说明：
返回根节点，系统会输出整颗二叉树对比结果，重建结果如题面图示

示例2
输入：
[1],[1]
返回值：
{1}

示例3
输入：
[1,2,3,4,5,6,7],[3,2,4,1,6,5,7]
返回值：
{1,2,5,3,4,6,7}

*/

/*
从一颗完全二叉树抽象出来（索引才找的对）

	   1
  2         5
3   4     6   7

前序： [1, 2 , 3, 4, 5, 6, 7]
中序： [3, 2 , 4, 1, 6, 5, 7]
根节点： 1



	   1
  左        右
  2         5
3   4     6   7

递归左右子树

左子树根节点： 2
左前序：[2,3,4]
左中序：[3,2,4]

右子树根节点： 5
右前序： [5, 6, 7]
右中序： [6, 5, 7]

*/
func reConstructBinaryTree(pre []int, vin []int) *TreeNode {
	// write code here
	if len(pre) == 0 || len(vin) == 0 {
		return nil
	}
	// 根节点的值 就是pre[0]
	root := &TreeNode{Val: pre[0]}
	// 找到根节点在中序遍历中的索引, 此处就是分界线
	idx := findIndex(vin, pre[0])
	// 先重建左边的树 [2, 3, 4] [3, 2, 4]
	// 前序的左边  和  中序的左边
	root.Left = reConstructBinaryTree(pre[1:idx+1], vin[:idx])
	// 重建右节点  [5, 6, 7]  [6, 5, 7]
	// 前序的右边  和  中序的右边
	root.Right = reConstructBinaryTree(pre[idx+1:], vin[idx+1:])

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
