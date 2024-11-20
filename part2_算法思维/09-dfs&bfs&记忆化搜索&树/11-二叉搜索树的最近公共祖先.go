package main

import "fmt"

/*
BM37 二叉搜索树的最近公共祖先
https://www.nowcoder.com/practice/d9820119321945f588ed6a26f0a6991f?tpId=295&tqId=2290592&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
1.对于该题的最近的公共祖先定义:对于有根树T的两个节点p、q，最近公共祖先LCA(T,p,q)表示一个节点x，满足x是p和q的祖先且x的深度尽可能大。在这里，一个节点也可以是它自己的祖先.
2.二叉搜索树是若它的左子树不空，则左子树上所有节点的值均小于它的根节点的值； 若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值
3.所有节点的值都是唯一的。
4.p、q 为不同节点且均存在于给定的二叉搜索树中。

示例1
输入：
{7,1,12,0,4,11,14,#,#,3,5}, 1, 12
返回值：
7
说明：
节点1 和 节点12 的最近公共祖先是 7

                7
           1        12
          0  4   11  14
            3 5

示例2
输入：
{7,1,12,0,4,11,14,#,#,3,5}, 12,11
返回值：
12
说明：
因为一个节点也可以是它自己的祖先. 所以输出12

*/
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/*
我们也可以利用二叉搜索树的性质：对于某一个节点若是p与q都小于等于这个这个节点值，说明p、q都在这个节点的左子树，而最近的公共祖先也一定在这个节点的左子树；若是p与q都大于等于这个节点，说明p、q都在这个节点的右子树，而最近的公共祖先也一定在这个节点的右子树。而若是对于某个节点，p与q的值一个大于等于节点值，一个小于等于节点值，说明它们分布在该节点的两边，而这个节点就是最近的公共祖先，因此从上到下的其他祖先都将这个两个节点放到同一子树，只有最近公共祖先会将它们放入不同的子树，每次进入一个子树又回到刚刚的问题，因此可以使用递归。

具体做法：
step 1：首先检查空节点，空树没有公共祖先。
step 2：对于某个节点，比较与p、q的大小，若p、q在该节点两边说明这就是最近公共祖先。
step 3：如果p、q都在该节点的左边，则递归进入左子树。
step 4：如果p、q都在该节点的右边，则递归进入右子树。

class Solution:
    def lowestCommonAncestor(self , root: TreeNode, p: int, q: int) -> int:
        # 空树找不到公共祖先
        if not root:
            return -1
        # pq在该节点两边说明这就是最近公共祖先
        if (p >= root.val and q <= root.val) or (p <= root.val and q >= root.val):
            return root.val
        # pq都在该节点的左边
        elif p <= root.val and q <= root.val:
            # 进入左子树
            return self.lowestCommonAncestor(root.left, p, q)
        #pq都在该节点的右边
        else:
            # 进入右子树
            return self.lowestCommonAncestor(root.right, p, q)

*/

func lowestCommonAncestor(root *TreeNode, p int, q int) int {
	// write code here
	node := query(root, p, q)
	return node.Val
}

func query(root *TreeNode, p, q int) *TreeNode {
	// 在根节点左边
	if root.Val > p && root.Val > q {
		return query(root.Left, p, q)
	}
	// 在根节点右边
	if root.Val < p && root.Val < q {
		return query(root.Right, p, q)
	}
	// 不左不右在中间，也就是根节点
	return root
}

//func lowestCommonAncestor(root *TreeNode, o1 int, o2 int) int {
//	// write code here
//	node := lCA(root, o1, o2)
//	//if node == nil {
//	//	return 0
//	//}
//	return node.Val
//}
//
//func lCA(root *TreeNode, o1 int, o2 int) *TreeNode {
//	// break 条件， 空节点或者根节点的值等于o1, o2
//	if root == nil || root.Val == o1 || root.Val == o2 {
//		return root
//	}
//	//  lCA求得o1、o2的最近公共祖先，left去左子树中寻找，right去右子树中寻找
//	//  left为空说明左子树中不存在，我们去右子树中寻找即可；right为空说明右子树中不存在，我们去左子树中寻找；当left、right均不为空时，说明两个节点分别坐落于left和right中，说明是root
//	left := lCA(root.Left, o1, o2)
//	right := lCA(root.Right, o1, o2)
//	// 左边没有，右边有
//	if left == nil {
//		return right
//		//	右边没有，左边有
//	} else if right == nil {
//		return left
//	} else {
//		// left, right都存在， 一个在左子树，一个在右子树，显然root就是答案
//		return root
//	}
//}

func main14() {
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := lowestCommonAncestor(&node, 1, 5)
	fmt.Printf("%+v\n", result)
}
