package main

import "fmt"

/*
BM41 输出二叉树的右视图
https://www.nowcoder.com/practice/c9480213597e45f4807880c763ddd5f0?tpId=295&tqId=1073834&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
请根据二叉树的前序遍历，中序遍历 恢复二叉树，并打印出二叉树的右视图（从右往左看）

如输入[1,2,4,5,3], [4,2,5,1,3]时
通过前序遍历的结果[1,2,4,5,3]和中序遍历的结果[4,2,5,1,3]
可重建出以下二叉树：
对应的右视图是[1,3,5]

示例1
输入：
[1,2,4,5,3],[4,2,5,1,3]
返回值：
[1,3,5]

*/

/*
先做17题重建二叉树

可以发现解这道题，我们有两个步骤：
	建树
	打印右视图

首先建树方面，前序遍历是根左右的顺序，中序遍历是左根右的顺序，因为节点值互不相同，我们可以根据在前序遍历中找到根节点（每个子树部分第一个就是），再在中序遍历中找到对应的值，从其左右分割开，左边就是该树的左子树，右边就是该树的右子树，于是将问题划分为了子问题。

而打印右视图即找到二叉树每层最右边的节点元素，我们可以采取dfs（深度优先搜索）遍历树，根据记录的深度找到最右值。
具体做法：
step 1：首先检查两个遍历序列的大小，若是为0，则空树不用打印。
step 2：建树函数根据上述说，每次利用前序遍历第一个元素就是根节点，在中序遍历中找到它将二叉树划分为左右子树，利用l1 r1 l2 r2分别记录子树部分在数组中分别对应的下标，并将子树的数组部分送入函数进行递归。
step 3：dfs打印右视图时，使用哈希表存储每个深度对应的最右边节点，初始化两个栈辅助遍历，第一个栈记录dfs时的节点，第二个栈记录遍历到的深度，根节点先入栈。
step 4：对于每个访问的节点，每次左子节点先进栈，右子节点再进栈，这样访问完一层后，因为栈的先进后出原理，每次都是右边被优先访问，因此我们在哈希表该层没有元素时，添加第一个该层遇到的元素就是最右边的节点。
step 5：使用一个变量逐层维护深度最大值，最后遍历每个深度，从哈希表中读出每个深度的最右边节点加入数组中。

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func solve(xianxu []int, zhongxu []int) []int {
	// write code here
	if len(xianxu) == 0 || len(zhongxu) == 0 {
		return []int{}
	}
	// 重建树
	root := dfs18(xianxu, zhongxu)
	// 做层序遍历
	// 一个栈维护遍历到的节点  左节点先进，右节点后进
	// 一个栈维护深度
	// 遍历每层深度，从hash表中取出右边的节点加入到结果中
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	res := make([]int, 0)
	for len(queue) != 0 {
		// n 代表节点个数
		n := len(queue)
		// 实现了一个层序遍历
		for i := 0; i < n; i++ {
			// 根, 左节点
			node := queue[0]
			// 子节点（右节点，也可以添加进左右节点）
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			// 走到最后，取出最后一个节点的值（可能是根节点，也可能是右节点）
			if i == n-1 {
				res = append(res, node.Val)
			}
		}
	}

	return res
}

func dfs18(xianxu []int, zhongxu []int) *TreeNode {
	if len(xianxu) == 0 {
		return nil
	}
	root := &TreeNode{Val: xianxu[0]}
	index := findIndex2(zhongxu, xianxu[0])
	root.Left = dfs18(xianxu[1:index+1], zhongxu[:index])
	root.Right = dfs18(xianxu[index+1:], zhongxu[index+1:])
	return root
}

func findIndex2(arr []int, val int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func main18() {
	//node := TreeNode{
	//	Val: 1,
	//	Left: &TreeNode{
	//		Val:   2,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//	Right: &TreeNode{
	//		Val: 3,
	//		Left: &TreeNode{
	//			Val:   4,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   5,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}

	result := solve([]int{1, 2, 4, 5, 3}, []int{4, 2, 5, 1, 3})
	fmt.Printf("%+v\n", result)
}

/*
很不错的代码

func solve( xianxu []int ,  zhongxu []int ) []int {
    // write code here
    if len(xianxu) == 0 {
        return nil
    }
	// 构建树
    root := buildTree(xianxu,zhongxu)

    queue := make([]*TreeLinkNode, 0)
    queue = append(queue, root)
    res := make([]int,0)
    for len(queue) != 0 {
        n := len(queue)
        for i := 0 ; i < n ;i++ {
            node := queue[0]
            queue = queue[1:]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
            if i == n-1{
                res = append(res,node.Val)
            }
        }
    }

    return res
}

func buildTree(preorder []int, inorder []int) *TreeLinkNode {
    if len(preorder) == 0 {
        return nil
    }

    root := preorder[0]
    var index int
    for idx ,v := range inorder {
        if v == root {
            index = idx
            break
        }
    }

    return &TreeLinkNode{
        Val: root,
        Left: buildTree(preorder[1:index+1],inorder[:index]),
        Right: buildTree(preorder[index+1:],inorder[index+1:]),
    }
}

*/
