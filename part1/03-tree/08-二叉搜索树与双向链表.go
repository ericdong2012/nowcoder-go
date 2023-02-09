package main

import "fmt"

/*
BM30 二叉搜索树与双向链表
https://www.nowcoder.com/practice/947f6eb80d944a84850b0538bf0ec3a5?tpId=295&tqId=23253&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的双向链表。如下图所示


数据范围：输入二叉树的节点数 0 \le n \le 10000≤n≤1000，二叉树中每个节点的值 0\le val \le 10000≤val≤1000
要求：空间复杂度O(1)O(1)（即在原树上操作），时间复杂度 O(n)O(n)

注意:
1.要求不能创建任何新的结点，只能调整树中结点指针的指向。当转化完成以后，树中节点的左指针需要指向前驱，树中节点的右指针需要指向后继
2.返回链表中的第一个节点的指针
3.函数返回的TreeNode，有左右指针，其实可以看成一个双向链表的数据结构
4.你不用输出双向链表，程序会根据你的返回值自动打印输出
输入描述：
二叉树的根节点
返回值描述：
双向链表的其中一个头节点。
示例1
输入：
{10,6,14,4,8,12,16}
复制
返回值：
From left to right are:4,6,8,10,12,14,16;From right to left are:16,14,12,10,8,6,4;
复制
说明：
输入题面图中二叉树，输出的时候将双向链表的头节点返回即可。
示例2
输入：
{5,4,#,3,#,2,#,1}
复制
返回值：
From left to right are:1,2,3,4,5;From right to left are:5,4,3,2,1;
复制
说明：
                    5
                  /
                4
              /
            3
          /
        2
      /
    1
树的形状如上图

*/

/*
将中序遍历的结果用数组存储下来，得到的数组是有从小到大顺序的。最后将数组中的结点依次连接即可。
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

var prev *TreeNode
var head *TreeNode

func Convert(root *TreeNode) *TreeNode {
	// write code here
	/*
	       # write code here
	       # 首先进行中序排序
	       def inorder(root, res):
	            if not root:
	                return None
	            inorder(root.left, res)
	            res.append(root)
	            inorder(root.right, res)

	        res = []
	        if not pRootOfTree:
	            return None
	        inorder(pRootOfTree, res)
	        if len(res) == 1:
	            return pRootOfTree
	        # 构造双向链表
	        for i in range(len(res)-1):
	            res[i].right = res[i+1]
	            res[i+1].left = res[i]
	        return res[0]
	*/

	if root == nil {
		return nil
	}
	prev = nil
	head = nil
	inorder(root)
	return head
}
func inorder(cur *TreeNode) {
	if cur == nil {
		return
	}
	inorder(cur.Left)
	if prev == nil {
		head = cur
	} else {
		// 构建双向链表
		prev.Right = cur
		cur.Left = prev
	}
	prev = cur
	inorder(cur.Right)
}

func main() {
	// test
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}

	result := Convert(&node)
	fmt.Printf("%+v\n", result)
	fmt.Printf("%+v\n", result.Left.Val)
	fmt.Printf("%+v\n", result.Right.Val)
}
