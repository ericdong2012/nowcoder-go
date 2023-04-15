package main

/*
NC195 二叉树的直径
https://www.nowcoder.com/practice/15f977cedc5a4ffa8f03a3433d18650d?tpId=117&tqId=39370&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=5051,1263&title=

描述
给定一颗二叉树，求二叉树的直径。
1.该题的直径定义为：树上任意两个节点路径长度的最大值
2.该题路径长度定义为：不需要从根节点开始，也不需要在叶子节点结束，也不需要必须从父节点到子节点，一个节点到底另外一个节点走的边的数目
3.这个路径可能穿过根节点，也可能不穿过
4.树为空时，返回 0
如，输入{1,2,3,#,#,4,5}，二叉树如下:

那么：
从4到5的路径为4=>3=>5，路径长度为2
从4到2的路径为4=>3=>1=>2，路径长度为3

如，输入{1,2,3,#,#,4,5,9,#,#,6,#,7,#,8}，二叉树如下:
那么路径长度最长为:7=>9=>4=>3=>5=>6=>8，长度为6

示例1
输入：
{1,2,3,#,#,4,5}
返回值：
3

示例2
输入：
{1,2,3,#,#,4,5,9,#,#,6,#,7,#,8}
返回值：
6

示例3
输入：
{1,2,3}
返回值：
2

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

var res int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	// write code here
	// 树上任意两个节点路径长度的最大值
	/*
			以 root 为头结点的树，最大直径就是左右两个子树的高度相加
			左子树的直径是以 root->left 为头结点的树，最大直径同理
			右子树的直径是以 root->right 为头结点的树，最大直径同理

			int result = 0;
		    int dfs(TreeNode* root){
		        if(root == NULL)
		            return 0;
		        int left = dfs(root->left);
		        int right = dfs(root->right);
		        result = max(result,max(left,right) + 1);
		        result = max(result,left + right + 1);
		        return (left > right ? left : right) + 1;
		    }
		    int diameterOfBinaryTree(TreeNode* root) {
		        // write code here
		        if(root == NULL)
		            return 0;
		        dfs(root);
		        return result - 1;
		    }
	*/
	if root == nil {
		return 0
	}

	dfs2(root)
	return res

}

func dfs2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs2(root.Left)
	right := dfs2(root.Right)
	res = max2(res, max2(left, right)+1)
	res = max2(res, left+right+1)

	if left > right {
		return left + 1
	}
	return right + 1

}

func max2(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}
