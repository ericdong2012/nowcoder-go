package main

import (
	"fmt"
	"math"
)

/*
BM34 判断是不是二叉搜索树
https://www.nowcoder.com/practice/a69242b39baf45dea217815c7dedb52b?tpId=295&tqId=2288088&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个二叉树根节点，请你判断这棵树是不是二叉搜索树。

二叉搜索树: 每个节点的左子树上的所有节点均小于当前节点且右子树上的所有节点均大于当前节点。
它或者是一棵空树，或者是具有下列性质的 二叉树 ： 若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值

示例1
输入：
{1,2,3}
返回值：
false

示例2
输入：
{2,1,3}
返回值：
true

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

var pre = math.MaxInt32

func isValidBST(root *TreeNode) bool {
	// 方式一: 递归  二叉树的递归，则是将某个节点的左子树、右子树看成一颗完整的树，那么对于子树的访问或者操作就是对于原树的访问或者操作的子问题，因此可以自我调用函数不断进入子树
	/*
	   pre = -sys.maxsize - 1
	   def isValidBST(self , root: TreeNode) -> bool:
	       if not root:
	           return True
	       # 先进入左子树
	       if not self.isValidBST(root.left):
	           return False
	       if(root.val <= self.pre):
	           return False
	       # 更新最值
	       self.pre = root.val
	       # 再进入右子树
	       if not self.isValidBST(root.right):
	           return False
	       return True
	*/

	// 方式二: 栈
	/*
		二叉搜索树的特性就是中序遍历是递增序。既然是判断是否是二叉搜索树，那我们可以使用中序递归遍历。只要之前的节点是二叉树搜索树，那么如果当前的节点小于上一个节点值那么就可以向下判断。*只不过在过程中我们要求反退出。比如一个链表1->2->3->4，只要for循环遍历如果中间有不是递增的直接返回false即可。
		if(root.val < pre)
		    return false;

		具体做法：
			step 1：首先递归到最左，初始化maxLeft与pre。
			step 2：然后往后遍历整棵树，依次连接pre与当前节点，并更新pre。
			step 3：左子树如果不是二叉搜索树返回false。
			step 4：判断当前节点是不是小于前置节点，更新前置节点。
			step 5：最后由右子树的后面节点决定。
	*/

	/*
		class Solution:
		    def isValidBST(self , root: TreeNode) -> bool:
		        # 设置栈用于遍历 和 记录中序遍历的数组
		        s ,sort = [], []
		        head = root
		        while head or s:
		            # 直到没有左节点
		            while head:
		                s.append(head)
		                head = head.left
		            head = s[-1]
		            s.pop()
		            # 访问节点
		            sort.append(head.val)
		            head = head.right
		        # 遍历中序结果
		        for i in range(1, len(sort)):
		            # 一旦有降序，则不是搜索树
		            if sort[i - 1] > sort[i]:
		                return False
		        return True
	*/
	// 中序遍历中拿到相应的值，做判断
	if root == nil {
		return true
	}
	// 判断左边
	if !isValidBST(root.Left) {
		return false
	}
	// 此时左子节点满足要求，中序 判断root.Val 跟上一次的结果(上一次的根节点是当前的左节点)
	if root.Val < pre {
		return false
	}
	// 拿到中间的值
	pre = root.Val
	// 判断右边
	return isValidBST(root.Right)
}

func main11() {
	// 本地无法完成自测
	node := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 5,
		},
	}

	result := isValidBST(&node)
	fmt.Printf("%+v\n", result)
}
