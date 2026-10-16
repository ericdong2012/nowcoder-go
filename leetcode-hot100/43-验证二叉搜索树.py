"""98. 验证二叉搜索树
题目：判断二叉树是否是合法 BST（左<根<右，整体约束）。
思路：递归传上下界 (lo, hi)；或中序遍历严格递增。
复杂度：时间 O(n)，空间 O(h)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        def check(node, lo, hi) -> bool:
            if not node:
                return True
            if not (lo < node.val < hi):      # 注意：不能只和直接父亲比
                return False
            return check(node.left, lo, node.val) and check(node.right, node.val, hi)

        return check(root, float("-inf"), float("inf"))


if __name__ == "__main__":
    s = Solution()
    bad = TreeNode(5, TreeNode(4), TreeNode(6, TreeNode(3), TreeNode(7)))  # 3 < 5 但在右子树
    good = TreeNode(2, TreeNode(1), TreeNode(3))
    assert s.isValidBST(bad) is False
    assert s.isValidBST(good) is True
    print("43 OK")
