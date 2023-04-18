package main

import (
	"fmt"
)

/*
BM26 求二叉树的层序遍历
https://www.nowcoder.com/practice/04a5560e43e24e9db4595865dc9c63a3?tpId=295&tqId=644&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295


描述
给定一个二叉树，返回该二叉树层序遍历的结果，（从左到右，一层一层地遍历）
例如：
给定的二叉树是{3,9,20,#,#,15,7},
该二叉树层序遍历的结果是
[
[3],
[9,20],
[15,7]
]

示例1
输入：
{1,2}
返回值：
[[1],[2]]

示例2
输入:
{1,2,3,4,#,#,5}
返回值:
[[1],[2,3],[4,5]]

*/

/*
https://blog.csdn.net/u013834525/article/details/80421684
层序遍历嘛，就是按层，从上到下，从左到右遍历，这个没啥好说的。  bfs
*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
	// write code here
	/*
	   vector<vector<int>> vect;
	   if (root==NULL)
	   {
	       return vect;
	   }
	   //边界检测
	   queue<TreeNode*> que;
	   //创建队列
	   que.push(root);
	   //根节点入队
	   while (!que.empty())//队列不为空，继续循环
	   {
	       int LevelNum = que.size();
	       //记录每层节点的个数
	       vector<int> vec;
	       //记录每层节点的元素
	       for (int i = 0; i < LevelNum; i++)
	       {
	           TreeNode* t = que.front();
	           //接收队列中头结点元素
	           que.pop();
	           //删除头结点元素
	           vec.push_back(t->val);
	           //将值存放到数组中
	           if (t->left!=nullptr)
	           {
	               que.push(t->left);
	               //如果左子树不为空，将左子树的根节点装入队列
	           }
	           if (t->right!=nullptr)
	           {
	               que.push(t->right);
	               //如果右子树不为空，将右子树的根节点装入队列
	           }
	       }
	       vect.push_back(vec);
	   }
	   return vect;
	*/

	/*
	   if root is None: # 判断根节点是否为空
	       return
	   res = []
	   q = [root]
	   while q:
	       tep = []
	       for _ in range(len(q)): # 对于每一层进行遍历
	           node = q.pop(0)
	           tep.append(node.val)
	           # 判断左右节点是否存在
	           if node.left:
	               q.append(node.left)
	           if node.right:
	               q.append(node.right)
	       res.append(tep)
	   return res
	*/

	if root == nil {
		return nil
	}
	res := [][]int{}
	q := []*TreeNode{root}
	for q != nil {
		// 存储每层的节点值
		temp := []int{}
		// 存储每层的节点
		// newQueue := []*TreeNode{} 报错
		var newQueue []*TreeNode
		for i := 0; i < len(q); i++ {
			temp = append(temp, q[i].Val)
			if q[i].Left != nil {
				newQueue = append(newQueue, q[i].Left)
			}
			if q[i].Right != nil {
				newQueue = append(newQueue, q[i].Right)
			}
		}
		res = append(res, temp)
		q = newQueue
	}
	return res
}

func main04() {
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

	result := levelOrder(&node)
	fmt.Printf("%+v\n", result)
}
