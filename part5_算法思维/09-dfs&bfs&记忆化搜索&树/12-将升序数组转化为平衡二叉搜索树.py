"""
NC11 将升序数组转化为平衡二叉搜索树
https://www.nowcoder.com/practice/7e5b00f94b254da599a9472fe5ab283d?tpId=117&tqId=37720&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D196

给定一个升序排序的数组，将其转化为平衡二叉搜索树（BST）.
平衡二叉搜索树指 树上每个节点 node 都满足左子树中所有节点的的值都小于 node 的值，
右子树中所有节点的值都大于 node 的值，并且左右子树的节点数量之差不大于1

示例1
输入：
[-1,0,1,2]
返回值：
{1,0,2,-1}

示例2
输入：
[]
返回值：
{}
"""


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def sorted_array_to_bst(num):
    """write code here"""
    # 二分 + 递归，利用 BST 中序遍历为升序序列的特性
    if len(num) == 0:
        return None
    return create_bst(num, 0, len(num) - 1)


def create_bst(num, left, right):
    if left > right:
        return None
    mid = (left + right) // 2
    root = TreeNode(val=num[mid])
    root.left = create_bst(num, left, mid - 1)
    root.right = create_bst(num, mid + 1, right)
    return root
