package main

/*
NC81 二叉搜索树的第k个节点
https://www.nowcoder.com/practice/57aa0bab91884a10b5136ca2c087f8ff?tpId=117&tqId=37783&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=583&title=


描述
给定一棵结点数为n 二叉搜索树，请找出其中的第 k 小的TreeNode结点值。
1.返回第k小的节点值即可
2.不能查找的情况，如二叉树为空，则返回-1，或者k大于n等等，也返回-1
3.保证n个节点的值不一样

如输入{5,3,7,2,4,6,8},3时，二叉树{5,3,7,2,4,6,8}如下图所示：
该二叉树所有节点按结点值升序排列后可得[2,3,4,5,6,7,8]，所以第3个结点的结点值为4，故返回对应结点值为4的结点即可。

示例1
输入：
{5,3,7,2,4,6,8},3
返回值：
4

示例2
输入：
{},1
返回值：
-1

*/

/*
二叉搜索树本身是有序的，且满足左< 根 <右， 只需按照中序遍历的过程，将结果保存即可，取出第k个(索引 k - 1 )
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func KthNode(proot *TreeNode, k int) int {
	// write code here
	if proot == nil || k == 0  {
		return -1
	}
	res := make([]int, 0)
	helper20(proot, &res)
	if k> len(res) {
		return -1
	}
	return res[k-1]
}

func helper20(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	helper20(root.Left, res)
	*res = append(*res, root.Val)
	helper20(root.Right, res)
}