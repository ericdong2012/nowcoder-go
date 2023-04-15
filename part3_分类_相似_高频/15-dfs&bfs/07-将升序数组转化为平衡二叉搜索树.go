package main

/*
NC11 将升序数组转化为平衡二叉搜索树
https://www.nowcoder.com/practice/7e5b00f94b254da599a9472fe5ab283d?tpId=117&tqId=37720&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=


描述
给定一个升序排序的数组，将其转化为平衡二叉搜索树（BST）.

平衡二叉搜索树指树上每个节点 node 都满足左子树中所有节点的的值都小于 node 的值，右子树中所有节点的值都大于 node 的值，并且左右子树的节点数量之差不大于1

数据范围：0 \le n \le 100000≤n≤10000，数组中每个值满足 |val| \le 5000∣val∣≤5000
进阶：空间复杂度 O(n)O(n) ，时间复杂度 O(n)O(n)

例如当输入的升序数组为[-1,0,1,2]时，转化后的平衡二叉搜索树（BST）可以为{1,0,2,-1}，如下图所示：

或为{0,-1,1,#,#,#,2}，如下图所示：

返回任意一种即可。
示例1
输入：
[-1,0,1,2]
复制
返回值：
{1,0,2,-1}
复制
示例2
输入：
[]
复制
返回值：
{}

*/

func create_bst(num []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	root := &TreeNode{Val: num[mid]}
	root.Left = create_bst(num, left, mid-1)
	root.Right = create_bst(num, mid+1, right)
	return root

}

func sortedArrayToBST(num []int) *TreeNode {
	// write code here
	// 二分 + 递归   利用BST中序遍历为升序序列的特性
	if len(num) == 0 {
		return nil
	}
	return create_bst(num, 0, len(num)-1)
}
