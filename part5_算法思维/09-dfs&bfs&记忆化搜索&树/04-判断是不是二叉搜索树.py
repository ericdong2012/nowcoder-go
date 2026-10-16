"""
BM34 判断是不是二叉搜索树
https://www.nowcoder.com/practice/a69242b39baf45dea217815c7dedb52b?tpId=295&tqId=2288088&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个二叉树根节点，请你判断这棵树是不是二叉搜索树。

二叉搜索树: 每个节点的左子树上的所有节点均小于当前节点且右子树上的所有节点均大于当前节点。
它或者是一棵空树，或者是具有下列性质的二叉树：
若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值

示例1
输入:
{1,2,3}
返回值:
false

示例2
输入:
{2,1,3}
返回值:
true
"""


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


# 方式一：中序遍历递增（注释保留）
# 二叉搜索树的特性就是中序遍历是递增序。既然是判断是否是二叉搜索树，
# 那我们可以使用中序递归遍历。只要之前的节点是二叉树搜索树，
# 那么如果当前的节点小于上一个节点值那么就可以向下判断。


# 更好的解法：限定上下界递归
def is_valid_bst(root):
    """write code here"""
    if root is None:
        return True
    left, right = -2**31 - 1, 2**31
    return helper04(root, left, right)


def helper04(root, left, right):
    if root is None:
        return True
    if root.val < left or root.val > right:
        return False
    return helper04(root.left, left, root.val) and helper04(root.right, root.val, right)


if __name__ == "__main__":
    node = TreeNode(
        val=3,
        left=TreeNode(val=2),
        right=TreeNode(val=5),
    )
    result = is_valid_bst(node)
    print(result)
