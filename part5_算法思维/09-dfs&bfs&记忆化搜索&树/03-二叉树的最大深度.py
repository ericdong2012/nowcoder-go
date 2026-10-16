"""
BM28 二叉树的最大深度
https://www.nowcoder.com/practice/8a2b2bf6c19b4f23a9bdb9b233eefa73?tpId=295&tqId=642&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj%3Ftab%3D%25E7%25AE%2597%25E6%25B3%2595%25E7%25AF%2587%26topicId%3D295

描述
求给定二叉树的最大深度，
深度是指树的根节点到任一叶子节点路径上节点的数量。
最大深度是所有叶子节点的深度的最大值。
（注：叶子节点是指没有子节点的节点。）

示例1
输入：
{1,2}
返回值：
2

示例2
输入：
{1,2,3,4,#,#,5}
返回值：
3
"""


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def max_depth(root):
    """write code here"""
    # 层序遍历，队列
    if root is None:
        return 0
    q = [root]
    height = 0
    while q:
        # 存储每层的节点值
        temp = []
        # 存储每层的节点
        new_queue = []
        for node in q:
            temp.append(node.val)
            if node.left is not None:
                new_queue.append(node.left)
            if node.right is not None:
                new_queue.append(node.right)
        q = new_queue
        height += 1
    return height


if __name__ == "__main__":
    node = TreeNode(
        val=3,
        left=TreeNode(val=9),
        right=TreeNode(
            val=20,
            left=TreeNode(val=15),
            right=TreeNode(val=7),
        ),
    )
    result = max_depth(node)
    print(result)
