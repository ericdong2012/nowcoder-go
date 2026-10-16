"""94. 二叉树的中序遍历（迭代栈）
题目：非递归中序遍历。
思路：一路压左；弹栈访问后转向右子树。
复杂度：时间 O(n)，空间 O(h)。
"""

from typing import List, Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        res, stack = [], []
        cur = root
        while cur or stack:
            while cur:                    # 左链全压栈
                stack.append(cur)
                cur = cur.left
            cur = stack.pop()             # 访问
            res.append(cur.val)
            cur = cur.right               # 转向右子树
        return res


def build():
    #     1
    #      \
    #       2
    #      /
    #     3
    n3 = TreeNode(3)
    n2 = TreeNode(2, n3)
    return TreeNode(1, None, n2)


if __name__ == "__main__":
    s = Solution()
    assert s.inorderTraversal(build()) == [1, 3, 2]
    assert s.inorderTraversal(None) == []
    print("36 OK")
