"""236. 二叉树的最近公共祖先（LCA）
题目：p、q 的最近公共祖先（节点可以是自己的祖先）。
思路：递归——左找 p/q，右找 p/q；
      当前节点是 p 或 q 直接返回；两侧都有 → 当前即 LCA；
      只有一侧有 → 返回那一侧。
复杂度：时间 O(n)，空间 O(h)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def lowestCommonAncestor(self, root: "TreeNode", p: "TreeNode", q: "TreeNode") -> "TreeNode":
        if not root or root is p or root is q:
            return root
        left = self.lowestCommonAncestor(root.left, p, q)
        right = self.lowestCommonAncestor(root.right, p, q)
        if left and right:
            return root                     # 两侧各找到一个 → 当前是 LCA
        return left or right


if __name__ == "__main__":
    s = Solution()
    n5 = TreeNode(5, TreeNode(6), TreeNode(2, TreeNode(7), TreeNode(4)))
    n1 = TreeNode(1, TreeNode(0), TreeNode(8))
    t = TreeNode(3, n5, n1)
    n4 = n5.right.right
    assert s.lowestCommonAncestor(t, n5, n1) is t
    assert s.lowestCommonAncestor(t, n5, n4) is n5
    print("49 OK")
