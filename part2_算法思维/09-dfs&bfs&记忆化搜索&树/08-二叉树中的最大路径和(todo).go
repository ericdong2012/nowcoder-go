package main

import (
	"fmt"
)

/*

NC6 二叉树中的最大路径和
https://www.nowcoder.com/practice/da785ea0f64b442488c125b441a4ba4a?tpId=117&tqId=37716&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117&difficulty=undefined&judgeStatus=undefined&tags=583&title=

描述
二叉树里面的路径被定义为:从该树的任意节点出发，经过父=>子 或者  子=>父的连接，达到任意节点的序列。
注意:
1. 同一个节点在一条二叉树路径里中最多出现一次
2. 一条路径至少包含一个节点，且不一定经过根节点

给定一个二叉树的根节点root，请你计算它的最大路径和
给出以下的二叉树:
	2
1       3
最优路径是:2=>1=>3，或者3=>1=>2，最大路径和 2+1+3=6

示例1
输入:
{1,2,3}
 	1
2      3
返回值:
6

示例2
输入:
{-20,8,20,#,#,15,6}
   -20
8       20
	 15    6
返回值:
41
说明:
其中一条最大路径为: 15=>20=>6，路径和为 15 + 20 + 6 = 41

示例3
输入:
{-2,#,-3}
    -2
		-3
返回值:
-2

*/

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

/*

public int backtrace(TreeNode root){
    if(root == null) return 0;
    // 后序遍历，最后计算和，避免根节点值重复计算
    // 计算左、右子结点最大路径和
    int left = backtrace(root.left);
    int right = backtrace(root.right);
    int cur = root.val;
    if(left > 0){
        // 获取左子节点值
        cur += left;
    }
    if(right > 0){
        // 获取右子节点值
        cur += right;
    }
    // 统计出现的最大路径和
    max = Math.max(max, cur);
    // 返回当前最大路径和，不能直接返回cur，因为cur包含了两条路径，而返回只能一条路径
    return Math.max(root.val, Math.max(left, right) + root.val);
}

*/


/*
【A,B,C,D,E,null,F】


以A为根节点的这棵树的最大路径和：
单纯的存在于以B为根节点的子树中，单纯的存在于以C为根节点的子树中A节点到达左子树B中某节点的路径，A节点到达右子树C中某节点的路径，左子树B中某节点到达右子树C中某节点的路径，必然经过A节点A节点(假如左、右子树的最大路径均为负数)
因为需要获得从A节点到达其子树中某节点的路径和，所以我们想到深度优先遍历。又因为计算以A为根节点的树的最大路径和计算其子树B、C中的最大路径和函数功能相同，所以同样采用递归算法。


class Solution:
    def maxPathSum(self , root ):
        # write code here
        # 初始化 最小值
        res = float('-inf')

        def dfs(root):
            nonlocal res
            if not root:
                return float('-inf')
            # 递归左右子树
            left_max = dfs(root.left)
            right_max = dfs(root.right)
            # 分四种情况
            res = max(res, left_max+root.val, right_max+root.val, right_max+left_max+root.val, root.val)
            return max(root.val, root.val+left_max, root.val+right_max)

        # 深度搜索遍历
        dfs(root)
        return res

*/
var res = -10000000

func maxPathSum(root *TreeNode) int {
	dfs(root)
	return res
}

// 解法一
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 后序遍历
	// 分别求出左右子树的最大贡献度
	// 只有在最大贡献值大于 0 时，才会选取对应子节点
	left := max(dfs(root.Left), 0)
	right := max(dfs(root.Right), 0)
	// 根， 根 + 左， 根 + 右  三者的较大值
	curSum := max(root.Val, max(root.Val+right, root.Val+left))
	// 较大值 和 根 + 左 + 右  二者的比较
	curMax := max(curSum, left+right+root.Val)
	// 最大值
	res = max(res, curMax)
	// todo  why  返回节点的最大贡献值   node.val + max(leftGain, rightGain)
	return curSum
}

//解法2
func dfs2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 后序遍历
	// 分别求出左右子树的最大贡献度
	// 只有在最大贡献值大于 0 时，才会选取对应子节点
	left := max(dfs(root.Left), 0)
	right := max(dfs(root.Right), 0)
	// 统计出现的最大路径和 =  左子树的最大深度加右子树的最大深度加本节点，  如果其中有负值，该路线会被舍弃掉
	res = max(res, root.Val + left + right)
	// 返回当前节点的最大路径
	num := max(left, right) + root.Val
	return num
}



func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main19() {
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

	result := maxPathSum(&node)
	fmt.Printf("%+v\n", result)
}
