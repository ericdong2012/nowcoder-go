"""543. 二叉树的直径
题目：任意两节点间最长路径的边数（可不过根）。
思路：直径 = 某节点的 (左深 + 右深)，DFS 返回深度时顺手更新全局最大。
复杂度：时间 O(n)，空间 O(h)。
"""

from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        self.best = 0

        def depth(node) -> int:
            if not node:
                return 0
            l = depth(node.left)
            r = depth(node.right)
            self.best = max(self.best, l + r)     # 过该节点的路径
            return 1 + max(l, r)

        depth(root)
        return self.best


if __name__ == "__main__":
    s = Solution()
    t = TreeNode(1, TreeNode(2, TreeNode(4), TreeNode(5)), TreeNode(3))
    assert s.diameterOfBinaryTree(t) == 3
    assert s.diameterOfBinaryTree(TreeNode(1)) == 0
    print("40 OK")
