"""226. 翻转二叉树（镜像）
题目：左右子树整体镜像交换。
思路：递归交换每个节点的左右孩子。
复杂度：时间 O(n)，空间 O(h)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        if not root:
            return None
        root.left, root.right = self.invertTree(root.right), self.invertTree(root.left)
        return root


def preorder(root, out):
    if root:
        out.append(root.val)
        preorder(root.left, out)
        preorder(root.right, out)


if __name__ == "__main__":
    s = Solution()
    t = TreeNode(4, TreeNode(2, TreeNode(1), TreeNode(3)), TreeNode(7, TreeNode(6), TreeNode(9)))
    out = []
    preorder(s.invertTree(t), out)
    # 翻转后：4(7(9,6), 2(3,1))，先序 = 4,7,9,6,2,3,1
    assert out == [4, 7, 9, 6, 2, 3, 1]
    print("38 OK")
