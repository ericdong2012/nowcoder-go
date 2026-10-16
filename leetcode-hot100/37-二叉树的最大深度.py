"""104. 二叉树的最大深度
题目：根到最深叶子节点的节点数。
思路：递归 max(左深, 右深) + 1；或 BFS 层数。
复杂度：时间 O(n)，空间 O(h)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        return 1 + max(self.maxDepth(root.left), self.maxDepth(root.right))


if __name__ == "__main__":
    s = Solution()
    #     3
    #    / \
    #   9  20
    #      / \
    #     15  7
    t = TreeNode(3, TreeNode(9), TreeNode(20, TreeNode(15), TreeNode(7)))
    assert s.maxDepth(t) == 3
    assert s.maxDepth(None) == 0
    print("37 OK")
