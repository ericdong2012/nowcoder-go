"""
BM37 二叉搜索树的最近公共祖先
https://www.nowcoder.com/practice/d9820119321945f588ed6a26f0a6991f?tpId=295&tqId=2290592&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
1.对于该题的最近的公共祖先定义:对于有根树T的两个节点p、q，
最近公共祖先LCA(T,p,q)表示一个节点x，满足x是p和q的祖先且x的深度尽可能大。
在这里，一个节点也可以是它自己的祖先.
2.二叉搜索树是若它的左子树不空，则左子树上所有节点的值均小于它的根节点的值；
若它的右子树不空，则右子树上所有节点的值均大于它的根节点的值
3.所有节点的值都是唯一的。
4.p、q 为不同节点且均存在于给定的二叉搜索树中。

示例1
输入：
{7,1,12,0,4,11,14,#,#,3,5}, 1, 12
返回值：
7

示例2
输入：
{7,1,12,0,4,11,14,#,#,3,5}, 12, 11
返回值：
12
"""

# 利用二叉搜索树的性质：对于某一个节点若是p与q都小于等于这个节点值，
# 说明p、q都在这个节点的左子树，而最近的公共祖先也一定在这个节点的左子树；
# 若是p与q都大于等于这个节点，说明p、q都在这个节点的右子树，
# 而最近的公共祖先也一定在这个节点的右子树。
# 而若是对于某个节点，p与q的值一个大于等于节点值，一个小于等于节点值，
# 说明它们分布在该节点的两边，而这个节点就是最近的公共祖先。


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def lowest_common_ancestor(root, p, q):
    """write code here"""
    node = query(root, p, q)
    return node.val


def query(root, p, q):
    # 在根节点左边
    if root.val > p and root.val > q:
        return query(root.left, p, q)
    # 在根节点右边
    if root.val < p and root.val < q:
        return query(root.right, p, q)
    # 不左不右在中间，也就是根节点
    return root


if __name__ == "__main__":
    node = TreeNode(
        val=3,
        left=TreeNode(val=1),
        right=TreeNode(
            val=5,
            left=TreeNode(val=4),
            right=TreeNode(val=6),
        ),
    )
    result = lowest_common_ancestor(node, 1, 5)
    print(result)
