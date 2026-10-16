"""101. 对称二叉树
题目：判断二叉树是否轴对称。
思路：递归比较 (l.left, r.right) 与 (l.right, r.left)；
     空对空 True，单空 False，值相等再递归。
复杂度：时间 O(n)，空间 O(h)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def mirror(a, b):
            if not a and not b:
                return True
            if not a or not b:
                return False
            return a.val == b.val and mirror(a.left, b.right) and mirror(a.right, b.left)

        return mirror(root, root)


if __name__ == "__main__":
    s = Solution()
    sym = TreeNode(1, TreeNode(2, TreeNode(3), TreeNode(4)),
                   TreeNode(2, TreeNode(4), TreeNode(3)))
    asym = TreeNode(1, TreeNode(2, None, TreeNode(3)), TreeNode(2, None, TreeNode(3)))
    assert s.isSymmetric(sym) is True
    assert s.isSymmetric(asym) is False
    print("39 OK")
