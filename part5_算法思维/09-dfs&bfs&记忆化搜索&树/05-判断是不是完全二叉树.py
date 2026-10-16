"""
BM35 判断是不是完全二叉树
https://www.nowcoder.com/practice/8daa4dff9e36409abba2adbe413d6fae?tpId=295&tqId=2299105&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
给定一个二叉树，确定他是否是一个完全二叉树。

完全二叉树的定义：若二叉树的深度为h，除第 h 层外，其它各层的结点数都达到最大个数，
第 h 层所有的叶子结点都连续集中在最左边，这就是完全二叉树。

示例1
输入：
{1,2,3,4,5,6}
返回值：
true

示例2
输入：
{1,2,3,4,5,6,7}
返回值：
true

示例3
输入：
{1,2,3,4,5,#,6}
返回值：
false
"""

# 层序遍历, 队列：从上到下遍历所有层，每层从左到右，
# 只有次下层和最下层才有叶子节点，其他层出现叶子节点就意味着不是完全二叉树。


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def is_complete_tree(root):
    """write code here"""
    # 层序遍历的中间结果
    stack = [root]
    # flag 标识：遇到叶子节点(左右都空)后置为 True，之后若再遇到非空子节点则不是完全二叉树
    flag = False
    while stack:
        new_stack = []
        for v in stack:
            # 情形1: 已经遇到叶子节点后又出现非空子节点
            # 情形2: 左子树空但右子树不空
            if (flag and (v.left is not None or v.right is not None)) or (v.left is None and v.right is not None):
                return False
            if v.left is not None:
                new_stack.append(v.left)
            if v.right is not None:
                new_stack.append(v.right)
            # 左右子树都空，叶子节点，flag 置 true
            if v.left is None and v.right is None:
                flag = True
        stack = new_stack

    return True


if __name__ == "__main__":
    node = TreeNode(
        val=1,
        left=TreeNode(
            val=2,
            left=TreeNode(val=4),
            right=TreeNode(val=5),
        ),
        right=TreeNode(
            val=3,
            left=None,
            right=TreeNode(val=7),
        ),
    )
    result = is_complete_tree(node)
    print(result)
