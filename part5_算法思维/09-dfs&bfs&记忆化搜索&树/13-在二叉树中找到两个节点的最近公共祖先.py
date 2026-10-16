"""
BM38 在二叉树中找到两个节点的最近公共祖先
https://www.nowcoder.com/practice/e0cc33a83afe4530bcec46eba3325116?tpId=295&tqId=1024325&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一棵二叉树(保证非空)以及这棵树上的两个节点对应的val值 o1 和 o2，请找到 o1 和 o2 的最近公共祖先节点
节点本身可以视为自己的祖先

示例1
输入：
{3,5,1,6,2,0,8,#,#,7,4},5,1
返回值：
3

示例2
输入：
{3,5,1,6,2,0,8,#,#,7,4},2,7
返回值：
2
"""

# 和11题不同的地方在于不是二叉搜索树
# dfs 找根节点


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def lowest_common_ancestor2(root, o1, o2):
    """write code here"""
    node = lca2(root, o1, o2)
    return node.val


def lca2(root, o1, o2):
    # break 条件：空节点或者根节点的值等于 o1, o2
    if root is None or root.val == o1 or root.val == o2:
        return root
    # lca2 求得 o1、o2 的最近公共祖先
    # left 去左子树中寻找，right 去右子树中寻找
    left = lca2(root.left, o1, o2)
    right = lca2(root.right, o1, o2)
    # 左边没有，右边有
    if left is None:
        return right
    # 右边没有，左边有
    elif right is None:
        return left
    else:
        # left, right 都存在，一个在左子树，一个在右子树，显然 root 就是答案
        return root
