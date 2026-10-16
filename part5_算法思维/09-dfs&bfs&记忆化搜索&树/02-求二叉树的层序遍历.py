"""
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
"""

# 层序遍历，就是按层，从上到下，从左到右遍历。bfs


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def level_order(root):
    """write code here"""
    if root is None:
        return None
    res = []
    q = [root]
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
        res.append(temp)
        q = new_queue
    return res


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
    result = level_order(node)
    print(result)
