"""
NC81 二叉搜索树的第k个节点
https://www.nowcoder.com/practice/57aa0bab91884a10b5136ca2c087f8ff?tpId=117&tqId=37783&rp=1&ru=/exam/oj&qru=/exam/oj&sourceUrl=%2Fexam%2Foj%3Fpage%3D1%26pageSize%3D50%26search%3D%26tab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D117

描述
给定一棵结点数为n 二叉搜索树，请找出其中的第 k 小的TreeNode结点值。
1.返回第k小的节点值即可
2.不能查找的情况，如二叉树为空，则返回-1，或者k大于n等等，也返回-1
3.保证n个节点的值不一样

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
"""

# 二叉搜索树本身是有序的，且满足左<根<右，
# 只需按照中序遍历的过程，将结果保存即可，取出第k个(索引 k-1)


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def kth_node(proot, k):
    """write code here"""
    if proot is None or k == 0:
        return -1
    res = []
    helper20(proot, res)
    if k > len(res):
        return -1
    return res[k - 1]


def helper20(root, res):
    if root is None:
        return
    helper20(root.left, res)
    res.append(root.val)
    helper20(root.right, res)
